package collection

import (
	"testing"

	"github.com/rexgreenway/collection-app/internal/storage"
	"github.com/rexgreenway/collection-app/internal/storage/mocks"
	"go.uber.org/zap"
)

// newTestService builds a service backed by an in-memory store.
func newTestService(t *testing.T, opts ...Option) *CollectionService {
	t.Helper()

	logger := zap.NewNop().Sugar()

	store, err := storage.StorageFactory(storage.IN_MEMORY, logger)
	if err != nil {
		t.Fatalf("failed to create store: %v", err)
	}

	return NewService(logger, store, opts...)
}

// newMockedService builds a service backed by a mocked store so tests can force
// storage errors and verify how the service translates them.
func newMockedService(t *testing.T, opts ...Option) (*CollectionService, *mocks.MockStore) {
	t.Helper()

	store := mocks.NewMockStore(t)

	svc := NewService(zap.NewNop().Sugar(), store, opts...)

	return svc, store
}
