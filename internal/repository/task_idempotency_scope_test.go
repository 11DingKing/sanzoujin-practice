package repository_test

import (
	"context"
	"errors"
	"testing"

	"github.com/11DingKing/sanzoujin-practice/internal/domain"
	"github.com/11DingKing/sanzoujin-practice/internal/repository"
	"github.com/11DingKing/sanzoujin-practice/internal/storage"
)

func TestIdempotencyReplayIsScopedToMethodAndPath(t *testing.T) {
	db, err := storage.Open(context.Background(), t.TempDir()+"/practice.db")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	repo := repository.IdempotencyRepo{DB: db.SQL}
	if err := repo.Put(context.Background(), "shared-key", "DELETE", "/api/v1/enrollments/old", "hash-a", `{"status":"withdrawn"}`); err != nil {
		t.Fatal(err)
	}
	if replay, err := repo.Get(context.Background(), "shared-key", "POST", "/api/v1/enrollments"); !errors.Is(err, domain.ErrNotFound) {
		t.Fatalf("different operation replayed a persisted response: response=%s err=%v", replay, err)
	}
	if err := repo.Put(context.Background(), "post-key", "POST", "/api/v1/enrollments", "hash-b", `{"status":"authorized"}`); err != nil {
		t.Fatal(err)
	}
	replay, err := repo.Get(context.Background(), "post-key", "POST", "/api/v1/enrollments")
	if err != nil || replay != `{"status":"authorized"}` {
		t.Fatalf("same operation could not replay its response: response=%s err=%v", replay, err)
	}
}
