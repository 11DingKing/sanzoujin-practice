package audit_test

import (
	"context"
	"testing"

	"github.com/11DingKing/sanzoujin-practice/internal/audit"
	"github.com/11DingKing/sanzoujin-practice/internal/repository"
	"github.com/11DingKing/sanzoujin-practice/internal/storage"
)

func TestCancelledRequestDoesNotAppendAuditEvent(t *testing.T) {
	db, err := storage.Open(context.Background(), t.TempDir()+"/practice.db")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	repo := repository.AuditRepo{DB: db.SQL}
	svc := audit.Service{Repo: repo}
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if err := svc.Record(ctx, "school", "guardian", "authorization", "auth-1", "approve", "ok", "request-cancelled"); err == nil {
		t.Fatal("cancelled audit append unexpectedly succeeded")
	}
	events, err := repo.List(context.Background(), "authorization", "auth-1")
	if err != nil {
		t.Fatal(err)
	}
	if len(events) != 0 {
		t.Fatalf("cancelled request changed audit chain: %+v", events)
	}
	if err := svc.Record(context.Background(), "school", "guardian", "authorization", "auth-1", "approve", "ok", "request-live"); err != nil {
		t.Fatalf("live audit append failed: %v", err)
	}
	events, err = repo.List(context.Background(), "authorization", "auth-1")
	if err != nil || len(events) != 1 {
		t.Fatalf("live audit event missing: events=%+v err=%v", events, err)
	}
}
