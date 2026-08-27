package service_test

import (
	"context"
	"testing"
	"time"

	"github.com/11DingKing/sanzoujin-practice/internal/audit"
	"github.com/11DingKing/sanzoujin-practice/internal/repository"
	"github.com/11DingKing/sanzoujin-practice/internal/service"
	"github.com/11DingKing/sanzoujin-practice/internal/storage"
)

func newGroupAssignmentService(t *testing.T, suffix string) (*storage.DB, service.GroupService) {
	t.Helper()
	db, err := storage.Open(context.Background(), t.TempDir()+"/practice.db")
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = db.Close() })
	now := time.Now().UTC().Format(time.RFC3339Nano)
	for _, stmt := range []string{
		`INSERT INTO users(id,org_id,name,email,password_hash,role,active,created_at) VALUES('student-` + suffix + `','school','Student','student-` + suffix + `@example.test','x','student',1,'` + now + `')`,
		`INSERT INTO users(id,org_id,name,email,password_hash,role,active,created_at) VALUES('guardian-` + suffix + `','school','Guardian','guardian-` + suffix + `@example.test','x','guardian',1,'` + now + `')`,
		`INSERT INTO venues(id,name,address,capacity,open,created_at) VALUES('venue-` + suffix + `','Community','Road',20,1,'` + now + `')`,
		`INSERT INTO practice_projects(id,org_id,title,description,venue_id,capacity,capacity_used,starts_at,ends_at,status,version,created_by,created_at,updated_at) VALUES('project-` + suffix + `','school','Summer Practice','Community visit','venue-` + suffix + `',10,1,'` + now + `','` + now + `','published',1,'student-` + suffix + `','` + now + `','` + now + `')`,
		`INSERT INTO enrollments(id,project_id,student_id,guardian_id,status,qualifications,idempotency_key,created_at,updated_at) VALUES('enrollment-` + suffix + `','project-` + suffix + `','student-` + suffix + `','guardian-` + suffix + `','authorized','[]','key-` + suffix + `','` + now + `','` + now + `')`,
		`INSERT INTO groups(id,project_id,name,capacity,status,version,created_at) VALUES('group-` + suffix + `','project-` + suffix + `','Team',2,'forming',1,'` + now + `')`,
	} {
		if _, err := db.SQL.ExecContext(context.Background(), stmt); err != nil {
			t.Fatalf("seed %s: %v", suffix, err)
		}
	}
	groups := repository.GroupRepo{DB: db.SQL}
	enrollments := repository.EnrollmentRepo{DB: db.SQL}
	return db, service.GroupService{DB: db, Groups: groups, Enrollments: enrollments, Audit: audit.Service{Repo: repository.AuditRepo{DB: db.SQL}}}
}

func TestGroupAssignmentRollsBackAsAUnit(t *testing.T) {
	db, svc := newGroupAssignmentService(t, "rollback")
	if _, err := db.SQL.ExecContext(context.Background(), `CREATE TRIGGER reject_match BEFORE UPDATE OF status ON enrollments WHEN NEW.status='matched' BEGIN SELECT RAISE(ABORT, 'status store unavailable'); END`); err != nil {
		t.Fatal(err)
	}
	err := svc.Add(context.Background(), "group-rollback", "enrollment-rollback", "school", "coordinator")
	if err == nil {
		t.Fatal("group assignment unexpectedly succeeded")
	}
	members, err := svc.Members(context.Background(), "group-rollback")
	if err != nil {
		t.Fatal(err)
	}
	if len(members) != 0 {
		t.Fatalf("failed assignment left persisted group members: %+v", members)
	}

	_, healthy := newGroupAssignmentService(t, "success")
	if err := healthy.Add(context.Background(), "group-success", "enrollment-success", "school", "coordinator"); err != nil {
		t.Fatalf("valid group assignment failed: %v", err)
	}
	got, err := healthy.Members(context.Background(), "group-success")
	if err != nil || len(got) != 1 || got[0].EnrollmentID != "enrollment-success" {
		t.Fatalf("valid assignment was not persisted: members=%+v err=%v", got, err)
	}
}
