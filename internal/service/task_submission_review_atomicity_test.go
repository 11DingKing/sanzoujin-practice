package service

import (
	"context"
	"path/filepath"
	"testing"

	"github.com/11DingKing/sanzoujin-practice/internal/audit"
	"github.com/11DingKing/sanzoujin-practice/internal/domain"
	"github.com/11DingKing/sanzoujin-practice/internal/repository"
	"github.com/11DingKing/sanzoujin-practice/internal/storage"
)

func TestSubmissionReviewIsAtomicWithReviewerMetadata(t *testing.T) {
	ctx := context.Background()
	db, err := storage.Open(ctx, filepath.Join(t.TempDir(), "review.db"))
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	for _, statement := range []string{
		`INSERT INTO users(id,org_id,name,email,password_hash,role,created_at) VALUES('student','school','Student','student@example.test','x','student','2026-08-26T00:00:00Z')`,
		`INSERT INTO users(id,org_id,name,email,password_hash,role,created_at) VALUES('reviewer','street','Reviewer','reviewer@example.test','x','mentor','2026-08-26T00:00:00Z')`,
		`INSERT INTO venues(id,name,address,capacity,created_at) VALUES('venue','Community Hall','Road 1',10,'2026-08-26T00:00:00Z')`,
		`INSERT INTO practice_projects(id,org_id,title,description,venue_id,capacity,starts_at,ends_at,status,version,created_by,created_at,updated_at) VALUES('project','street','Research','Field research','venue',10,'2026-08-26T00:00:00Z','2026-09-26T00:00:00Z','published',1,'reviewer','2026-08-26T00:00:00Z','2026-08-26T00:00:00Z')`,
		`INSERT INTO submissions(id,project_id,student_id,version,status,content) VALUES('blocked','project','student',1,'submitted','survey')`,
		`INSERT INTO submissions(id,project_id,student_id,version,status,content) VALUES('normal','project','student',1,'submitted','interview')`,
		`CREATE TRIGGER reject_review_metadata BEFORE UPDATE OF reviewer_id ON submissions WHEN NEW.id='blocked' BEGIN SELECT RAISE(ABORT, 'review metadata unavailable'); END`,
	} {
		if _, err := db.SQL.ExecContext(ctx, statement); err != nil {
			t.Fatal(err)
		}
	}
	repo := repository.SubmissionRepo{DB: db.SQL}
	svc := SubmissionService{Repo: repo, Audit: audit.Service{Repo: repository.AuditRepo{DB: db.SQL}}}
	if err := svc.Review(ctx, "blocked", domain.SubmissionApproved, 1, "reviewer", "street"); err == nil {
		t.Fatal("review unexpectedly succeeded when reviewer metadata was rejected")
	}
	var status string
	var version int
	var reviewedAt, reviewer any
	if err := db.SQL.QueryRowContext(ctx, `SELECT status,version,reviewed_at,reviewer_id FROM submissions WHERE id='blocked'`).Scan(&status, &version, &reviewedAt, &reviewer); err != nil {
		t.Fatal(err)
	}
	if status != "submitted" || version != 1 || reviewedAt != nil || reviewer != nil {
		t.Fatalf("failed review leaked partial state: status=%s version=%d reviewed_at=%v reviewer=%v", status, version, reviewedAt, reviewer)
	}
	if err := svc.Review(ctx, "normal", domain.SubmissionApproved, 1, "reviewer", "street"); err != nil {
		t.Fatalf("normal review failed: %v", err)
	}
	if err := db.SQL.QueryRowContext(ctx, `SELECT status,version,reviewed_at,reviewer_id FROM submissions WHERE id='normal'`).Scan(&status, &version, &reviewedAt, &reviewer); err != nil {
		t.Fatal(err)
	}
	if status != "approved" || version != 2 || reviewedAt == nil || reviewer != "reviewer" {
		t.Fatalf("normal review incomplete: status=%s version=%d reviewed_at=%v reviewer=%v", status, version, reviewedAt, reviewer)
	}
}
