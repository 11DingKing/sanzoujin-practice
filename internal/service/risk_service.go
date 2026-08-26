package service

import (
	"context"
	"github.com/11DingKing/sanzoujin-practice/internal/audit"
	"github.com/11DingKing/sanzoujin-practice/internal/domain"
	"github.com/11DingKing/sanzoujin-practice/internal/repository"
	"time"
)

type RiskService struct {
	Repo     repository.RiskRepo
	Projects repository.ProjectRepo
	Audit    audit.Service
}

func (s RiskService) Report(ctx context.Context, r domain.RiskEvent, org, actor string) error {
	if r.Severity < 1 || r.Severity > 5 || r.Description == "" {
		return domain.ErrInvalid
	}
	r.ID = randomID()
	r.ReporterID = actor
	r.Status = domain.RiskOpen
	r.Version = 1
	r.CreatedAt = time.Now()
	if err := s.Repo.Create(ctx, r); err != nil {
		return err
	}
	if r.Severity >= 4 {
		_ = s.Projects.Transition(ctx, r.ProjectID, domain.ProjectPaused, 1, time.Now().UTC().Format(time.RFC3339Nano))
	}
	return s.Audit.Record(ctx, org, actor, "risk", r.ID, "report", "ok", "")
}
func (s RiskService) Resolve(ctx context.Context, id string, version int, org, actor string) error {
	if err := s.Repo.Transition(ctx, id, domain.RiskResolved, version); err != nil {
		return err
	}
	return s.Audit.Record(ctx, org, actor, "risk", id, "resolve", "ok", "")
}
