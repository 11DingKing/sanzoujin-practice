package auth

import (
	"context"
	"github.com/11DingKing/sanzoujin-practice/internal/domain"
	"testing"
	"time"
)

type fakeUsers struct {
	byEmail map[string]domain.User
	byID    map[string]domain.User
}

func (f fakeUsers) ByEmail(_ context.Context, e string) (domain.User, error) {
	u, ok := f.byEmail[e]
	if !ok {
		return domain.User{}, domain.ErrNotFound
	}
	return u, nil
}
func (f fakeUsers) ByID(_ context.Context, id string) (domain.User, error) {
	u, ok := f.byID[id]
	if !ok {
		return domain.User{}, domain.ErrNotFound
	}
	return u, nil
}

type fakeSessions struct {
	created []domain.Session
	active  domain.Session
	revoked bool
}

func (f *fakeSessions) Create(_ context.Context, s domain.Session) error {
	f.created = append(f.created, s)
	f.active = s
	return nil
}
func (f *fakeSessions) Active(_ context.Context, _ string, _ string) (domain.Session, error) {
	if f.active.ID == "" {
		return domain.Session{}, domain.ErrExpired
	}
	return f.active, nil
}
func (f *fakeSessions) Revoke(_ context.Context, _ string, _ time.Time) error {
	f.revoked = true
	return nil
}
func (f *fakeSessions) Purge(context.Context, time.Time) error { return nil }

func TestPasswordHashing(t *testing.T) {
	h := HashPassword("secret")
	if h == "" {
		t.Fatal("empty hash")
	}
	if !CheckPassword(h, "secret") {
		t.Fatal("valid password rejected")
	}
	if CheckPassword(h, "wrong") {
		t.Fatal("wrong password accepted")
	}
	if HashPassword("secret") == HashPassword("other") {
		t.Fatal("hash collision")
	}
}
func TestLoginAndLogout(t *testing.T) {
	u := domain.User{ID: "u1", Email: "s@example.com", PasswordHash: HashPassword("pw"), Role: domain.RoleStudent, Active: true}
	users := fakeUsers{byEmail: map[string]domain.User{u.Email: u}, byID: map[string]domain.User{u.ID: u}}
	sessions := &fakeSessions{}
	s := Service{Users: users, Sessions: repositoryAdapter{sessions}, TTL: time.Hour}
	got, sess, tok, e := s.Login(context.Background(), u.Email, "pw")
	if e != nil {
		t.Fatal(e)
	}
	if got.ID != u.ID || sess.ID == "" || tok == "" {
		t.Fatal("login response")
	}
	if len(sessions.created) != 1 {
		t.Fatal("session not saved")
	}
	if e = s.Logout(context.Background(), sess); e != nil {
		t.Fatal(e)
	}
	if !sessions.revoked {
		t.Fatal("not revoked")
	}
}
func TestLoginRejectsInactiveAndBadPassword(t *testing.T) {
	for _, u := range []domain.User{{ID: "inactive", Email: "i", PasswordHash: HashPassword("pw"), Active: false}, {ID: "active", Email: "a", PasswordHash: HashPassword("pw"), Active: true}} {
		users := fakeUsers{byEmail: map[string]domain.User{u.Email: u}}
		s := Service{Users: users, Sessions: repositoryAdapter{&fakeSessions{}}, TTL: time.Hour}
		pw := "pw"
		if u.Active {
			pw = "bad"
		}
		_, _, _, e := s.Login(context.Background(), u.Email, pw)
		if e == nil {
			t.Fatalf("accepted invalid user %s", u.Email)
		}
	}
}
func TestAuthenticateExpired(t *testing.T) {
	sessions := &fakeSessions{}
	s := Service{Users: fakeUsers{}, Sessions: repositoryAdapter{sessions}, TTL: time.Hour}
	if _, _, e := s.Authenticate(context.Background(), "token"); e != domain.ErrExpired {
		t.Fatalf("got %v", e)
	}
}

type repositoryAdapter struct{ *fakeSessions }
