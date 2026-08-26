package repository

import (
	"context"
	"database/sql"
	"github.com/11DingKing/sanzoujin-practice/internal/domain"
	"time"
)

type OutboxRepo struct{ DB DBer }

func (r OutboxRepo) Enqueue(ctx context.Context, m domain.OutboxMessage) error {
	_, e := r.DB.ExecContext(ctx, `INSERT INTO outbox_messages(id,topic,payload,attempts,next_attempt_at,created_at) VALUES(?,?,?,?,?,?)`, m.ID, m.Topic, m.Payload, m.Attempts, ts(m.NextAttemptAt), ts(m.CreatedAt))
	return e
}
func (r OutboxRepo) Due(ctx context.Context, limit int) ([]domain.OutboxMessage, error) {
	rows, e := r.DB.QueryContext(ctx, `SELECT id,topic,payload,attempts,next_attempt_at,sent_at,last_error,created_at FROM outbox_messages WHERE sent_at IS NULL AND next_attempt_at<=? ORDER BY next_attempt_at LIMIT ?`, ts(time.Now()), limit)
	if e != nil {
		return nil, e
	}
	defer rows.Close()
	out := []domain.OutboxMessage{}
	for rows.Next() {
		var m domain.OutboxMessage
		var next, created string
		var sent sql.NullString
		if e := rows.Scan(&m.ID, &m.Topic, &m.Payload, &m.Attempts, &next, &sent, &m.LastError, &created); e != nil {
			return nil, e
		}
		m.NextAttemptAt = parse(next)
		m.SentAt = nullTime(sent)
		m.CreatedAt = parse(created)
		out = append(out, m)
	}
	return out, rows.Err()
}
func (r OutboxRepo) Mark(ctx context.Context, id string, attempts int, errText string, next time.Time, sent bool) error {
	if sent {
		_, e := r.DB.ExecContext(ctx, `UPDATE outbox_messages SET sent_at=?,attempts=? WHERE id=?`, ts(time.Now()), attempts, id)
		return e
	}
	_, e := r.DB.ExecContext(ctx, `UPDATE outbox_messages SET attempts=?,last_error=?,next_attempt_at=? WHERE id=?`, attempts, errText, ts(next), id)
	return e
}
