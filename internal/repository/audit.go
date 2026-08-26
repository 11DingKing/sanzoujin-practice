package repository

import (
	"context"
	"database/sql"
	"github.com/11DingKing/sanzoujin-practice/internal/domain"
)

type AuditRepo struct{ DB DBer }

func (r AuditRepo) LastHash(ctx context.Context, org string) (string, error) {
	var h string
	e := r.DB.QueryRowContext(ctx, `SELECT hash FROM audit_events WHERE org_id=? ORDER BY created_at DESC LIMIT 1`, org).Scan(&h)
	if e == sql.ErrNoRows {
		return "", nil
	}
	return h, e
}
func (r AuditRepo) Append(ctx context.Context, e domain.AuditEvent) error {
	_, err := r.DB.ExecContext(ctx, `INSERT INTO audit_events(id,org_id,actor_id,object_type,object_id,action,result,request_id,prev_hash,hash,created_at) VALUES(?,?,?,?,?,?,?,?,?,?,?)`, e.ID, e.OrgID, e.ActorID, e.ObjectType, e.ObjectID, e.Action, e.Result, e.RequestID, e.PrevHash, e.Hash, ts(e.CreatedAt))
	return err
}
func (r AuditRepo) List(ctx context.Context, obj, id string) ([]domain.AuditEvent, error) {
	rows, e := r.DB.QueryContext(ctx, `SELECT id,org_id,actor_id,object_type,object_id,action,result,request_id,prev_hash,hash,created_at FROM audit_events WHERE object_type=? AND object_id=? ORDER BY created_at`, obj, id)
	if e != nil {
		return nil, e
	}
	defer rows.Close()
	out := []domain.AuditEvent{}
	for rows.Next() {
		var a domain.AuditEvent
		var at string
		if e := rows.Scan(&a.ID, &a.OrgID, &a.ActorID, &a.ObjectType, &a.ObjectID, &a.Action, &a.Result, &a.RequestID, &a.PrevHash, &a.Hash, &at); e != nil {
			return nil, e
		}
		a.CreatedAt = parse(at)
		out = append(out, a)
	}
	return out, rows.Err()
}

type IdempotencyRepo struct{ DB DBer }

func (r IdempotencyRepo) Get(ctx context.Context, key, method, path string) (string, error) {
	var v string
	e := r.DB.QueryRowContext(ctx, `SELECT response FROM idempotency_keys WHERE key=? AND method=? AND path=?`, key, method, path).Scan(&v)
	if e == sql.ErrNoRows {
		return "", domain.ErrNotFound
	}
	return v, e
}
func (r IdempotencyRepo) Put(ctx context.Context, key, method, path, hash, response string) error {
	_, e := r.DB.ExecContext(ctx, `INSERT INTO idempotency_keys(key,method,path,request_hash,response,created_at) VALUES(?,?,?,?,?,datetime('now'))`, key, method, path, hash, response)
	return e
}
