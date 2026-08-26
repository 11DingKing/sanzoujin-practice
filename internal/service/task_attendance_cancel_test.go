package service

import (
	"context"
	"errors"
	"path/filepath"
	"testing"

	"github.com/11DingKing/sanzoujin-practice/internal/audit"
	"github.com/11DingKing/sanzoujin-practice/internal/repository"
	"github.com/11DingKing/sanzoujin-practice/internal/storage"
)

func TestCancelledCheckoutDoesNotPersistCompletion(t *testing.T) {
	ctx := context.Background()
	db, err := storage.Open(ctx, filepath.Join(t.TempDir(), "attendance.db"))
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	for _, statement := range []string{
		`INSERT INTO users(id,org_id,name,email,password_hash,role,created_at) VALUES('student','school','Student','student@example.test','x','student','2026-08-26T00:00:00Z')`,
		`INSERT INTO users(id,org_id,name,email,password_hash,role,created_at) VALUES('guardian','school','Guardian','guardian@example.test','x','guardian','2026-08-26T00:00:00Z')`,
		`INSERT INTO users(id,org_id,name,email,password_hash,role,created_at) VALUES('teacher','school','Teacher','teacher@example.test','x','teacher','2026-08-26T00:00:00Z')`,
		`INSERT INTO venues(id,name,address,capacity,created_at) VALUES('venue','Community Hall','Road 1',10,'2026-08-26T00:00:00Z')`,
		`INSERT INTO practice_projects(id,org_id,title,description,venue_id,capacity,starts_at,ends_at,status,version,created_by,created_at,updated_at) VALUES('project','school','Practice','Service','venue',10,'2026-08-26T00:00:00Z','2026-09-26T00:00:00Z','published',1,'teacher','2026-08-26T00:00:00Z','2026-08-26T00:00:00Z')`,
		`INSERT INTO enrollments(id,project_id,student_id,guardian_id,status,qualifications,idempotency_key,created_at,updated_at) VALUES('cancelled','project','student','guardian','matched','[]','cancelled','2026-08-26T00:00:00Z','2026-08-26T00:00:00Z')`,
		`INSERT INTO enrollments(id,project_id,student_id,guardian_id,status,qualifications,idempotency_key,created_at,updated_at) VALUES('normal','project','teacher','guardian','matched','[]','normal','2026-08-26T00:00:00Z','2026-08-26T00:00:00Z')`,
		`INSERT INTO attendance(id,enrollment_id,state,checked_in_at,minutes,note) VALUES('a1','cancelled','present','2026-08-26T00:00:00Z',0,'arrived')`,
		`INSERT INTO attendance(id,enrollment_id,state,checked_in_at,minutes,note) VALUES('a2','normal','present','2026-08-26T00:00:00Z',0,'arrived')`,
	} {
		if _, err := db.SQL.ExecContext(ctx, statement); err != nil {
			t.Fatal(err)
		}
	}
	svc := AttendanceService{
		Repo:        repository.AttendanceRepo{DB: db.SQL},
		Enrollments: repository.EnrollmentRepo{DB: db.SQL},
		Audit:       audit.Service{Repo: repository.AuditRepo{DB: db.SQL}},
	}
	cancelled, cancel := context.WithCancel(context.Background())
	cancel()
	if _, err := svc.CheckOut(cancelled, "cancelled", "teacher", "school"); !errors.Is(err, context.Canceled) {
		t.Fatalf("cancelled checkout returned %v", err)
	}
	var attendanceState, enrollmentState string
	var checkedOut any
	var minutes int
	if err := db.SQL.QueryRowContext(ctx, `SELECT state,checked_out_at,minutes FROM attendance WHERE enrollment_id='cancelled'`).Scan(&attendanceState, &checkedOut, &minutes); err != nil {
		t.Fatal(err)
	}
	if err := db.SQL.QueryRowContext(ctx, `SELECT status FROM enrollments WHERE id='cancelled'`).Scan(&enrollmentState); err != nil {
		t.Fatal(err)
	}
	if attendanceState != "present" || checkedOut != nil || minutes != 0 || enrollmentState != "matched" {
		t.Fatalf("cancelled checkout changed persisted state: attendance=%s checked_out=%v minutes=%d enrollment=%s", attendanceState, checkedOut, minutes, enrollmentState)
	}
	if _, err := svc.CheckOut(ctx, "normal", "teacher", "school"); err != nil {
		t.Fatalf("active checkout failed: %v", err)
	}
	if err := db.SQL.QueryRowContext(ctx, `SELECT state,checked_out_at,minutes FROM attendance WHERE enrollment_id='normal'`).Scan(&attendanceState, &checkedOut, &minutes); err != nil {
		t.Fatal(err)
	}
	if err := db.SQL.QueryRowContext(ctx, `SELECT status FROM enrollments WHERE id='normal'`).Scan(&enrollmentState); err != nil {
		t.Fatal(err)
	}
	if attendanceState != "left" || checkedOut == nil || minutes < 1 || enrollmentState != "completed" {
		t.Fatalf("active checkout incomplete: attendance=%s checked_out=%v minutes=%d enrollment=%s", attendanceState, checkedOut, minutes, enrollmentState)
	}
}
