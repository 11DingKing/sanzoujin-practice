package repository

import (
	"context"
	"github.com/11DingKing/sanzoujin-practice/internal/domain"
)

type VenueRepo struct{ DB DBer }

func (r VenueRepo) Create(ctx context.Context, v domain.Venue) error {
	_, e := r.DB.ExecContext(ctx, `INSERT INTO venues(id,name,address,capacity,open,created_at) VALUES(?,?,?,?,?,?)`, v.ID, v.Name, v.Address, v.Capacity, v.Open, ts(v.CreatedAt))
	return e
}
func (r VenueRepo) ByID(ctx context.Context, id string) (domain.Venue, error) {
	var v domain.Venue
	var open int
	var at string
	e := r.DB.QueryRowContext(ctx, `SELECT id,name,address,capacity,open,created_at FROM venues WHERE id=?`, id).Scan(&v.ID, &v.Name, &v.Address, &v.Capacity, &open, &at)
	if e != nil {
		return v, e
	}
	v.Open = open == 1
	v.CreatedAt = parse(at)
	return v, nil
}
func (r VenueRepo) SetOpen(ctx context.Context, id string, open bool) error {
	_, e := r.DB.ExecContext(ctx, `UPDATE venues SET open=? WHERE id=?`, open, id)
	return e
}
