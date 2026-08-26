package service

import (
	"context"
	"github.com/11DingKing/sanzoujin-practice/internal/audit"
	"github.com/11DingKing/sanzoujin-practice/internal/domain"
	"github.com/11DingKing/sanzoujin-practice/internal/repository"
	"time"
)

type SubmissionService struct {
	Repo  repository.SubmissionRepo
	Audit audit.Service
}

func (s SubmissionService) Create(ctx context.Context, x domain.Submission, org, actor string) error {
	if x.ProjectID == "" || x.StudentID == "" || x.Content == "" {
		return domain.ErrInvalid
	}
	x.ID = randomID()
	x.Status = domain.SubmissionDraft
	x.Version = 1
	if err := s.Repo.Create(ctx, x); err != nil {
		return err
	}
	return s.Audit.Record(ctx, org, actor, "submission", x.ID, "create", "ok", "")
}
func (s SubmissionService) Submit(ctx context.Context, id string, version int, org, actor string) error {
	if err := s.Repo.Transition(ctx, id, domain.SubmissionSubmitted, version, ""); err != nil {
		return err
	}
	return s.Audit.Record(ctx, org, actor, "submission", id, "submit", "ok", "")
}
func (s SubmissionService) Review(ctx context.Context, id string, next domain.SubmissionStatus, version int, reviewer, org string) error {
	if next != domain.SubmissionApproved && next != domain.SubmissionRejected {
		return domain.ErrInvalid
	}
	if err := s.Repo.Transition(ctx, id, next, version, reviewer); err != nil {
		return err
	}
	return s.Audit.Record(ctx, org, reviewer, "submission", id, "review", "ok", "")
}
func submissionDeadline(start time.Time) time.Time { return start.Add(30 * 24 * time.Hour) }
