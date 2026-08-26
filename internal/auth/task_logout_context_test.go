package auth_test

import (
	"context"
	"testing"
	"time"

	"github.com/11DingKing/sanzoujin-practice/internal/auth"
	"github.com/11DingKing/sanzoujin-practice/internal/domain"
)

type logoutSessionStore struct {
	revoked []string
}

func (s *logoutSessionStore) Create(context.Context, domain.Session) error { return nil }
func (s *logoutSessionStore) Active(context.Context, string, string) (domain.Session, error) {
	return domain.Session{}, domain.ErrNotFound
}
func (s *logoutSessionStore) Purge(context.Context, time.Time) error { return nil }
func (s *logoutSessionStore) Revoke(ctx context.Context, id string, _ time.Time) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	s.revoked = append(s.revoked, id)
	return nil
}

func TestCancelledLogoutDoesNotRevokeSession(t *testing.T) {
	store := &logoutSessionStore{}
	svc := auth.Service{Sessions: store}
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if err := svc.Logout(ctx, domain.Session{ID: "session-cancelled"}); err == nil {
		t.Fatal("cancelled logout unexpectedly returned success")
	}
	if len(store.revoked) != 0 {
		t.Fatalf("cancelled logout revoked sessions: %+v", store.revoked)
	}
	if err := svc.Logout(context.Background(), domain.Session{ID: "session-active"}); err != nil {
		t.Fatalf("active logout failed: %v", err)
	}
	if len(store.revoked) != 1 || store.revoked[0] != "session-active" {
		t.Fatalf("active logout did not revoke the intended session: %+v", store.revoked)
	}
}
