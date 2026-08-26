package service

import (
    "context"
    "testing"
    "github.com/11DingKing/sanzoujin-practice/internal/domain"
)

func TestCancelledEnrollmentResumePreservesState(t *testing.T) {
    ctx, cancel := context.WithCancel(context.Background()); cancel()
    original := EnrollmentResume{EnrollmentID:"enroll-01", Status:domain.EnrollmentAuthorized}
    got, err := ResumeEnrollmentT01(ctx, original)
    if err == nil { t.Fatalf("cancelled resume succeeded: %+v", got) }
    if got.Changed { t.Fatalf("cancelled resume changed state: %+v", got) }
}
