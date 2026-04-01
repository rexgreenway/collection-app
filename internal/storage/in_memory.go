package storage

import (
	"github.com/rexgreenway/collection-app/internal/entities"
)

const InMemory StorageType = "in_memory"

// inMemoryStorage ???
type inMemoryStorage struct {
	store map[string]entities.Collection
}

// NewInMemoryStorage ???
func newInMemoryStorage() (*inMemoryStorage, error) {
	return &inMemoryStorage{store: map[string]entities.Collection{}}, nil
}

// CreateCollection ???
func (s inMemoryStorage) CreateCollection(collection entities.Collection) (entities.Collection, error) {
	if _, ok := s.store[collection.ID]; ok {
		return entities.Collection{}, ErrCollectionAlreadyExists
	}
	s.store[collection.ID] = collection
	return collection, nil
}

// ListCollections ???
func (s inMemoryStorage) ListCollections() (map[string]entities.Collection, error) {
	return s.store, nil
}

// GetCollection ???
func (s inMemoryStorage) GetCollection(collectionID string) (entities.Collection, error) {
	collection, ok := s.store[collectionID]
	if !ok {
		return entities.Collection{}, ErrCollectionNotFound
	}

	return collection, nil
}

// UpdateCollection ???
// Change this to a DIFF method???
func (s *inMemoryStorage) UpdateCollection(id string, collection entities.Collection) (entities.Collection, error) {
	if _, ok := s.store[id]; !ok {
		return entities.Collection{}, ErrCollectionNotFound
	}

	s.store[collection.ID] = collection

	return collection, nil
}

// DeleteCollection ???
func (s *inMemoryStorage) DeleteCollection(collectionID string) error {
	if _, ok := s.store[collectionID]; !ok {
		return ErrCollectionNotFound
	}

	delete(s.store, collectionID)

	return nil
}
