package worker

import (
	"context"
	"github.com/11DingKing/sanzoujin-practice/internal/notify"
	"github.com/11DingKing/sanzoujin-practice/internal/repository"
	"log/slog"
	"time"
)

type Worker struct {
	Outbox   repository.OutboxRepo
	Sender   notify.Sender
	Interval time.Duration
	Logger   *slog.Logger
	stop     chan struct{}
	done     chan struct{}
}

func (w *Worker) Start(ctx context.Context) {
	w.stop = make(chan struct{})
	w.done = make(chan struct{})
	go w.loop(ctx)
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
		if err := w.Outbox.Acknowledge(ctx, m.ID); err != nil {
			return err
		}
		if err := w.Sender.Send(ctx, m.Topic, m.Payload); err != nil {
			attempt := m.Attempts + 1
			delay := time.Duration(1<<min(attempt, 6)) * time.Second
			_ = w.Outbox.Mark(ctx, m.ID, attempt, err.Error(), time.Now().Add(delay), false)
			continue
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
