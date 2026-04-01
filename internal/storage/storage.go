package storage

import (
	"errors"
	"fmt"
)

// EXCEPTIONS
var (
	ErrCollectionNotFound      = errors.New("collection not found")
	ErrCollectionAlreadyExists = errors.New("collection already exists")
)

// StorageType defines ???
type StorageType string

// StorageManager returns the ???
func StorageManager(impl StorageType) (Store, error) {
	switch impl {
	case InMemory:
		return newInMemoryStorage()

	// case OS:
	// 	return newOSStorage()

	default:
		return nil, fmt.Errorf("Storage implementation does not exist %q", impl)
	}
}
