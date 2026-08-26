package auth

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"github.com/11DingKing/sanzoujin-practice/internal/domain"
	"time"
)

type UserStore interface {
	ByEmail(context.Context, string) (domain.User, error)
	ByID(context.Context, string) (domain.User, error)
}
type SessionStore interface {
	Create(context.Context, domain.Session) error
	Active(context.Context, string, string) (domain.Session, error)
	Revoke(context.Context, string, time.Time) error
	Purge(context.Context, time.Time) error
}
type Service struct {
	Users    UserStore
	Sessions SessionStore
	TTL      time.Duration
}

func randomID() string          { b := make([]byte, 16); _, _ = rand.Read(b); return hex.EncodeToString(b) }
func tokenHash(t string) string { h := sha256.Sum256([]byte(t)); return hex.EncodeToString(h[:]) }
func (s Service) Login(ctx context.Context, email, password string) (domain.User, domain.Session, string, error) {
	u, e := s.Users.ByEmail(ctx, email)
	if e != nil {
		return u, domain.Session{}, "", e
	}
	if !u.Active || !CheckPassword(u.PasswordHash, password) {
		return u, domain.Session{}, "", errors.New("invalid credentials")
	}
	tok := randomID() + randomID()
	sess := domain.Session{ID: randomID(), UserID: u.ID, TokenHash: tokenHash(tok), ExpiresAt: time.Now().Add(s.TTL), CreatedAt: time.Now()}
	if e = s.Sessions.Create(ctx, sess); e != nil {
		return u, sess, "", e
	}
	return u, sess, tok, nil
}
func (s Service) Authenticate(ctx context.Context, tok string) (domain.User, domain.Session, error) {
	sess, e := s.Sessions.Active(ctx, tokenHash(tok), time.Now().UTC().Format(time.RFC3339Nano))
	if e != nil {
		return domain.User{}, sess, e
	}
	u, e := s.Users.ByID(ctx, sess.UserID)
	return u, sess, e
}
func (s Service) Logout(ctx context.Context, sess domain.Session) error {
	return s.Sessions.Revoke(ctx, sess.ID, time.Now())
}
