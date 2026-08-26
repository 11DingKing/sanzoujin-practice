package service

import (
	"context"
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/11DingKing/sanzoujin-practice/internal/audit"
	"github.com/11DingKing/sanzoujin-practice/internal/domain"
	"github.com/11DingKing/sanzoujin-practice/internal/repository"
	"github.com/11DingKing/sanzoujin-practice/internal/storage"
	"time"
)

type EnrollmentService struct {
	DB          *storage.DB
	Projects    repository.ProjectRepo
	Enrollments repository.EnrollmentRepo
	Audit       audit.Service
	Idempotency repository.IdempotencyRepo
}

func (s EnrollmentService) Enroll(ctx context.Context, e domain.Enrollment, method, path, requestKey string) (domain.Enrollment, error) {
	if e.StudentID == "" || e.GuardianID == "" || e.ProjectID == "" {
		return e, domain.ErrInvalid
	}
	if requestKey != "" {
		if cached, err := s.Idempotency.Get(ctx, requestKey, method, path); err == nil {
			_ = json.Unmarshal([]byte(cached), &e)
			return e, nil
		}
	}
	p, err := s.Projects.ByID(ctx, e.ProjectID)
	if err != nil {
		return e, err
	}
	if p.Status != domain.ProjectPublished || !domain.WithinWindow(time.Now(), p.StartsAt, p.EndsAt) {
		return e, domain.ErrConflict
	}
	if old, err := s.Enrollments.ByStudentProject(ctx, e.StudentID, e.ProjectID); err == nil {
		return old, nil
	}
	e.ID = randomID()
	e.CreatedAt = time.Now()
	e.UpdatedAt = e.CreatedAt
	e.Status = domain.EnrollmentAuthorized
	if e.IdempotencyKey == "" {
		e.IdempotencyKey = requestKey
	}
	err = s.DB.Tx(ctx, func(tx *sql.Tx) error {
		if err := s.Projects.Reserve(ctx, tx, e.ProjectID); err != nil {
			return err
		}
		if err := s.Enrollments.Create(ctx, e); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return e, fmt.Errorf("enroll transaction: %w", err)
	}
	if err = s.Audit.Record(ctx, p.OrgID, e.StudentID, "enrollment", e.ID, "enroll", "ok", ""); err != nil {
		return e, fmt.Errorf("enrollment audit: %w", err)
	}
	if requestKey != "" {
		b, _ := json.Marshal(e)
		h := sha256.Sum256(b)
		_ = s.Idempotency.Put(ctx, requestKey, method, path, hex.EncodeToString(h[:]), string(b))
	}
	return e, nil
}
func (s EnrollmentService) Authorize(ctx context.Context, id string, guardian, org string) (domain.Enrollment, error) {
	e, err := s.Enrollments.ByID(ctx, id)
	if err != nil {
		return e, err
	}
	if e.GuardianID != guardian || !e.CanTransition(domain.EnrollmentAuthorized) {
		return e, domain.ErrForbidden
	}
	err = s.Enrollments.SetStatus(ctx, id, domain.EnrollmentAuthorized)
	if err == nil {
		err = s.Audit.Record(ctx, org, guardian, "enrollment", id, "authorize", "ok", "")
	}
	return e, err
}
func (s EnrollmentService) List(ctx context.Context, pid string) ([]domain.Enrollment, error) {
	return s.Enrollments.ListByProject(ctx, pid)
}
