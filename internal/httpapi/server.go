package httpapi

import (
	"encoding/json"
	"errors"
	"github.com/11DingKing/sanzoujin-practice/internal/auth"
	"github.com/11DingKing/sanzoujin-practice/internal/domain"
	"github.com/11DingKing/sanzoujin-practice/internal/middleware"
	"github.com/11DingKing/sanzoujin-practice/internal/service"
	"github.com/11DingKing/sanzoujin-practice/internal/storage"
	"log/slog"
	"net/http"
	"strings"
)

type Server struct {
	DB          *storage.DB
	Auth        auth.Service
	Projects    service.ProjectService
	Enrollments service.EnrollmentService
	Groups      service.GroupService
	Attendance  service.AttendanceService
	Risks       service.RiskService
	Submissions service.SubmissionService
	Logger      *slog.Logger
}

func (s *Server) Handler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, r *http.Request) { jsonWrite(w, 200, map[string]string{"status": "ok"}) })
	mux.HandleFunc("GET /readyz", s.ready)
	mux.HandleFunc("POST /api/v1/auth/login", s.login)
	mux.HandleFunc("POST /api/v1/auth/logout", s.logout)
	mux.HandleFunc("GET /api/v1/projects", s.listProjects)
	mux.HandleFunc("POST /api/v1/projects", s.createProject)
	mux.HandleFunc("POST /api/v1/projects/", s.projectAction)
	mux.HandleFunc("POST /api/v1/enrollments", s.enroll)
	mux.HandleFunc("POST /api/v1/enrollments/", s.enrollmentAction)
	mux.HandleFunc("POST /api/v1/groups", s.createGroup)
	mux.HandleFunc("POST /api/v1/groups/", s.groupAction)
	mux.HandleFunc("POST /api/v1/attendance/", s.attendance)
	mux.HandleFunc("POST /api/v1/risks", s.reportRisk)
	mux.HandleFunc("POST /api/v1/submissions", s.createSubmission)
	mux.HandleFunc("POST /api/v1/submissions/", s.submissionAction)
	return middleware.Chain(mux, s.Logger)
}
func (s *Server) ready(w http.ResponseWriter, r *http.Request) {
	if err := s.DB.Ping(r.Context()); err != nil {
		jsonWrite(w, 503, map[string]string{"status": "not_ready"})
		return
	}
	jsonWrite(w, 200, map[string]string{"status": "ready"})
}
func (s *Server) actor(r *http.Request) (domain.User, domain.Session, error) {
	h := r.Header.Get("Authorization")
	if !strings.HasPrefix(h, "Bearer ") {
		return domain.User{}, domain.Session{}, domain.ErrForbidden
	}
	return s.Auth.Authenticate(r.Context(), strings.TrimPrefix(h, "Bearer "))
}
func (s *Server) login(w http.ResponseWriter, r *http.Request) {
	var in struct{ Email, Password string }
	if !decode(r, &in) {
		jsonWrite(w, 400, map[string]string{"error": "invalid json"})
		return
	}
	u, sess, tok, e := s.Auth.Login(r.Context(), in.Email, in.Password)
	if e != nil {
		jsonWrite(w, 401, map[string]string{"error": "invalid credentials"})
		return
	}
	jsonWrite(w, 200, map[string]any{"token": tok, "expires_at": sess.ExpiresAt, "user": u})
}
func (s *Server) logout(w http.ResponseWriter, r *http.Request) {
	_, sess, e := s.actor(r)
	if e != nil {
		jsonWrite(w, 401, map[string]string{"error": "unauthorized"})
		return
	}
	if e = s.Auth.Logout(r.Context(), sess); e != nil {
		jsonWrite(w, 500, map[string]string{"error": "logout failed"})
		return
	}
	jsonWrite(w, 204, nil)
}
func (s *Server) listProjects(w http.ResponseWriter, r *http.Request) {
	if _, _, e := s.actor(r); e != nil {
		jsonWrite(w, 401, map[string]string{"error": "unauthorized"})
		return
	}
	items, e := s.Projects.List(r.Context(), r.URL.Query().Get("status"), 20)
	if e != nil {
		jsonWrite(w, 500, map[string]string{"error": e.Error()})
		return
	}
	jsonWrite(w, 200, items)
}
func (s *Server) createProject(w http.ResponseWriter, r *http.Request) {
	u, _, e := s.actor(r)
	if e != nil || !domain.CanManageProject(u.Role) {
		jsonWrite(w, 403, map[string]string{"error": "forbidden"})
		return
	}
	var p domain.Project
	if !decode(r, &p) {
		jsonWrite(w, 400, map[string]string{"error": "invalid json"})
		return
	}
	p.CreatedBy = u.ID
	p.OrgID = u.OrgID
	e = s.Projects.Create(r.Context(), p)
	if e != nil {
		mapError(w, e)
		return
	}
	jsonWrite(w, 201, p)
}
func (s *Server) projectAction(w http.ResponseWriter, r *http.Request) {
	u, _, e := s.actor(r)
	if e != nil {
		jsonWrite(w, 401, map[string]string{"error": "unauthorized"})
		return
	}
	id := strings.TrimPrefix(r.URL.Path, "/api/v1/projects/")
	var in struct {
		Status  domain.ProjectStatus
		Version int
	}
	if !decode(r, &in) {
		jsonWrite(w, 400, map[string]string{"error": "invalid json"})
		return
	}
	e = s.Projects.ChangeStatus(r.Context(), id, in.Status, in.Version, u.ID, u.OrgID)
	if e != nil {
		mapError(w, e)
		return
	}
	jsonWrite(w, 200, map[string]string{"status": "updated"})
}
func (s *Server) enroll(w http.ResponseWriter, r *http.Request) {
	u, _, e := s.actor(r)
	if e != nil {
		jsonWrite(w, 401, map[string]string{"error": "unauthorized"})
		return
	}
	var in domain.Enrollment
	if !decode(r, &in) {
		jsonWrite(w, 400, map[string]string{"error": "invalid json"})
		return
	}
	in.StudentID = u.ID
	created, e2 := s.Enrollments.Enroll(r.Context(), in, r.Method, r.URL.Path, r.Header.Get("Idempotency-Key"))
	if e2 != nil {
		mapError(w, e2)
		return
	}
	jsonWrite(w, 201, created)
}
func (s *Server) enrollmentAction(w http.ResponseWriter, r *http.Request) {
	u, _, e := s.actor(r)
	if e != nil {
		jsonWrite(w, 401, map[string]string{"error": "unauthorized"})
		return
	}
	id := strings.TrimPrefix(r.URL.Path, "/api/v1/enrollments/")
	if strings.HasSuffix(id, "/authorize") {
		id = strings.TrimSuffix(id, "/authorize")
		_, e = s.Enrollments.Authorize(r.Context(), id, u.ID, u.OrgID)
	}
	if e != nil {
		mapError(w, e)
		return
	}
	jsonWrite(w, 200, map[string]string{"status": "ok"})
}
func (s *Server) createGroup(w http.ResponseWriter, r *http.Request) {
	u, _, e := s.actor(r)
	if e != nil || !domain.CanManageProject(u.Role) {
		jsonWrite(w, 403, map[string]string{"error": "forbidden"})
		return
	}
	var g domain.Group
	if !decode(r, &g) {
		jsonWrite(w, 400, map[string]string{"error": "invalid json"})
		return
	}
	g.MentorID = u.ID
	e = s.Groups.Create(r.Context(), g, u.OrgID, u.ID)
	if e != nil {
		mapError(w, e)
		return
	}
	jsonWrite(w, 201, g)
}
func (s *Server) groupAction(w http.ResponseWriter, r *http.Request) {
	u, _, e := s.actor(r)
	if e != nil {
		jsonWrite(w, 401, map[string]string{"error": "unauthorized"})
		return
	}
	id := strings.TrimPrefix(r.URL.Path, "/api/v1/groups/")
	var in struct{ EnrollmentID string }
	if !decode(r, &in) {
		jsonWrite(w, 400, map[string]string{"error": "invalid json"})
		return
	}
	e = s.Groups.Add(r.Context(), id, in.EnrollmentID, u.OrgID, u.ID)
	if e != nil {
		mapError(w, e)
		return
	}
	jsonWrite(w, 200, map[string]string{"status": "matched"})
}
func (s *Server) attendance(w http.ResponseWriter, r *http.Request) {
	u, _, e := s.actor(r)
	if e != nil {
		jsonWrite(w, 401, map[string]string{"error": "unauthorized"})
		return
	}
	action := strings.TrimPrefix(r.URL.Path, "/api/v1/attendance/")
	var in struct {
		EnrollmentID string
		Note         string
	}
	if !decode(r, &in) {
		jsonWrite(w, 400, map[string]string{"error": "invalid json"})
		return
	}
	var a domain.Attendance
	if action == "check-in" {
		a, e = s.Attendance.CheckIn(r.Context(), in.EnrollmentID, u.ID, u.OrgID, in.Note)
	} else {
		a, e = s.Attendance.CheckOut(r.Context(), in.EnrollmentID, u.ID, u.OrgID)
	}
	if e != nil {
		mapError(w, e)
		return
	}
	jsonWrite(w, 200, a)
}
func (s *Server) reportRisk(w http.ResponseWriter, r *http.Request) {
	u, _, e := s.actor(r)
	if e != nil {
		jsonWrite(w, 401, map[string]string{"error": "unauthorized"})
		return
	}
	var x domain.RiskEvent
	if !decode(r, &x) {
		jsonWrite(w, 400, map[string]string{"error": "invalid json"})
		return
	}
	e = s.Risks.Report(r.Context(), x, u.OrgID, u.ID)
	if e != nil {
		mapError(w, e)
		return
	}
	jsonWrite(w, 201, x)
}
func (s *Server) createSubmission(w http.ResponseWriter, r *http.Request) {
	u, _, e := s.actor(r)
	if e != nil {
		jsonWrite(w, 401, map[string]string{"error": "unauthorized"})
		return
	}
	var x domain.Submission
	if !decode(r, &x) {
		jsonWrite(w, 400, map[string]string{"error": "invalid json"})
		return
	}
	x.StudentID = u.ID
	e = s.Submissions.Create(r.Context(), x, u.OrgID, u.ID)
	if e != nil {
		mapError(w, e)
		return
	}
	jsonWrite(w, 201, x)
}
func (s *Server) submissionAction(w http.ResponseWriter, r *http.Request) {
	u, _, e := s.actor(r)
	if e != nil {
		jsonWrite(w, 401, map[string]string{"error": "unauthorized"})
		return
	}
	parts := strings.Split(strings.TrimPrefix(r.URL.Path, "/api/v1/submissions/"), "/")
	var in struct {
		Version int
		Status  domain.SubmissionStatus
	}
	if !decode(r, &in) {
		jsonWrite(w, 400, map[string]string{"error": "invalid json"})
		return
	}
	if len(parts) > 1 && parts[1] == "review" {
		e = s.Submissions.Review(r.Context(), parts[0], in.Status, in.Version, u.ID, u.OrgID)
	} else {
		e = s.Submissions.Submit(r.Context(), parts[0], in.Version, u.OrgID, u.ID)
	}
	if e != nil {
		mapError(w, e)
		return
	}
	jsonWrite(w, 200, map[string]string{"status": "ok"})
}
func decode(r *http.Request, v any) bool { return json.NewDecoder(r.Body).Decode(v) == nil }
func jsonWrite(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if v != nil {
		_ = json.NewEncoder(w).Encode(v)
	}
}
func mapError(w http.ResponseWriter, e error) {
	status := 500
	if errors.Is(e, domain.ErrNotFound) {
		status = 404
	} else if errors.Is(e, domain.ErrForbidden) {
		status = 403
	} else if errors.Is(e, domain.ErrInvalid) {
		status = 400
	} else if errors.Is(e, domain.ErrConflict) || errors.Is(e, domain.ErrCapacity) {
		status = 409
	}
	jsonWrite(w, status, map[string]string{"error": e.Error()})
}
