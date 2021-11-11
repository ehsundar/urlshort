package pg

import (
	"context"
	"database/sql"
	"urlshort/storage"
)

type pgStorage struct {
	q *Queries
}

func NewStorage(db *sql.DB) storage.Storage {
	return &pgStorage{
		q: New(db),
	}
}

func (s *pgStorage) GetLong(ctx context.Context, short string) (long string, err error) {
	return s.q.GetLong(ctx, short)
}

func (s *pgStorage) Create(ctx context.Context, short, long string) (err error) {
	_, err = s.q.Create(ctx, CreateParams{
		Short: short,
		Long:  long,
	})
	return err
}

func (s *pgStorage) Revoke(ctx context.Context, short string) (err error) {
	return s.q.Revoke(ctx, short)
}
