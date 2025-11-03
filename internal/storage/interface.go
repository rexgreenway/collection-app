package storage

import (
	"errors"

	"github.com/rexgreenway/collection-app/internal/entities"
)

var ErrAlreadyExists = errors.New("collection already exists")

// Storage defines the methodologies for fetching and editing collection data.
type Storage interface {
	// COLLECTION LEVEL
	// ListCollections ???
	ListCollections() (map[string]entities.Collection, error)

	// CreateCollection adds the provided collection into storage.
	CreateCollection(entities.Collection) (entities.Collection, error)

	// GetCollection looks up and returns a collection from storage given an ID.
	GetCollection(collectionID string) (entities.Collection, error)

	// // DeleteCollection ???
	// DeleteCollection(collectionID string) error
}
