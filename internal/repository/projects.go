package repository

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/11DingKing/sanzoujin-practice/internal/domain"
	"time"
)

type ProjectRepo struct{ DB DBer }

func (r ProjectRepo) Create(ctx context.Context, p domain.Project) error {
	_, e := r.DB.ExecContext(ctx, `INSERT INTO practice_projects(id,org_id,title,description,venue_id,capacity,starts_at,ends_at,status,version,created_by,created_at,updated_at) VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?)`, p.ID, p.OrgID, p.Title, p.Description, p.VenueID, p.Capacity, ts(p.StartsAt), ts(p.EndsAt), p.Status, p.Version, p.CreatedBy, ts(p.CreatedAt), ts(p.UpdatedAt))
	return e
}
func scanProject(row interface{ Scan(...any) error }) (domain.Project, error) {
	var p domain.Project
	var st, sa, ea, ca, ua string
	e := row.Scan(&p.ID, &p.OrgID, &p.Title, &p.Description, &p.VenueID, &p.Capacity, &p.CapacityUsed, &sa, &ea, &st, &p.Version, &p.CreatedBy, &ca, &ua)
	if e != nil {
		return p, e
	}
	p.StartsAt = parse(sa)
	p.EndsAt = parse(ea)
	p.Status = domain.ProjectStatus(st)
	p.CreatedAt = parse(ca)
	p.UpdatedAt = parse(ua)
	return p, nil
}
func (r ProjectRepo) ByID(ctx context.Context, id string) (domain.Project, error) {
	p, e := scanProject(r.DB.QueryRowContext(ctx, `SELECT id,org_id,title,description,venue_id,capacity,capacity_used,starts_at,ends_at,status,version,created_by,created_at,updated_at FROM practice_projects WHERE id=?`, id))
	if e == sql.ErrNoRows {
		return p, domain.ErrNotFound
	}
	if e != nil {
		return p, fmt.Errorf("project: %w", e)
	}
	return p, nil
}
func (r ProjectRepo) List(ctx context.Context, status string, limit int) ([]domain.Project, error) {
	if limit <= 0 || limit > 100 {
		limit = 20
	}
	rows, e := r.DB.QueryContext(ctx, `SELECT id,org_id,title,description,venue_id,capacity,capacity_used,starts_at,ends_at,status,version,created_by,created_at,updated_at FROM practice_projects WHERE (?='' OR status=?) ORDER BY starts_at LIMIT ?`, status, status, limit)
	if e != nil {
		return nil, e
	}
	defer rows.Close()
	out := make([]domain.Project, 0)
	for rows.Next() {
		p, e := scanProject(rows)
		if e != nil {
			return nil, e
		}
		out = append(out, p)
	}
	return out, rows.Err()
}
func (r ProjectRepo) Transition(ctx context.Context, id string, next domain.ProjectStatus, version int, now string) error {
	res, e := r.DB.ExecContext(ctx, `UPDATE practice_projects SET status=?,version=version+1,updated_at=? WHERE id=? AND version=?`, next, now, id, version)
	if e != nil {
		return e
	}
	n, _ := res.RowsAffected()
	if n == 0 {
		return domain.ErrConflict
	}
	return nil
}
func (r ProjectRepo) Reserve(ctx context.Context, tx *sql.Tx, id string) error {
	res, e := tx.ExecContext(ctx, `UPDATE practice_projects SET capacity_used=capacity_used+1,updated_at=? WHERE id=? AND status=? AND capacity_used<capacity`, ts(time.Now()), id, domain.ProjectPublished)
	if e != nil {
		return e
	}
	n, _ := res.RowsAffected()
	if n == 0 {
		return domain.ErrCapacity
	}
	return nil
}
