package storage

import (
	"context"
	"errors"
	"time"
)

var (
	ErrAlreadyExists = errors.New("already exists")
	ErrNotFound      = errors.New("not found")
)

type Item struct {
	Short     string
	Long      string
	CreatedAt time.Time
}

type Storage interface {
	GetLong(ctx context.Context, short string) (long string, err error)
	Create(ctx context.Context, short, long string) (err error)
	Delete(ctx context.Context, short string) (err error)
	List(ctx context.Context) (items []Item, err error)
}
