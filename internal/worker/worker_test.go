package worker

import (
	"context"
	"github.com/11DingKing/sanzoujin-practice/internal/domain"
	"github.com/11DingKing/sanzoujin-practice/internal/notify"
	"github.com/11DingKing/sanzoujin-practice/internal/repository"
	"github.com/11DingKing/sanzoujin-practice/internal/storage"
	"io"
	"log/slog"
	"path/filepath"
	"sync"
	"testing"
	"time"
)

// blockingSender signals readiness exactly once, then blocks on the delivery
// context until shutdown cancels it. The arrived channel is set at construction
// and never reassigned, so waiting on it is race-free.
type blockingSender struct {
	once    sync.Once
	arrived chan struct{}
}

func newBlockingSender() *blockingSender {
	return &blockingSender{arrived: make(chan struct{})}
}

func (s *blockingSender) Send(ctx context.Context, topic, payload string) error {
	s.once.Do(func() { close(s.arrived) })
	<-ctx.Done()
	return ctx.Err()
}

func newTestWorker(t *testing.T) (*Worker, *storage.DB) {
	t.Helper()
	db, err := storage.Open(context.Background(), filepath.Join(t.TempDir(), "worker.db"))
	if err != nil {
		t.Fatalf("open db: %v", err)
	}
	t.Cleanup(func() { _ = db.Close() })
	return &Worker{Outbox: repository.OutboxRepo{DB: db.SQL}, Sender: &notify.MemorySender{}, Interval: 10 * time.Millisecond, Logger: slog.New(slog.NewTextHandler(io.Discard, nil))}, db
}

func enqueueDue(t *testing.T, repo repository.OutboxRepo, id string) {
	t.Helper()
	if err := repo.Enqueue(context.Background(), domain.OutboxMessage{ID: id, Topic: "practice.general", Payload: "p", NextAttemptAt: time.Now(), CreatedAt: time.Now()}); err != nil {
		t.Fatalf("enqueue: %v", err)
	}
}

// TestShutdownAbortsDeliveryAndKeepsMessageRecoverable reproduces the reported
// bug: shutdown fires exactly when delivery begins. The worker must not mark the
// message as delivered; it must stay recoverable for the next restart.
func TestShutdownAbortsDeliveryAndKeepsMessageRecoverable(t *testing.T) {
	db, err := storage.Open(context.Background(), filepath.Join(t.TempDir(), "shutdown.db"))
	if err != nil {
		t.Fatalf("open db: %v", err)
	}
	defer db.Close()
	repo := repository.OutboxRepo{DB: db.SQL}
	enqueueDue(t, repo, "msg-shutdown")

	sender := newBlockingSender()
	w := &Worker{Outbox: repo, Sender: sender, Interval: 10 * time.Millisecond, Logger: slog.New(slog.NewTextHandler(io.Discard, nil))}
	w.Start(context.Background())

	// Wait until the worker has picked the message up and is mid-delivery.
	<-sender.arrived
	// Shutdown arrives precisely while delivery is in flight.
	w.Stop()

	// The row must NOT be marked as delivered.
	var sent *string
	if err := db.SQL.QueryRowContext(context.Background(), `SELECT sent_at FROM outbox_messages WHERE id=?`, "msg-shutdown").Scan(&sent); err != nil {
		t.Fatalf("query: %v", err)
	}
	if sent != nil {
		t.Fatalf("message marked sent_at=%q during shutdown; must stay recoverable", *sent)
	}
}

// TestRestartRecoversAbortedMessage: after a shutdown-aborted delivery, a fresh
// worker (simulating restart) re-delivers and marks the message sent.
func TestRestartRecoversAbortedMessage(t *testing.T) {
	db, err := storage.Open(context.Background(), filepath.Join(t.TempDir(), "restart.db"))
	if err != nil {
		t.Fatalf("open db: %v", err)
	}
	defer db.Close()
	repo := repository.OutboxRepo{DB: db.SQL}
	enqueueDue(t, repo, "msg-restart")

	sender := newBlockingSender()
	w := &Worker{Outbox: repo, Sender: sender, Interval: 10 * time.Millisecond, Logger: slog.New(slog.NewTextHandler(io.Discard, nil))}
	w.Start(context.Background())
	<-sender.arrived
	w.Stop()

	// Simulate restart: new worker with a real (non-blocking) sender.
	restart := &Worker{Outbox: repo, Sender: &notify.MemorySender{}, Interval: 10 * time.Millisecond, Logger: slog.New(slog.NewTextHandler(io.Discard, nil))}
	restart.Start(context.Background())
	defer restart.Stop()

	deadline := time.Now().Add(2 * time.Second)
	for time.Now().Before(deadline) {
		var sent *string
		_ = db.SQL.QueryRowContext(context.Background(), `SELECT sent_at FROM outbox_messages WHERE id=?`, "msg-restart").Scan(&sent)
		if sent != nil {
			return
		}
		time.Sleep(5 * time.Millisecond)
	}
	t.Fatal("message not recovered and sent after restart")
}

// TestNormalDeliveryMarksSent ensures the happy path still confirms delivery.
func TestNormalDeliveryMarksSent(t *testing.T) {
	w, db := newTestWorker(t)
	repo := repository.OutboxRepo{DB: db.SQL}
	enqueueDue(t, repo, "msg-normal")
	w.Start(context.Background())
	defer w.Stop()

	deadline := time.Now().Add(2 * time.Second)
	for time.Now().Before(deadline) {
		var sent *string
		_ = db.SQL.QueryRowContext(context.Background(), `SELECT sent_at FROM outbox_messages WHERE id=?`, "msg-normal").Scan(&sent)
		if sent != nil {
			return
		}
		time.Sleep(5 * time.Millisecond)
	}
	t.Fatal("normal delivery not marked sent")
}

// TestFailedDeliverySchedulesRetry keeps a failed message recoverable.
func TestFailedDeliverySchedulesRetry(t *testing.T) {
	db, err := storage.Open(context.Background(), filepath.Join(t.TempDir(), "retry.db"))
	if err != nil {
		t.Fatalf("open db: %v", err)
	}
	defer db.Close()
	repo := repository.OutboxRepo{DB: db.SQL}
	enqueueDue(t, repo, "msg-retry")

	sender := &notify.MemorySender{Fail: true}
	w := &Worker{Outbox: repo, Sender: sender, Interval: 10 * time.Millisecond, Logger: slog.New(slog.NewTextHandler(io.Discard, nil))}
	w.Start(context.Background())
	defer w.Stop()

	deadline := time.Now().Add(2 * time.Second)
	for time.Now().Before(deadline) {
		var attempts int
		var lastError string
		var sent *string
		_ = db.SQL.QueryRowContext(context.Background(), `SELECT attempts, last_error, sent_at FROM outbox_messages WHERE id=?`, "msg-retry").Scan(&attempts, &lastError, &sent)
		if attempts > 0 && lastError != "" && sent == nil {
			return
		}
		time.Sleep(5 * time.Millisecond)
	}
	t.Fatal("failed delivery did not schedule retry")
}
