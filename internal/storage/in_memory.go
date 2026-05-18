package storage

import (
	"maps"
	"slices"

	"go.uber.org/zap"

	"github.com/rexgreenway/collection-app/internal/entities"
)

const InMemory StorageType = "in_memory"

// inMemoryStorage ???
type inMemoryStorage struct {
	collections map[string]entities.Collection
	items       map[string][]entities.Item

	logger *zap.SugaredLogger
}

// NewInMemoryStorage ???
func newInMemoryStorage(logger *zap.SugaredLogger) (*inMemoryStorage, error) {
	return &inMemoryStorage{collections: map[string]entities.Collection{}, logger: logger}, nil
}

// CreateCollection ???
func (s inMemoryStorage) CreateCollection(collection entities.Collection) (entities.Collection, error) {
	if _, ok := s.collections[collection.ID]; ok {
		return entities.Collection{}, ErrCollectionAlreadyExists
	}
	s.collections[collection.ID] = collection
	return collection, nil
}

// ListCollections ???
func (s inMemoryStorage) ListCollections(pagination *entities.Pagination) ([]entities.Collection, error) {
	total := int32(len(s.collections))

	validateTransformPagination(pagination, total)

	result := slices.Collect(maps.Values(s.collections))[pagination.Start:pagination.End]

	return result, nil
}

// GetCollection ???
func (s inMemoryStorage) GetCollection(collectionID string) (entities.Collection, error) {
	collection, ok := s.collections[collectionID]
	if !ok {
		return entities.Collection{}, ErrCollectionNotFound
	}

	return collection, nil
}

// UpdateCollection ???
// Change this to a DIFF method???
func (s *inMemoryStorage) UpdateCollection(id string, collection entities.Collection) (entities.Collection, error) {
	if _, ok := s.collections[id]; !ok {
		return entities.Collection{}, ErrCollectionNotFound
	}

	s.collections[collection.ID] = collection

	return collection, nil
}

// DeleteCollection ???
func (s *inMemoryStorage) DeleteCollection(collectionID string) error {
	if _, ok := s.collections[collectionID]; !ok {
		return ErrCollectionNotFound
	}

	delete(s.collections, collectionID)

	return nil
}

// GetItemCountByCollection ???
func (s *inMemoryStorage) GetItemCountByCollection(collectionID string) int32 {
	if items, ok := s.items[collectionID]; !ok {
		return 0
	} else {
		return int32(len(items))
	}
}
