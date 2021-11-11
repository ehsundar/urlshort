package mem

import (
	"sync"
	"urlshort/storage"
)

type memStorage struct {
	m    map[string]string
	lock sync.RWMutex
}

func NewMemStorage() storage.URLStorage {
	return &memStorage{
		m: make(map[string]string),
	}
}

func (s *memStorage) GetLong(short string) (lng string, err error) {
	s.lock.RLock()
	defer s.lock.RUnlock()

	lng, ok := s.m[short]
	if !ok {
		err = storage.ErrNotFound
	}
	return
}

func (s *memStorage) Revoke(short string) (err error) {
	_, errGet := s.GetLong(short)
	if errGet != nil {
		return errGet
	}

	s.lock.Lock()
	defer s.lock.Unlock()
	delete(s.m, short)
	return
}

func (s *memStorage) Create(short, lng string) (err error) {
	_, errGet := s.GetLong(short)
	if errGet != storage.ErrNotFound {
		return storage.ErrAlreadyExists
	}

	s.lock.Lock()
	defer s.lock.Unlock()

	s.m[short] = lng
	return
}
