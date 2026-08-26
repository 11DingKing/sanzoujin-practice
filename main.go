package main

import (
	"context"
	"github.com/11DingKing/sanzoujin-practice/internal/audit"
	"github.com/11DingKing/sanzoujin-practice/internal/auth"
	"github.com/11DingKing/sanzoujin-practice/internal/config"
	"github.com/11DingKing/sanzoujin-practice/internal/httpapi"
	"github.com/11DingKing/sanzoujin-practice/internal/notify"
	"github.com/11DingKing/sanzoujin-practice/internal/repository"
	"github.com/11DingKing/sanzoujin-practice/internal/service"
	"github.com/11DingKing/sanzoujin-practice/internal/storage"
	"github.com/11DingKing/sanzoujin-practice/internal/worker"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func buildServer(ctx context.Context, c config.Config, logger *slog.Logger) (*httpapi.Server, *worker.Worker, *storage.DB, error) {
	db, e := storage.Open(ctx, c.DBPath)
	if e != nil {
		return nil, nil, nil, e
	}
	aud := audit.Service{Repo: repository.AuditRepo{DB: db.SQL}}
	srv := &httpapi.Server{DB: db, Logger: logger}
	srv.Auth = auth.Service{Users: repository.UserRepo{DB: db.SQL}, Sessions: repository.SessionRepo{DB: db.SQL}, TTL: c.SessionTTL}
	srv.Projects = service.ProjectService{DB: db, Projects: repository.ProjectRepo{DB: db.SQL}, Venues: repository.VenueRepo{DB: db.SQL}, Audit: aud}
	srv.Enrollments = service.EnrollmentService{DB: db, Projects: repository.ProjectRepo{DB: db.SQL}, Enrollments: repository.EnrollmentRepo{DB: db.SQL}, Audit: aud, Idempotency: repository.IdempotencyRepo{DB: db.SQL}}
	srv.Groups = service.GroupService{DB: db, Groups: repository.GroupRepo{DB: db.SQL}, Enrollments: repository.EnrollmentRepo{DB: db.SQL}, Audit: aud}
	srv.Attendance = service.AttendanceService{Repo: repository.AttendanceRepo{DB: db.SQL}, Enrollments: repository.EnrollmentRepo{DB: db.SQL}, Audit: aud}
	srv.Risks = service.RiskService{Repo: repository.RiskRepo{DB: db.SQL}, Projects: repository.ProjectRepo{DB: db.SQL}, Audit: aud}
	srv.Submissions = service.SubmissionService{Repo: repository.SubmissionRepo{DB: db.SQL}, Audit: aud}
	w := &worker.Worker{Outbox: repository.OutboxRepo{DB: db.SQL}, Sender: &notify.MemorySender{}, Interval: c.WorkerInterval, Logger: logger}
	return srv, w, db, nil
}
func run() int {
	cfg := config.Load()
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()
	srv, w, db, e := buildServer(ctx, cfg, logger)
	if e != nil {
		logger.Error("startup", "error", e)
		return 1
	}
	defer db.Close()
	w.Start(ctx)
	defer w.Stop()
	server := &http.Server{Addr: cfg.HTTPAddr, Handler: srv.Handler(), ReadHeaderTimeout: 5 * time.Second, ReadTimeout: 15 * time.Second, WriteTimeout: 15 * time.Second, IdleTimeout: 60 * time.Second}
	go func() {
		<-ctx.Done()
		shutdown, stop := context.WithTimeout(context.Background(), 10*time.Second)
		defer stop()
		_ = server.Shutdown(shutdown)
	}()
	logger.Info("server started", "addr", cfg.HTTPAddr)
	if e := server.ListenAndServe(); e != nil && e != http.ErrServerClosed {
		logger.Error("server", "error", e)
		return 1
	}
	return 0
}
func Main() { os.Exit(run()) }
func main() { Main() }
