package notify

import (
	"context"
	"errors"
)

type Sender interface {
	Send(context.Context, string, string) error
}

// DeliveryContext derives the context used for an in-flight delivery.
//
// Unlike a decoupled (non-cancellable) context, the delivery context still
// propagates the poller's shutdown signal: when the service shuts down
// mid-delivery the sender observes the cancellation, aborts the send, and the
// worker marks the message for retry instead of as delivered. A message that
// was never confirmed delivered therefore stays unsent (sent_at IS NULL) and is
// recovered on restart. If the sender completed before cancellation reached it,
// the success is still acknowledged. Both paths are safe.
func DeliveryContext(ctx context.Context) context.Context {
	return ctx
}

type MemorySender struct {
	Fail bool
	Sent []string
}

func (s *MemorySender) Send(ctx context.Context, topic, payload string) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	if s.Fail {
		return errors.New("temporary notification failure")
	}
	s.Sent = append(s.Sent, topic+":"+payload)
	return nil
}
