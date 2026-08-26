package repository

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/11DingKing/sanzoujin-practice/internal/domain"
)

type UserRepo struct{ DB DBer }

func (r UserRepo) Create(ctx context.Context, u domain.User) error {
	_, err := r.DB.ExecContext(ctx, `INSERT INTO users(id,org_id,name,email,password_hash,role,active,created_at) VALUES(?,?,?,?,?,?,?,?)`, u.ID, u.OrgID, u.Name, u.Email, u.PasswordHash, string(u.Role), u.Active, ts(u.CreatedAt))
	return err
}
func (r UserRepo) ByEmail(ctx context.Context, email string) (domain.User, error) {
	var u domain.User
	var active int
	var created string
	err := r.DB.QueryRowContext(ctx, `SELECT id,org_id,name,email,password_hash,role,active,created_at FROM users WHERE email=?`, email).Scan(&u.ID, &u.OrgID, &u.Name, &u.Email, &u.PasswordHash, &u.Role, &active, &created)
	if err == sql.ErrNoRows {
		return u, domain.ErrNotFound
	}
	if err != nil {
		return u, fmt.Errorf("find user: %w", err)
	}
	u.Active = active == 1
	u.CreatedAt = parse(created)
	return u, nil
}
func (r UserRepo) ByID(ctx context.Context, id string) (domain.User, error) {
	var u domain.User
	var active int
	var created string
	err := r.DB.QueryRowContext(ctx, `SELECT id,org_id,name,email,password_hash,role,active,created_at FROM users WHERE id=?`, id).Scan(&u.ID, &u.OrgID, &u.Name, &u.Email, &u.PasswordHash, &u.Role, &active, &created)
	if err == sql.ErrNoRows {
		return u, domain.ErrNotFound
	}
	if err != nil {
		return u, err
	}
	u.Active = active == 1
	u.CreatedAt = parse(created)
	return u, nil
}
