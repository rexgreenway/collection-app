package storage

import (
	"errors"

	"github.com/rexgreenway/collection-app/internal/entities"
)

const InMemory StorageType = "in_memory"

// inMemoryStorage ???
type inMemoryStorage struct {
	store map[string]entities.Collection
}

// NewInMemoryStorage ???
func newInMemoryStorage() (inMemoryStorage, error) {
	return inMemoryStorage{store: map[string]entities.Collection{}}, nil
}

// CreateCollection ???
func (s inMemoryStorage) CreateCollection(collection entities.Collection) (entities.Collection, error) {
	if _, ok := s.store[collection.ID]; ok {
		return entities.Collection{}, errors.New("collection already exists")
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
		return entities.Collection{}, errors.New("collection not found")
	}
	return collection, nil
}

// UpdateCollection ???
// Change this to a DIFF method???
func (s *inMemoryStorage) UpdateCollection(collection entities.Collection) error {
	if _, ok := s.store[collection.ID]; !ok {
		return errors.New("collection not found")
	}
	s.store[collection.ID] = collection
	return nil
}

// DeleteCollection ???
func (s *inMemoryStorage) DeleteCollection(collectionID string) error {
	if _, ok := s.store[collectionID]; !ok {
		return errors.New("collection not found")
	}
	delete(s.store, collectionID)
	return nil
}
