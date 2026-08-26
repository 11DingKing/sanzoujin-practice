package service

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/11DingKing/sanzoujin-practice/internal/audit"
	"github.com/11DingKing/sanzoujin-practice/internal/domain"
	"github.com/11DingKing/sanzoujin-practice/internal/repository"
	"github.com/11DingKing/sanzoujin-practice/internal/storage"
	"time"
)

type GroupService struct {
	DB          *storage.DB
	Groups      repository.GroupRepo
	Enrollments repository.EnrollmentRepo
	Audit       audit.Service
}

func (s GroupService) Create(ctx context.Context, g domain.Group, org, actor string) error {
	if g.Capacity <= 0 || g.ProjectID == "" {
		return domain.ErrInvalid
	}
	if g.Status == "" {
		g.Status = domain.GroupForming
	}
	if g.Version == 0 {
		g.Version = 1
	}
	if g.CreatedAt.IsZero() {
		g.CreatedAt = time.Now()
	}
	if e := s.Groups.Create(ctx, g); e != nil {
		return e
	}
	return s.Audit.Record(ctx, org, actor, "group", g.ID, "create", "ok", "")
}
func (s GroupService) Add(ctx context.Context, gid, eid, org, actor string) error {
	e, err := s.Enrollments.ByID(ctx, eid)
	if err != nil {
		return err
	}
	if e.Status != domain.EnrollmentAuthorized {
		return domain.ErrConflict
	}
	return s.DB.Tx(ctx, func(tx *sql.Tx) error {
		if err := s.Groups.AddMember(ctx, tx, gid, eid); err != nil {
			return fmt.Errorf("group capacity: %w", err)
		}
		if _, err := tx.ExecContext(ctx, `UPDATE enrollments SET status='matched',updated_at=? WHERE id=? AND status='authorized'`, time.Now(), eid); err != nil {
			return err
		}
		return s.Audit.Record(ctx, org, actor, "group", gid, "add_member", "ok", "")
	})
}
func (s GroupService) Members(ctx context.Context, gid string) ([]domain.GroupMember, error) {
	return s.Groups.Members(ctx, gid)
}
