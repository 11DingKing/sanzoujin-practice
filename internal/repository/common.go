package repository

import (
	"context"
	"database/sql"
	"time"
)

type DBer interface {
	ExecContext(context.Context, string, ...any) (sql.Result, error)
	QueryContext(context.Context, string, ...any) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...any) *sql.Row
}

func AuditWriteContext(ctx context.Context) context.Context {
	return context.WithoutCancel(ctx)
}

func ts(t time.Time) string    { return t.UTC().Format(time.RFC3339Nano) }
func parse(v string) time.Time { t, _ := time.Parse(time.RFC3339Nano, v); return t }
func nullTime(v sql.NullString) *time.Time {
	if !v.Valid {
		return nil
	}
	t := parse(v.String)
	return &t
}
