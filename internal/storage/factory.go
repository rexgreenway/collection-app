package storage

import (
	"fmt"

	"go.uber.org/zap"
)

// StorageType ???
type StorageType string

// StorageManager ???
func StorageManager(impl StorageType, logger *zap.SugaredLogger) (Store, error) {
	switch impl {
	case InMemory:
		return newInMemoryStorage(logger)

	// case OS:
	// 	return newOSStorage()

	default:
		return nil, fmt.Errorf("Storage implementation does not exist %q", impl)
	}
}
