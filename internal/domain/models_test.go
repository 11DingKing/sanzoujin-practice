package domain

import (
	"testing"
	"time"
)

func TestProjectTransitions(t *testing.T) {
	cases := []struct {
		name string
		from ProjectStatus
		to   ProjectStatus
		want bool
	}{
		{"draft publish", ProjectDraft, ProjectPublished, true},
		{"draft running", ProjectDraft, ProjectRunning, false},
		{"published running", ProjectPublished, ProjectRunning, true},
		{"published pause", ProjectPublished, ProjectPaused, true},
		{"published close", ProjectPublished, ProjectClosed, true},
		{"running pause", ProjectRunning, ProjectPaused, true},
		{"running close", ProjectRunning, ProjectClosed, true},
		{"paused run", ProjectPaused, ProjectRunning, true},
		{"paused close", ProjectPaused, ProjectClosed, true},
		{"closed draft", ProjectClosed, ProjectDraft, false},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			p := Project{Status: tc.from}
			if got := p.CanTransition(tc.to); got != tc.want {
				t.Fatalf("got %v want %v", got, tc.want)
			}
		})
	}
}

func TestEnrollmentTransitions(t *testing.T) {
	cases := []struct {
		name     string
		from, to EnrollmentStatus
		want     bool
	}{
		{"pending authorized", EnrollmentPending, EnrollmentAuthorized, true},
		{"pending wait", EnrollmentPending, EnrollmentWaitlisted, true},
		{"pending withdraw", EnrollmentPending, EnrollmentWithdrawn, true},
		{"authorized matched", EnrollmentAuthorized, EnrollmentMatched, true},
		{"authorized withdraw", EnrollmentAuthorized, EnrollmentWithdrawn, true},
		{"wait authorized", EnrollmentWaitlisted, EnrollmentAuthorized, true},
		{"matched complete", EnrollmentMatched, EnrollmentCompleted, true},
		{"matched withdraw", EnrollmentMatched, EnrollmentWithdrawn, true},
		{"completed pending", EnrollmentCompleted, EnrollmentPending, false},
		{"withdrawn matched", EnrollmentWithdrawn, EnrollmentMatched, false},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := (Enrollment{Status: tc.from}).CanTransition(tc.to); got != tc.want {
				t.Fatalf("got %v", got)
			}
		})
	}
}

func TestGroupTransitions(t *testing.T) {
	cases := []struct {
		from, to GroupStatus
		want     bool
	}{{GroupForming, GroupReady, true}, {GroupForming, GroupDeparted, false}, {GroupReady, GroupDeparted, true}, {GroupReady, GroupFinished, false}, {GroupDeparted, GroupFinished, true}, {GroupFinished, GroupReady, false}}
	for _, tc := range cases {
		if got := (Group{Status: tc.from}).CanTransition(tc.to); got != tc.want {
			t.Errorf("%s to %s got %v", tc.from, tc.to, got)
		}
	}
}
func TestRiskTransitions(t *testing.T) {
	cases := []struct {
		from, to RiskStatus
		want     bool
	}{{RiskOpen, RiskInvestigating, true}, {RiskOpen, RiskCancelled, true}, {RiskInvestigating, RiskResolved, true}, {RiskInvestigating, RiskOpen, true}, {RiskResolved, RiskOpen, false}, {RiskCancelled, RiskResolved, false}}
	for _, tc := range cases {
		if got := (RiskEvent{Status: tc.from}).CanTransition(tc.to); got != tc.want {
			t.Errorf("risk transition %s/%s", tc.from, tc.to)
		}
	}
}
func TestSubmissionTransitions(t *testing.T) {
	cases := []struct {
		from, to SubmissionStatus
		want     bool
	}{{SubmissionDraft, SubmissionSubmitted, true}, {SubmissionSubmitted, SubmissionApproved, true}, {SubmissionSubmitted, SubmissionRejected, true}, {SubmissionRejected, SubmissionSubmitted, true}, {SubmissionApproved, SubmissionRejected, false}}
	for _, tc := range cases {
		if got := (Submission{Status: tc.from}).CanTransition(tc.to); got != tc.want {
			t.Errorf("submission transition")
		}
	}
}
func TestWindowAndPolicy(t *testing.T) {
	now := time.Date(2026, 8, 1, 10, 0, 0, 0, time.UTC)
	if !WithinWindow(now, now.Add(-time.Hour), now.Add(time.Hour)) {
		t.Fatal("window")
	}
	if WithinWindow(now, now.Add(time.Hour), now.Add(2*time.Hour)) {
		t.Fatal("early")
	}
	if CanManageProject(RoleStudent) {
		t.Fatal("student manage")
	}
	if !CanManageProject(RoleCoordinator) {
		t.Fatal("coordinator")
	}
	if !CanReview(RoleMentor) {
		t.Fatal("mentor review")
	}
	if CanAttend(RoleGuardian) {
		t.Fatal("guardian attend")
	}
	for _, score := range []int{1, 2, 3, 4, 5} {
		if _, e := NormalizeScore(score); e != nil {
			t.Fatal(e)
		}
	}
	if _, e := NormalizeScore(0); e == nil {
		t.Fatal("zero score")
	}
	if _, e := NormalizeScore(6); e == nil {
		t.Fatal("high score")
	}
}

func TestErrorsAreSentinels(t *testing.T) {
	if ErrNotFound == nil || ErrConflict == nil || ErrForbidden == nil || ErrInvalid == nil || ErrCapacity == nil || ErrIdempotency == nil || ErrExpired == nil {
		t.Fatal("sentinel missing")
	}
	if !IsTerminal("sent") || !IsTerminal("permanent_failed") || IsTerminal("pending") {
		t.Fatal("terminal classification")
	}
}
