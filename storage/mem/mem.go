package mem

import (
	"context"
	"sync"
	"time"
	"urlshort/storage"
)

type memStorage struct {
	m    map[string]string
	lock sync.RWMutex
}

func NewMemStorage() storage.Storage {
	return &memStorage{
		m: make(map[string]string),
	}
}

func (s *memStorage) GetLong(_ context.Context, short string) (long string, err error) {
	s.lock.RLock()
	defer s.lock.RUnlock()

	long, ok := s.m[short]
	if !ok {
		err = storage.ErrNotFound
	}
	return
}

func (s *memStorage) Delete(ctx context.Context, short string) (err error) {
	_, errGet := s.GetLong(ctx, short)
	if errGet != nil {
		return errGet
	}

	s.lock.Lock()
	defer s.lock.Unlock()
	delete(s.m, short)
	return
}

func (s *memStorage) Create(ctx context.Context, short, long string) (err error) {
	_, errGet := s.GetLong(ctx, short)
	if errGet != storage.ErrNotFound {
		return storage.ErrAlreadyExists
	}

	s.lock.Lock()
	defer s.lock.Unlock()

	s.m[short] = long
	return
}

func (s *memStorage) List(ctx context.Context) (items []storage.Item, err error) {
	s.lock.RLock()
	defer s.lock.RUnlock()

	for short, long := range s.m {
		items = append(items, storage.Item{
			Short:     short,
			Long:      long,
			CreatedAt: time.Now(),
		})
	}

	return
}
