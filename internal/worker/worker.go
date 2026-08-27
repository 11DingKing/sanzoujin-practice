package worker

import (
	"context"
	"errors"
	"github.com/11DingKing/sanzoujin-practice/internal/notify"
	"github.com/11DingKing/sanzoujin-practice/internal/repository"
	"log/slog"
	"sync"
	"time"
)

type Worker struct {
	Outbox   repository.OutboxRepo
	Sender   notify.Sender
	Interval time.Duration
	Logger   *slog.Logger
	stop     chan struct{}
	done     chan struct{}
	mu       sync.Mutex
	cancel   context.CancelFunc
}

func (w *Worker) Start(ctx context.Context) {
	w.stop = make(chan struct{})
	w.done = make(chan struct{})
	runCtx, cancel := context.WithCancel(ctx)
	w.mu.Lock()
	w.cancel = cancel
	w.mu.Unlock()
	go w.loop(runCtx)
}
func (w *Worker) loop(ctx context.Context) {
	defer close(w.done)
	t := time.NewTicker(w.Interval)
	defer t.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case <-w.stop:
			return
		case <-t.C:
			_ = w.Process(ctx)
		}
	}
}
func (w *Worker) Stop() {
	if w.stop != nil {
		// Cancel in-flight deliveries first so the worker stops calling the
		// sender and marks the message for retry instead of as delivered. This
		// keeps the outbox row recoverable after restart (sent_at stays NULL).
		w.mu.Lock()
		if w.cancel != nil {
			w.cancel()
		}
		w.mu.Unlock()
		close(w.stop)
		<-w.done
	}
}
func (w *Worker) Process(ctx context.Context) error {
	items, e := w.Outbox.Due(ctx, 20)
	if e != nil {
		return e
	}
	for _, m := range items {
		if err := w.Sender.Send(notify.DeliveryContext(ctx), m.Topic, m.Payload); err != nil {
			markCtx := context.WithoutCancel(ctx)
			if errors.Is(err, context.Canceled) {
				// Shutdown aborted an in-flight delivery. The message was never
				// confirmed delivered, so keep it immediately recoverable: leave
				// attempts unchanged and requeue for now. A restart (or the next
				// poll, since Stop() owns the lifecycle) re-delivers it.
				if mErr := w.Outbox.Mark(markCtx, m.ID, m.Attempts, err.Error(), time.Now(), false); mErr != nil && w.Logger != nil {
					w.Logger.Error("outbox mark shutdown", "id", m.ID, "error", mErr)
				}
				continue
			}
			// Genuine delivery failure: apply exponential backoff for retry.
			attempt := m.Attempts + 1
			delay := time.Duration(1<<min(attempt, 6)) * time.Second
			if mErr := w.Outbox.Mark(markCtx, m.ID, attempt, err.Error(), time.Now().Add(delay), false); mErr != nil && w.Logger != nil {
				w.Logger.Error("outbox mark retry", "id", m.ID, "error", mErr)
			}
			continue
		}
		if mErr := w.Outbox.Mark(context.WithoutCancel(ctx), m.ID, m.Attempts, "", time.Time{}, true); mErr != nil && w.Logger != nil {
			w.Logger.Error("outbox mark sent", "id", m.ID, "error", mErr)
		}
	}
	return nil
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
