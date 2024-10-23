package storage

import "sync"

// Storage is a simple map with thread safe access.
type Storage struct {
	values map[string]string

	lock sync.Mutex
}

// New returns a new instance of Storage.
func New() *Storage {
	return &Storage{
		values: map[string]string{},
		lock:   sync.Mutex{},
	}
}

// Get returns the value at location "key", and a boolean indicating whether it exists or not.
func (s *Storage) Get(key string) (string, bool, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	value, exists := s.values[key]
	return value, exists, nil
}

// Set replaces the value at "key" with the provided value.
func (s *Storage) Set(key, value string) error {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.values[key] = value
	return nil
}
