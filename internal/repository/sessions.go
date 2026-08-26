package repository

import (
	"context"
	"database/sql"
	"github.com/11DingKing/sanzoujin-practice/internal/domain"
	"time"
)

type SessionRepo struct{ DB DBer }

func (r SessionRepo) Create(ctx context.Context, s domain.Session) error {
	_, e := r.DB.ExecContext(ctx, `INSERT INTO sessions(id,user_id,token_hash,expires_at,created_at) VALUES(?,?,?,?,?)`, s.ID, s.UserID, s.TokenHash, ts(s.ExpiresAt), ts(s.CreatedAt))
	return e
}
func (r SessionRepo) Active(ctx context.Context, hash string, now string) (domain.Session, error) {
	var s domain.Session
	var exp, created string
	var rev sql.NullString
	e := r.DB.QueryRowContext(ctx, `SELECT id,user_id,token_hash,expires_at,revoked_at,created_at FROM sessions WHERE token_hash=? AND revoked_at IS NULL AND expires_at>?`, hash, now).Scan(&s.ID, &s.UserID, &s.TokenHash, &exp, &rev, &created)
	if e == sql.ErrNoRows {
		return s, domain.ErrExpired
	}
	if e != nil {
		return s, e
	}
	s.ExpiresAt = parse(exp)
	s.CreatedAt = parse(created)
	s.RevokedAt = nullTime(rev)
	return s, nil
}
func (r SessionRepo) Revoke(ctx context.Context, id string, now time.Time) error {
	_, e := r.DB.ExecContext(ctx, `UPDATE sessions SET revoked_at=? WHERE id=? AND revoked_at IS NULL`, ts(now), id)
	return e
}
func (r SessionRepo) Purge(ctx context.Context, now time.Time) error {
	_, e := r.DB.ExecContext(ctx, `DELETE FROM sessions WHERE expires_at<? OR revoked_at IS NOT NULL`, ts(now))
	return e
}
