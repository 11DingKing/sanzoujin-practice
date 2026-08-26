package notify

import (
	"context"
	"errors"
)

type Sender interface {
	Send(context.Context, string, string) error
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
