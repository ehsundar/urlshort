package storage

import "errors"

var (
	ErrAlreadyExists = errors.New("already exists")
	ErrNotFound      = errors.New("not found")
)

type URLStorage interface {
	GetLong(short string) (long string, err error)
	Create(short, long string) (err error)
	Revoke(short string) (err error)
}
