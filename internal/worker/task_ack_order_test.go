package worker_test

import (
	"context"
	"errors"
	"github.com/11DingKing/sanzoujin-practice/internal/repository"
	"github.com/11DingKing/sanzoujin-practice/internal/storage"
	"github.com/11DingKing/sanzoujin-practice/internal/worker"
	"testing"
	"time"
)

type selectiveSender struct{ sent []string }

func (s *selectiveSender) Send(_ context.Context, topic, _ string) error {
	if topic == "risk.failed" {
		return errors.New("gateway unavailable")
	}
	s.sent = append(s.sent, topic)
	return nil
}

func TestOutboxAcknowledgesOnlyAfterSuccessfulSend(t *testing.T) {
	db, err := storage.Open(context.Background(), t.TempDir()+"/practice.db")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	now := time.Now().UTC().Format(time.RFC3339Nano)
	for _, row := range []struct{ id, topic string }{{"msg-fail", "risk.failed"}, {"msg-ok", "group.ready"}} {
		if _, err := db.SQL.ExecContext(context.Background(), `INSERT INTO outbox_messages(id,topic,payload,attempts,next_attempt_at,created_at) VALUES(?,?,?,0,?,?)`, row.id, row.topic, "{}", now, now); err != nil {
			t.Fatal(err)
		}
	}
	sender := &selectiveSender{}
	w := worker.Worker{Outbox: repository.OutboxRepo{DB: db.SQL}, Sender: sender}
	if err := w.Process(context.Background()); err != nil {
		t.Fatal(err)
	}
	var failedSent, okSent *string
	var failedAttempts int
	if err := db.SQL.QueryRow(`SELECT sent_at,attempts FROM outbox_messages WHERE id='msg-fail'`).Scan(&failedSent, &failedAttempts); err != nil {
		t.Fatal(err)
	}
	if err := db.SQL.QueryRow(`SELECT sent_at FROM outbox_messages WHERE id='msg-ok'`).Scan(&okSent); err != nil {
		t.Fatal(err)
	}
	if failedSent != nil || failedAttempts != 1 {
		t.Fatalf("failed delivery was acknowledged or not scheduled: sent=%v attempts=%d", failedSent, failedAttempts)
	}
	if okSent == nil || len(sender.sent) != 1 || sender.sent[0] != "group.ready" {
		t.Fatalf("successful delivery was not sent and acknowledged: sent=%v calls=%v", okSent, sender.sent)
	}
}
