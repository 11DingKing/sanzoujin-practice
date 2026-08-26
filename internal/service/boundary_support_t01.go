package service

import "context"

func resumeContextErrorT01(ctx context.Context) error { return ctx.Err() }
