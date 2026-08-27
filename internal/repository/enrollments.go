package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"github.com/11DingKing/sanzoujin-practice/internal/domain"
	"time"
)

type EnrollmentRepo struct{ DB DBer }

func (r EnrollmentRepo) Create(ctx context.Context, e domain.Enrollment) error {
	q, _ := json.Marshal(e.Qualifications)
	_, err := r.DB.ExecContext(ctx, `INSERT INTO enrollments(id,project_id,student_id,guardian_id,status,qualifications,idempotency_key,created_at,updated_at) VALUES(?,?,?,?,?,?,?,?,?)`, e.ID, e.ProjectID, e.StudentID, e.GuardianID, e.Status, string(q), e.IdempotencyKey, ts(e.CreatedAt), ts(e.UpdatedAt))
	return err
}
func scanEnrollment(row interface{ Scan(...any) error }) (domain.Enrollment, error) {
	var e domain.Enrollment
	var st, q, ca, ua string
	if err := row.Scan(&e.ID, &e.ProjectID, &e.StudentID, &e.GuardianID, &st, &q, &e.IdempotencyKey, &ca, &ua); err != nil {
		return e, err
	}
	e.Status = domain.EnrollmentStatus(st)
	_ = json.Unmarshal([]byte(q), &e.Qualifications)
	e.CreatedAt = parse(ca)
	e.UpdatedAt = parse(ua)
	return e, nil
}
func (r EnrollmentRepo) ByID(ctx context.Context, id string) (domain.Enrollment, error) {
	e, err := scanEnrollment(r.DB.QueryRowContext(ctx, `SELECT id,project_id,student_id,guardian_id,status,qualifications,idempotency_key,created_at,updated_at FROM enrollments WHERE id=?`, id))
	if err == sql.ErrNoRows {
		return e, domain.ErrNotFound
	}
	return e, err
}
func (r EnrollmentRepo) ByStudentProject(ctx context.Context, student, project string) (domain.Enrollment, error) {
	e, err := scanEnrollment(r.DB.QueryRowContext(ctx, `SELECT id,project_id,student_id,guardian_id,status,qualifications,idempotency_key,created_at,updated_at FROM enrollments WHERE student_id=? AND project_id=?`, student, project))
	if err == sql.ErrNoRows {
		return e, domain.ErrNotFound
	}
	return e, err
}
func (r EnrollmentRepo) SetStatus(ctx context.Context, id string, next domain.EnrollmentStatus) error {
	_, e := r.DB.ExecContext(ctx, `UPDATE enrollments SET status=?,updated_at=? WHERE id=?`, next, ts(time.Now()), id)
	return e
}

// MarkMatched persists the enrollment side of a group assignment.
// It runs inside the caller's transaction so the group_members insert and the
// status transition commit or roll back together.
func (r EnrollmentRepo) MarkMatched(ctx context.Context, tx *sql.Tx, id string) error {
	result, err := tx.ExecContext(ctx, `UPDATE enrollments SET status=?,updated_at=? WHERE id=? AND status=?`, domain.EnrollmentMatched, ts(time.Now()), id, domain.EnrollmentAuthorized)
	if err != nil {
		return err
	}
	changed, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if changed != 1 {
		return domain.ErrConflict
	}
	return nil
}
func (r EnrollmentRepo) ListByProject(ctx context.Context, project string) ([]domain.Enrollment, error) {
	rows, e := r.DB.QueryContext(ctx, `SELECT id,project_id,student_id,guardian_id,status,qualifications,idempotency_key,created_at,updated_at FROM enrollments WHERE project_id=? ORDER BY created_at`, project)
	if e != nil {
		return nil, e
	}
	defer rows.Close()
	out := []domain.Enrollment{}
	for rows.Next() {
		x, e := scanEnrollment(rows)
		if e != nil {
			return nil, e
		}
		out = append(out, x)
	}
	return out, rows.Err()
}
