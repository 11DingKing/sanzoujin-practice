package storage

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
)

const schema = `CREATE TABLE IF NOT EXISTS schema_migrations(version INTEGER PRIMARY KEY, applied_at TEXT NOT NULL);
CREATE TABLE IF NOT EXISTS users(id TEXT PRIMARY KEY, org_id TEXT NOT NULL, name TEXT NOT NULL, email TEXT NOT NULL UNIQUE, password_hash TEXT NOT NULL, role TEXT NOT NULL, active INTEGER NOT NULL DEFAULT 1, created_at TEXT NOT NULL);
CREATE TABLE IF NOT EXISTS sessions(id TEXT PRIMARY KEY, user_id TEXT NOT NULL REFERENCES users(id), token_hash TEXT NOT NULL UNIQUE, expires_at TEXT NOT NULL, revoked_at TEXT, created_at TEXT NOT NULL);
CREATE TABLE IF NOT EXISTS venues(id TEXT PRIMARY KEY, name TEXT NOT NULL, address TEXT NOT NULL, capacity INTEGER NOT NULL, open INTEGER NOT NULL DEFAULT 1, created_at TEXT NOT NULL);
CREATE TABLE IF NOT EXISTS practice_projects(id TEXT PRIMARY KEY, org_id TEXT NOT NULL, title TEXT NOT NULL, description TEXT NOT NULL, venue_id TEXT NOT NULL REFERENCES venues(id), capacity INTEGER NOT NULL, capacity_used INTEGER NOT NULL DEFAULT 0, starts_at TEXT NOT NULL, ends_at TEXT NOT NULL, status TEXT NOT NULL, version INTEGER NOT NULL DEFAULT 1, created_by TEXT NOT NULL REFERENCES users(id), created_at TEXT NOT NULL, updated_at TEXT NOT NULL);
CREATE TABLE IF NOT EXISTS enrollments(id TEXT PRIMARY KEY, project_id TEXT NOT NULL REFERENCES practice_projects(id), student_id TEXT NOT NULL REFERENCES users(id), guardian_id TEXT NOT NULL REFERENCES users(id), status TEXT NOT NULL, qualifications TEXT NOT NULL, idempotency_key TEXT NOT NULL, created_at TEXT NOT NULL, updated_at TEXT NOT NULL, UNIQUE(project_id, student_id));
CREATE TABLE IF NOT EXISTS groups(id TEXT PRIMARY KEY, project_id TEXT NOT NULL REFERENCES practice_projects(id), name TEXT NOT NULL, capacity INTEGER NOT NULL, status TEXT NOT NULL, mentor_id TEXT REFERENCES users(id), version INTEGER NOT NULL DEFAULT 1, created_at TEXT NOT NULL);
CREATE TABLE IF NOT EXISTS group_members(group_id TEXT NOT NULL REFERENCES groups(id), enrollment_id TEXT NOT NULL REFERENCES enrollments(id), joined_at TEXT NOT NULL, PRIMARY KEY(group_id, enrollment_id), UNIQUE(enrollment_id));
CREATE TABLE IF NOT EXISTS attendance(id TEXT PRIMARY KEY, enrollment_id TEXT NOT NULL REFERENCES enrollments(id), state TEXT NOT NULL, checked_in_at TEXT, checked_out_at TEXT, minutes INTEGER NOT NULL DEFAULT 0, note TEXT NOT NULL DEFAULT '', UNIQUE(enrollment_id));
CREATE TABLE IF NOT EXISTS risk_events(id TEXT PRIMARY KEY, project_id TEXT NOT NULL REFERENCES practice_projects(id), reporter_id TEXT NOT NULL REFERENCES users(id), severity INTEGER NOT NULL, description TEXT NOT NULL, status TEXT NOT NULL, version INTEGER NOT NULL DEFAULT 1, created_at TEXT NOT NULL, resolved_at TEXT);
CREATE TABLE IF NOT EXISTS submissions(id TEXT PRIMARY KEY, project_id TEXT NOT NULL REFERENCES practice_projects(id), student_id TEXT NOT NULL REFERENCES users(id), version INTEGER NOT NULL DEFAULT 1, status TEXT NOT NULL, content TEXT NOT NULL, submitted_at TEXT, reviewed_at TEXT, reviewer_id TEXT REFERENCES users(id));
CREATE TABLE IF NOT EXISTS evaluations(id TEXT PRIMARY KEY, submission_id TEXT NOT NULL REFERENCES submissions(id), evaluator_id TEXT NOT NULL REFERENCES users(id), role TEXT NOT NULL, score INTEGER NOT NULL, comment TEXT NOT NULL, created_at TEXT NOT NULL, UNIQUE(submission_id, evaluator_id));
CREATE TABLE IF NOT EXISTS idempotency_keys(key TEXT NOT NULL, method TEXT NOT NULL, path TEXT NOT NULL, request_hash TEXT NOT NULL, response TEXT NOT NULL, created_at TEXT NOT NULL, PRIMARY KEY(key, method, path));
CREATE TABLE IF NOT EXISTS audit_events(id TEXT PRIMARY KEY, org_id TEXT NOT NULL, actor_id TEXT NOT NULL, object_type TEXT NOT NULL, object_id TEXT NOT NULL, action TEXT NOT NULL, result TEXT NOT NULL, request_id TEXT NOT NULL, prev_hash TEXT NOT NULL, hash TEXT NOT NULL, created_at TEXT NOT NULL);
CREATE TABLE IF NOT EXISTS outbox_messages(id TEXT PRIMARY KEY, topic TEXT NOT NULL, payload TEXT NOT NULL, attempts INTEGER NOT NULL DEFAULT 0, next_attempt_at TEXT NOT NULL, sent_at TEXT, last_error TEXT NOT NULL DEFAULT '', created_at TEXT NOT NULL);
CREATE TABLE IF NOT EXISTS worker_jobs(id TEXT PRIMARY KEY, kind TEXT NOT NULL, payload TEXT NOT NULL, status TEXT NOT NULL, attempts INTEGER NOT NULL DEFAULT 0, run_after TEXT NOT NULL, last_error TEXT NOT NULL DEFAULT '', created_at TEXT NOT NULL, updated_at TEXT NOT NULL);
CREATE INDEX IF NOT EXISTS idx_projects_status ON practice_projects(status, starts_at); CREATE INDEX IF NOT EXISTS idx_enrollments_project ON enrollments(project_id, status); CREATE INDEX IF NOT EXISTS idx_audit_object ON audit_events(object_type, object_id, created_at);`

func (d *DB) Migrate(ctx context.Context) error {
	return d.Tx(ctx, func(tx *sql.Tx) error {
		if _, err := tx.ExecContext(ctx, `CREATE TABLE IF NOT EXISTS schema_migrations(version INTEGER PRIMARY KEY, applied_at TEXT NOT NULL)`); err != nil {
			return err
		}
		var version int
		_ = tx.QueryRowContext(ctx, `SELECT COALESCE(MAX(version),0) FROM schema_migrations`).Scan(&version)
		if version >= 1 {
			return nil
		}
		for _, statement := range strings.Split(schema, ";") {
			stmt := strings.TrimSpace(statement)
			if stmt == "" {
				continue
			}
			if _, err := tx.ExecContext(ctx, stmt); err != nil {
				return fmt.Errorf("migration statement: %w", err)
			}
		}
		if _, err := tx.ExecContext(ctx, `INSERT INTO schema_migrations(version, applied_at) VALUES(1, datetime('now'))`); err != nil {
			return err
		}
		return nil
	})
}
