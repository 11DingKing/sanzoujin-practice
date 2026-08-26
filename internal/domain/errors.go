package domain

import "errors"

var (
	ErrNotFound    = errors.New("not found")
	ErrConflict    = errors.New("conflict")
	ErrForbidden   = errors.New("forbidden")
	ErrInvalid     = errors.New("invalid input")
	ErrCapacity    = errors.New("capacity exceeded")
	ErrIdempotency = errors.New("idempotency conflict")
	ErrExpired     = errors.New("session expired")
)

type ValidationError struct {
	Field   string
	Message string
}

func (e ValidationError) Error() string { return e.Field + ": " + e.Message }
func IsTerminal(status string) bool     { return status == "sent" || status == "permanent_failed" }
