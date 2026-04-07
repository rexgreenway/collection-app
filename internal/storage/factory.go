package storage

import (
	"fmt"

	"go.uber.org/zap"
)

// StorageType ???
type StorageType string

// StorageFactory ???
func StorageFactory(impl StorageType, logger *zap.SugaredLogger) (Store, error) {
	switch impl {
	case IN_MEMORY:
		return newInMemoryStorage(logger)

	// case OS:
	// 	return newOSStorage()

	default:
		return nil, fmt.Errorf("Storage implementation does not exist %q", impl)
	}
}
