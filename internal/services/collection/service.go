package collection

import (
	"go.uber.org/zap"

	"github.com/google/uuid"
	pb "github.com/rexgreenway/collection-app/internal/gen/v1/collection"
	"github.com/rexgreenway/collection-app/internal/storage"
)

// CollectionService satisfies the gRPC CollectionService server interface.
type CollectionService struct {
	// This adds forward compatibility to this implementation of the server
	pb.UnimplementedCollectionServiceServer

	logger *zap.SugaredLogger
	store  storage.Store
	utils  Utils
}

func NewService(logger *zap.SugaredLogger, store storage.Store, opts ...Option) *CollectionService {
	// Initialise the default utils for the service.
	utils := Utils{
		NewId: uuid.NewString,
	}

	// Apply any supplied utility mutations
	for _, o := range opts {
		o(&utils)
	}

	// Create the default service
	return &CollectionService{
		logger: logger,
		store:  store,
		utils:  utils,
	}
}
