package storage

import (
	"context"
	"errors"
)

var (
	ErrAlreadyExists = errors.New("already exists")
	ErrNotFound      = errors.New("not found")
)

type URLStorage interface {
	GetLong(ctx context.Context, short string) (long string, err error)
	Create(ctx context.Context, short, long string) (err error)
	Revoke(ctx context.Context, short string) (err error)
}
