package domain

import "time"

type Role string

const (
	RoleAdmin       Role = "admin"
	RoleCoordinator Role = "coordinator"
	RoleStudent     Role = "student"
	RoleGuardian    Role = "guardian"
	RoleMentor      Role = "mentor"
)

type ProjectStatus string

const (
	ProjectDraft     ProjectStatus = "draft"
	ProjectPublished ProjectStatus = "published"
	ProjectRunning   ProjectStatus = "running"
	ProjectPaused    ProjectStatus = "paused"
	ProjectClosed    ProjectStatus = "closed"
)

type EnrollmentStatus string

const (
	EnrollmentPending    EnrollmentStatus = "pending"
	EnrollmentAuthorized EnrollmentStatus = "authorized"
	EnrollmentWaitlisted EnrollmentStatus = "waitlisted"
	EnrollmentMatched    EnrollmentStatus = "matched"
	EnrollmentWithdrawn  EnrollmentStatus = "withdrawn"
	EnrollmentCompleted  EnrollmentStatus = "completed"
)

type GroupStatus string

const (
	GroupForming  GroupStatus = "forming"
	GroupReady    GroupStatus = "ready"
	GroupDeparted GroupStatus = "departed"
	GroupFinished GroupStatus = "finished"
)

type AttendanceState string

const (
	AttendanceAbsent  AttendanceState = "absent"
	AttendancePresent AttendanceState = "present"
	AttendanceLate    AttendanceState = "late"
	AttendanceLeft    AttendanceState = "left"
)

type RiskStatus string

const (
	RiskOpen          RiskStatus = "open"
	RiskInvestigating RiskStatus = "investigating"
	RiskResolved      RiskStatus = "resolved"
	RiskCancelled     RiskStatus = "cancelled"
)

type SubmissionStatus string

const (
	SubmissionDraft     SubmissionStatus = "draft"
	SubmissionSubmitted SubmissionStatus = "submitted"
	SubmissionApproved  SubmissionStatus = "approved"
	SubmissionRejected  SubmissionStatus = "rejected"
)

type User struct {
	ID           string
	OrgID        string
	Name         string
	Email        string
	PasswordHash string
	Role         Role
	Active       bool
	CreatedAt    time.Time
}
type Session struct {
	ID        string
	UserID    string
	TokenHash string
	ExpiresAt time.Time
	RevokedAt *time.Time
	CreatedAt time.Time
}
type Project struct {
	ID           string
	OrgID        string
	Title        string
	Description  string
	VenueID      string
	Capacity     int
	CapacityUsed int
	StartsAt     time.Time
	EndsAt       time.Time
	Status       ProjectStatus
	Version      int
	CreatedBy    string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
type Venue struct {
	ID        string
	Name      string
	Address   string
	Capacity  int
	Open      bool
	CreatedAt time.Time
}
type Enrollment struct {
	ID             string
	ProjectID      string
	StudentID      string
	GuardianID     string
	Status         EnrollmentStatus
	Qualifications []string
	IdempotencyKey string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
type Group struct {
	ID        string
	ProjectID string
	Name      string
	Capacity  int
	Status    GroupStatus
	MentorID  string
	Version   int
	CreatedAt time.Time
}
type GroupMember struct {
	GroupID      string
	EnrollmentID string
	JoinedAt     time.Time
}
type Attendance struct {
	ID           string
	EnrollmentID string
	State        AttendanceState
	CheckedInAt  *time.Time
	CheckedOutAt *time.Time
	Minutes      int
	Note         string
}
type RiskEvent struct {
	ID          string
	ProjectID   string
	ReporterID  string
	Severity    int
	Description string
	Status      RiskStatus
	Version     int
	CreatedAt   time.Time
	ResolvedAt  *time.Time
}
type Submission struct {
	ID          string
	ProjectID   string
	StudentID   string
	Version     int
	Status      SubmissionStatus
	Content     string
	SubmittedAt *time.Time
	ReviewedAt  *time.Time
	ReviewerID  string
}
type Evaluation struct {
	ID           string
	SubmissionID string
	EvaluatorID  string
	Role         Role
	Score        int
	Comment      string
	CreatedAt    time.Time
}
type AuditEvent struct {
	ID         string
	OrgID      string
	ActorID    string
	ObjectType string
	ObjectID   string
	Action     string
	Result     string
	RequestID  string
	PrevHash   string
	Hash       string
	CreatedAt  time.Time
}
type OutboxMessage struct {
	ID            string
	Topic         string
	Payload       string
	Attempts      int
	NextAttemptAt time.Time
	SentAt        *time.Time
	LastError     string
	CreatedAt     time.Time
}
type WorkerJob struct {
	ID        string
	Kind      string
	Payload   string
	Status    string
	Attempts  int
	RunAfter  time.Time
	LastError string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (p Project) CanTransition(next ProjectStatus) bool {
	switch p.Status {
	case ProjectDraft:
		return next == ProjectPublished
	case ProjectPublished:
		return next == ProjectRunning || next == ProjectPaused || next == ProjectClosed
	case ProjectRunning:
		return next == ProjectPaused || next == ProjectClosed
	case ProjectPaused:
		return next == ProjectRunning || next == ProjectClosed
	case ProjectClosed:
		return false
	}
	return false
}
func (e Enrollment) CanTransition(next EnrollmentStatus) bool {
	switch e.Status {
	case EnrollmentPending:
		return next == EnrollmentAuthorized || next == EnrollmentWaitlisted || next == EnrollmentWithdrawn
	case EnrollmentAuthorized:
		return next == EnrollmentMatched || next == EnrollmentWithdrawn
	case EnrollmentWaitlisted:
		return next == EnrollmentAuthorized || next == EnrollmentWithdrawn
	case EnrollmentMatched:
		return next == EnrollmentCompleted || next == EnrollmentWithdrawn
	case EnrollmentCompleted, EnrollmentWithdrawn:
		return false
	}
	return false
}
func (g Group) CanTransition(next GroupStatus) bool {
	switch g.Status {
	case GroupForming:
		return next == GroupReady
	case GroupReady:
		return next == GroupDeparted
	case GroupDeparted:
		return next == GroupFinished
	case GroupFinished:
		return false
	}
	return false
}
func (r RiskEvent) CanTransition(next RiskStatus) bool {
	switch r.Status {
	case RiskOpen:
		return next == RiskInvestigating || next == RiskCancelled
	case RiskInvestigating:
		return next == RiskResolved || next == RiskOpen
	case RiskResolved, RiskCancelled:
		return false
	}
	return false
}
func (s Submission) CanTransition(next SubmissionStatus) bool {
	switch s.Status {
	case SubmissionDraft:
		return next == SubmissionSubmitted
	case SubmissionSubmitted:
		return next == SubmissionApproved || next == SubmissionRejected
	case SubmissionRejected:
		return next == SubmissionSubmitted
	case SubmissionApproved:
		return false
	}
	return false
}
