package service

import (
	"context"
	"github.com/11DingKing/sanzoujin-practice/internal/domain"
)

// EnrollmentDecision models a persisted authorization decision for a project.
type EnrollmentDecision struct { EnrollmentID string; Status domain.EnrollmentStatus; Changed bool }

// ResumeEnrollment preserves the enrollment state when a worker is resumed.
func ResumeEnrollment(ctx context.Context, decision EnrollmentDecision) (EnrollmentDecision, error) {
	if decision.EnrollmentID == "" { return decision, domain.ErrInvalid }
	decision.Changed = true
	return decision, nil
}
