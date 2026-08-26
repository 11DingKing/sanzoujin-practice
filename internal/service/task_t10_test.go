package service

import (
    "context"
    "testing"
    "github.com/11DingKing/sanzoujin-practice/internal/domain"
)

func TestCancelledResumeDoesNotMutateEnrollment(t *testing.T) {
    ctx, cancel := context.WithCancel(context.Background())
    cancel()
    original := EnrollmentDecision{EnrollmentID: "enroll-10", Status: domain.EnrollmentAuthorized}
    got, err := ResumeEnrollment(ctx, original)
    if err == nil { t.Fatalf("cancelled resume returned success: %+v", got) }
    if got.Changed { t.Fatalf("cancelled resume mutated persisted decision: %+v", got) }
}
