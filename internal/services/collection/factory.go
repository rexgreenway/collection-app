package collection

import (
	"fmt"

	"go.uber.org/zap"

	"github.com/rexgreenway/collection-app/internal/storage"
)

// ServiceType ???
type ServiceType string

// CollectionServiceFactory ???
func CollectionServiceFactory(
	impl ServiceType,
	logger *zap.SugaredLogger,
	store storage.Store,
) (CollectionService, error) {
	switch impl {
	case GRPC:
		return newGrpcServer(logger, store), nil

	default:
		return nil, fmt.Errorf("Collection Service implementation %q does not exist.", impl)
	}
}
