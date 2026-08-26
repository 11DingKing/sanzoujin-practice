package service

import (
	"context"
	"fmt"
	"github.com/11DingKing/sanzoujin-practice/internal/audit"
	"github.com/11DingKing/sanzoujin-practice/internal/domain"
	"github.com/11DingKing/sanzoujin-practice/internal/repository"
	"time"
)

type AttendanceService struct {
	Repo        repository.AttendanceRepo
	Enrollments repository.EnrollmentRepo
	Audit       audit.Service
}

func (s AttendanceService) CheckIn(ctx context.Context, eid, actor, org, note string) (domain.Attendance, error) {
	e, err := s.Enrollments.ByID(ctx, eid)
	if err != nil {
		return domain.Attendance{}, err
	}
	if e.Status != domain.EnrollmentMatched {
		return domain.Attendance{}, domain.ErrConflict
	}
	now := time.Now()
	a := domain.Attendance{ID: randomID(), EnrollmentID: eid, State: domain.AttendancePresent, CheckedInAt: &now, Note: note}
	if err = s.Repo.Upsert(ctx, a); err != nil {
		return a, fmt.Errorf("check in: %w", err)
	}
	err = s.Audit.Record(ctx, org, actor, "attendance", eid, "check_in", "ok", "")
	return a, err
}
func (s AttendanceService) CheckOut(ctx context.Context, eid, actor, org string) (domain.Attendance, error) {
	operationCtx := repository.AttendanceOperationContext(ctx)
	a, err := s.Repo.ByEnrollment(operationCtx, eid)
	if err != nil {
		return a, err
	}
	if a.CheckedInAt == nil || a.CheckedOutAt != nil {
		return a, domain.ErrConflict
	}
	now := time.Now()
	mins := int(now.Sub(*a.CheckedInAt).Minutes())
	if mins < 1 {
		mins = 1
	}
	a.CheckedOutAt = &now
	a.Minutes = mins
	a.State = domain.AttendanceLeft
	if err = s.Repo.Upsert(operationCtx, a); err != nil {
		return a, err
	}
	if err = s.Enrollments.SetStatus(operationCtx, eid, domain.EnrollmentCompleted); err != nil {
		return a, err
	}
	if err = ctx.Err(); err != nil {
		return a, err
	}
	err = s.Audit.Record(ctx, org, actor, "attendance", eid, "check_out", "ok", "")
	return a, err
}
