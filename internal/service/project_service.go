package service

import (
	"context"
	"fmt"
	"github.com/11DingKing/sanzoujin-practice/internal/audit"
	"github.com/11DingKing/sanzoujin-practice/internal/domain"
	"github.com/11DingKing/sanzoujin-practice/internal/repository"
	"github.com/11DingKing/sanzoujin-practice/internal/storage"
	"time"
)

type ProjectService struct {
	DB       *storage.DB
	Projects repository.ProjectRepo
	Venues   repository.VenueRepo
	Audit    audit.Service
}

func (s ProjectService) Create(ctx context.Context, p domain.Project) error {
	if p.Title == "" || p.Capacity <= 0 || !p.EndsAt.After(p.StartsAt) {
		return domain.ErrInvalid
	}
	if p.Status == "" {
		p.Status = domain.ProjectDraft
	}
	if p.Version == 0 {
		p.Version = 1
	}
	if p.CreatedAt.IsZero() {
		p.CreatedAt = time.Now()
	}
	p.UpdatedAt = p.CreatedAt
	if e := s.Projects.Create(ctx, p); e != nil {
		return fmt.Errorf("create project: %w", e)
	}
	return s.Audit.Record(ctx, p.OrgID, p.CreatedBy, "project", p.ID, "create", "ok", "")
}
func (s ProjectService) Publish(ctx context.Context, id string, version int, actor, org string) error {
	p, e := s.Projects.ByID(ctx, id)
	if e != nil {
		return e
	}
	if p.Status != domain.ProjectDraft {
		return domain.ErrConflict
	}
	if e = s.Projects.Transition(ctx, id, domain.ProjectPublished, version, time.Now().UTC().Format(time.RFC3339Nano)); e != nil {
		return e
	}
	return s.Audit.Record(ctx, org, actor, "project", id, "publish", "ok", "")
}
func (s ProjectService) ChangeStatus(ctx context.Context, id string, next domain.ProjectStatus, version int, actor, org string) error {
	p, e := s.Projects.ByID(ctx, id)
	if e != nil {
		return e
	}
	if !p.CanTransition(next) {
		return domain.ErrConflict
	}
	if e = s.Projects.Transition(ctx, id, next, version, time.Now().UTC().Format(time.RFC3339Nano)); e != nil {
		return e
	}
	return s.Audit.Record(ctx, org, actor, "project", id, "status", "ok", "")
}
func (s ProjectService) List(ctx context.Context, status string, limit int) ([]domain.Project, error) {
	return s.Projects.List(ctx, status, limit)
}
