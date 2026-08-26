package repository

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/11DingKing/sanzoujin-practice/internal/domain"
	"time"
)

type GroupRepo struct{ DB DBer }

func (r GroupRepo) Create(ctx context.Context, g domain.Group) error {
	_, e := r.DB.ExecContext(ctx, `INSERT INTO groups(id,project_id,name,capacity,status,mentor_id,version,created_at) VALUES(?,?,?,?,?,?,?,?)`, g.ID, g.ProjectID, g.Name, g.Capacity, g.Status, g.MentorID, g.Version, ts(g.CreatedAt))
	return e
}
func (r GroupRepo) ByID(ctx context.Context, id string) (domain.Group, error) {
	var g domain.Group
	var st, ca string
	var mentor sql.NullString
	e := r.DB.QueryRowContext(ctx, `SELECT id,project_id,name,capacity,status,mentor_id,version,created_at FROM groups WHERE id=?`, id).Scan(&g.ID, &g.ProjectID, &g.Name, &g.Capacity, &st, &mentor, &g.Version, &ca)
	if e == sql.ErrNoRows {
		return g, domain.ErrNotFound
	}
	if e != nil {
		return g, e
	}
	g.Status = domain.GroupStatus(st)
	g.MentorID = mentor.String
	g.CreatedAt = parse(ca)
	return g, nil
}
func (r GroupRepo) AddMember(ctx context.Context, tx *sql.Tx, gid, eid string) error {
	var cap, used int
	if err := tx.QueryRowContext(ctx, `SELECT capacity, (SELECT COUNT(*) FROM group_members WHERE group_id=?) FROM groups WHERE id=?`, gid, gid).Scan(&cap, &used); err != nil {
		return err
	}
	if used >= cap {
		return domain.ErrCapacity
	}
	_, err := tx.ExecContext(ctx, `INSERT INTO group_members(group_id,enrollment_id,joined_at) VALUES(?,?,?)`, gid, eid, ts(time.Now()))
	return err
}
func (r GroupRepo) Members(ctx context.Context, gid string) ([]domain.GroupMember, error) {
	rows, e := r.DB.QueryContext(ctx, `SELECT group_id,enrollment_id,joined_at FROM group_members WHERE group_id=? ORDER BY joined_at`, gid)
	if e != nil {
		return nil, e
	}
	defer rows.Close()
	out := []domain.GroupMember{}
	for rows.Next() {
		var m domain.GroupMember
		var at string
		if e := rows.Scan(&m.GroupID, &m.EnrollmentID, &at); e != nil {
			return nil, e
		}
		m.JoinedAt = parse(at)
		out = append(out, m)
	}
	return out, rows.Err()
}

type AttendanceRepo struct{ DB DBer }

func (r AttendanceRepo) Upsert(ctx context.Context, a domain.Attendance) error {
	_, e := r.DB.ExecContext(ctx, `INSERT INTO attendance(id,enrollment_id,state,checked_in_at,checked_out_at,minutes,note) VALUES(?,?,?,?,?,?,?) ON CONFLICT(enrollment_id) DO UPDATE SET state=excluded.state,checked_in_at=excluded.checked_in_at,checked_out_at=excluded.checked_out_at,minutes=excluded.minutes,note=excluded.note`, a.ID, a.EnrollmentID, a.State, timeString(a.CheckedInAt), timeString(a.CheckedOutAt), a.Minutes, a.Note)
	return e
}
func timeString(t *time.Time) any {
	if t == nil {
		return nil
	}
	return ts(*t)
}
func (r AttendanceRepo) ByEnrollment(ctx context.Context, eid string) (domain.Attendance, error) {
	var a domain.Attendance
	var ci, co sql.NullString
	e := r.DB.QueryRowContext(ctx, `SELECT id,enrollment_id,state,checked_in_at,checked_out_at,minutes,note FROM attendance WHERE enrollment_id=?`, eid).Scan(&a.ID, &a.EnrollmentID, &a.State, &ci, &co, &a.Minutes, &a.Note)
	if e == sql.ErrNoRows {
		return a, domain.ErrNotFound
	}
	if e != nil {
		return a, fmt.Errorf("attendance: %w", e)
	}
	a.CheckedInAt = nullTime(ci)
	a.CheckedOutAt = nullTime(co)
	return a, nil
}

type RiskRepo struct{ DB DBer }

func (r RiskRepo) Create(ctx context.Context, x domain.RiskEvent) error {
	_, e := r.DB.ExecContext(ctx, `INSERT INTO risk_events(id,project_id,reporter_id,severity,description,status,version,created_at) VALUES(?,?,?,?,?,?,?,?)`, x.ID, x.ProjectID, x.ReporterID, x.Severity, x.Description, x.Status, x.Version, ts(x.CreatedAt))
	return e
}
func (r RiskRepo) Transition(ctx context.Context, id string, next domain.RiskStatus, version int) error {
	res, e := r.DB.ExecContext(ctx, `UPDATE risk_events SET status=?,version=version+1,resolved_at=CASE WHEN ?=? THEN ? ELSE resolved_at END WHERE id=? AND version=?`, next, next, domain.RiskResolved, ts(time.Now()), id, version)
	if e != nil {
		return e
	}
	n, _ := res.RowsAffected()
	if n == 0 {
		return domain.ErrConflict
	}
	return nil
}

type SubmissionRepo struct{ DB DBer }

func (r SubmissionRepo) Create(ctx context.Context, s domain.Submission) error {
	_, e := r.DB.ExecContext(ctx, `INSERT INTO submissions(id,project_id,student_id,version,status,content) VALUES(?,?,?,?,?,?)`, s.ID, s.ProjectID, s.StudentID, s.Version, s.Status, s.Content)
	return e
}
func (r SubmissionRepo) Transition(ctx context.Context, id string, next domain.SubmissionStatus, version int, reviewer string) error {
	res, e := r.DB.ExecContext(ctx, `UPDATE submissions SET status=?,version=version+1,reviewed_at=CASE WHEN ? IN ('approved','rejected') THEN ? ELSE reviewed_at END,reviewer_id=CASE WHEN ? IN ('approved','rejected') THEN ? ELSE reviewer_id END WHERE id=? AND version=?`, next, next, ts(time.Now()), next, reviewer, id, version)
	if e != nil {
		return e
	}
	n, _ := res.RowsAffected()
	if n == 0 {
		return domain.ErrConflict
	}
	return nil
}
