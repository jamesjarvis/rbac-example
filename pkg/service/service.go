package service

import "fmt"

// StorageClient provides access to a key-value map.
type StorageClient interface {
	Get(key string) (value string, exists bool, err error)
	Set(key string, value string) error
}

// AccessClient enables the ability to check authorisation for a given user.
type AccessClient interface {
	Check(userID string, action string) bool
}

// Service is an implementation of the key-value service
type Service struct {
	storage StorageClient
	access  AccessClient
}

// New returns a new instance of the key-value service.
func New(
	storage StorageClient,
	access AccessClient,
) *Service {
	return &Service{
		storage: storage,
		access:  access,
	}
}

// Get accepts a user and a key, and returns either the value associated with that key,
// or an error if encountered.
func (s *Service) Get(userID string, key string) (string, error) {
	if !s.access.Check(userID, "get") {
		return "", Error_UNAUTHORISED
	}

	value, ok, err := s.storage.Get(key)
	if err != nil {
		return "", fmt.Errorf("error getting value for key:%q error:%w", key, err)
	}
	if !ok {
		return "", fmt.Errorf("error getting value for key:%q error:%w", key, Error_NOTFOUND)
	}
	return value, nil
}

// Set accepts a user, key and value, and will attempt to set the value with that key,
// or return an error if encountered.
func (s *Service) Set(userID, key, value string) error {
	if !s.access.Check(userID, "set") {
		return Error_UNAUTHORISED
	}

	err := s.storage.Set(key, value)
	if err != nil {
		return fmt.Errorf("error setting value for key:%q error:%w", key, err)
	}
	return nil
}
