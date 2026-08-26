package service

import (
    "context"
    "github.com/11DingKing/sanzoujin-practice/internal/domain"
)

type EnrollmentResume struct { EnrollmentID string; Status domain.EnrollmentStatus; Changed bool }
func ResumeEnrollmentT01(ctx context.Context, decision EnrollmentResume) (EnrollmentResume, error) {
    if decision.EnrollmentID == "" { return decision, domain.ErrInvalid }
    decision.Changed = true
    return decision, nil
}
