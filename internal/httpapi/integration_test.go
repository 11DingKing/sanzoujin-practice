package httpapi_test

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/11DingKing/sanzoujin-practice/internal/app"
	"github.com/11DingKing/sanzoujin-practice/internal/auth"
	"github.com/11DingKing/sanzoujin-practice/internal/config"
	"github.com/11DingKing/sanzoujin-practice/internal/domain"
	"github.com/11DingKing/sanzoujin-practice/internal/repository"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"testing"
	"time"
)

func call(h http.Handler, method, path string, body any, token string) *httptest.ResponseRecorder {
	var r io.Reader
	if body != nil {
		b, _ := json.Marshal(body)
		r = bytes.NewReader(b)
	}
	req := httptest.NewRequest(method, path, r)
	if token != "" {
		req.Header.Set("Authorization", "Bearer "+token)
	}
	rr := httptest.NewRecorder()
	h.ServeHTTP(rr, req)
	return rr
}
func TestLoginSessionLogoutWithRealDatabase(t *testing.T) {
	logger := slog.New(slog.NewTextHandler(io.Discard, nil))
	srv, _, db, err := app.Build(context.Background(), config.Config{DBPath: filepath.Join(t.TempDir(), "api.db"), SessionTTL: time.Hour, WorkerInterval: time.Hour}, logger)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	u := domain.User{ID: "student-1", OrgID: "school-1", Name: "学生甲", Email: "student@example.com", PasswordHash: auth.HashPassword("safe-password"), Role: domain.RoleStudent, Active: true, CreatedAt: time.Now()}
	if err := (repository.UserRepo{DB: db.SQL}).Create(context.Background(), u); err != nil {
		t.Fatal(err)
	}
	h := srv.Handler()
	login := call(h, http.MethodPost, "/api/v1/auth/login", map[string]string{"email": u.Email, "password": "safe-password"}, "")
	if login.Code != http.StatusOK {
		t.Fatalf("login status %d", login.Code)
	}
	var result struct {
		Token string `json:"token"`
	}
	if err := json.NewDecoder(login.Body).Decode(&result); err != nil {
		t.Fatal(err)
	}
	if result.Token == "" {
		t.Fatal("missing token")
	}
	if got := call(h, http.MethodGet, "/api/v1/projects", nil, result.Token).Code; got != http.StatusOK {
		t.Fatalf("protected status %d", got)
	}
	if got := call(h, http.MethodPost, "/api/v1/auth/logout", nil, result.Token).Code; got != http.StatusNoContent {
		t.Fatalf("logout status %d", got)
	}
	if got := call(h, http.MethodGet, "/api/v1/projects", nil, result.Token).Code; got != http.StatusUnauthorized {
		t.Fatalf("revoked status %d", got)
	}
}
func TestExpiredSessionIsRejected(t *testing.T) {
	logger := slog.New(slog.NewTextHandler(io.Discard, nil))
	srv, _, db, err := app.Build(context.Background(), config.Config{DBPath: filepath.Join(t.TempDir(), "expiry.db"), SessionTTL: -time.Second, WorkerInterval: time.Hour}, logger)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	u := domain.User{ID: "coordinator-1", OrgID: "district-1", Name: "协调员", Email: "coordinator@example.com", PasswordHash: auth.HashPassword("pw"), Role: domain.RoleCoordinator, Active: true, CreatedAt: time.Now()}
	if err := (repository.UserRepo{DB: db.SQL}).Create(context.Background(), u); err != nil {
		t.Fatal(err)
	}
	h := srv.Handler()
	login := call(h, http.MethodPost, "/api/v1/auth/login", map[string]string{"email": u.Email, "password": "pw"}, "")
	var result struct {
		Token string `json:"token"`
	}
	_ = json.NewDecoder(login.Body).Decode(&result)
	if got := call(h, http.MethodGet, "/api/v1/projects", nil, result.Token).Code; got != http.StatusUnauthorized {
		t.Fatalf("expired status %d", got)
	}
}
