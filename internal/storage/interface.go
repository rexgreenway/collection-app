package storage

import (
	"github.com/rexgreenway/collection-app/internal/entities"
)

// Storage defines the methodologies for fetching and editing collection data.
type Store interface {
	// ListCollections ???
	ListCollections() (map[string]entities.Collection, error)

	// CreateCollection adds the provided collection into storage.
	CreateCollection(entities.Collection) (entities.Collection, error)

	// GetCollection looks up and returns a collection from storage given an ID.
	GetCollection(collectionID string) (entities.Collection, error)

	// UpdateCollection ???
	UpdateCollection(collectionID string, collection entities.Collection) (entities.Collection, error)

	// DeleteCollection ???
	DeleteCollection(collectionID string) error
}
