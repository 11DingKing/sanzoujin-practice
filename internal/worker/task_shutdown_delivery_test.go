package worker_test

import (
	"context"
	"sync"
	"testing"
	"time"

	"github.com/11DingKing/sanzoujin-practice/internal/repository"
	"github.com/11DingKing/sanzoujin-practice/internal/storage"
	"github.com/11DingKing/sanzoujin-practice/internal/worker"
)

type shutdownSender struct {
	cancel context.CancelFunc
	mu     sync.Mutex
	sent   []string
}

func (s *shutdownSender) Send(ctx context.Context, topic, payload string) error {
	s.cancel()
	if err := ctx.Err(); err != nil {
		return err
	}
	s.mu.Lock()
	s.sent = append(s.sent, topic+":"+payload)
	s.mu.Unlock()
	return nil
}

func TestShutdownDoesNotAcknowledgeInFlightNotification(t *testing.T) {
	db, err := storage.Open(context.Background(), t.TempDir()+"/practice.db")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	now := time.Now().UTC().Format(time.RFC3339Nano)
	if _, err := db.SQL.ExecContext(context.Background(), `INSERT INTO outbox_messages(id,topic,payload,attempts,next_attempt_at,created_at) VALUES('notice-1','group.departure','{"group":"g-1"}',0,?,?)`, now, now); err != nil {
		t.Fatal(err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	sender := &shutdownSender{cancel: cancel}
	w := worker.Worker{Outbox: repository.OutboxRepo{DB: db.SQL}, Sender: sender}
	_ = w.Process(ctx)

	sender.mu.Lock()
	sentCount := len(sender.sent)
	sender.mu.Unlock()
	if sentCount != 0 {
		t.Fatalf("shutdown worker still emitted notification: %+v", sender.sent)
	}
	var sentAt *string
	var attempts int
	if err := db.SQL.QueryRowContext(context.Background(), `SELECT sent_at,attempts FROM outbox_messages WHERE id='notice-1'`).Scan(&sentAt, &attempts); err != nil {
		t.Fatal(err)
	}
	if sentAt != nil || attempts != 0 {
		t.Fatalf("shutdown worker acknowledged or mutated pending message: sent_at=%v attempts=%d", sentAt, attempts)
	}
}
