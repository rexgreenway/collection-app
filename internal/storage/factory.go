package storage

import (
	"fmt"
	"time"

	"go.uber.org/zap"
)

// StorageType ???
type StorageType string

// StorageFactory ???
func StorageFactory(impl StorageType, logger *zap.SugaredLogger, opts ...Option) (Store, error) {
	// Initialise the default utils for the implementation.
	utils := Utils{
		NowUTC: func() time.Time { return time.Now().UTC() },
	}

	// Apply any supplied utility mutations
	for _, o := range opts {
		o(&utils)
	}

	switch impl {
	case IN_MEMORY:
		return newInMemoryStorage(logger, utils)

	// case OS:
	// 	return newOSStorage()

	default:
		return nil, fmt.Errorf("Storage implementation %q is not supported", impl)
	}
}
