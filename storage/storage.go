package storage

import "errors"

var (
	ErrAlreadyExists = errors.New("already exists")
	ErrNotFound      = errors.New("not found")
)

type URLStorage interface {
	Get(short string) (lng string, err error)
	Post(short, lng string) (err error)
}
