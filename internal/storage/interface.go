package storage

import (
	"github.com/rexgreenway/collection-app/internal/entities"
)

// Storage defines the methodologies for fetching and editing collection data.
type Store interface {
	// ListCollections ???
	ListCollections(pagination *entities.Pagination) ([]entities.Collection, error)

	// CreateCollection adds the provided collection into storage.
	CreateCollection(entities.Collection) (entities.Collection, error)

	// GetCollection looks up and returns a collection from storage given an ID.
	GetCollection(collectionId string) (entities.Collection, error)

	// UpdateCollection ???
	UpdateCollection(collectionId string, collection entities.Collection) (entities.Collection, error)

	// DeleteCollection ???
	DeleteCollection(collectionId string) error

	// ListItemsForCollection ???
	ListItemsByCollectionId(collectionId string, pagination *entities.Pagination) ([]entities.Item, error)

	// CreateItem ???
	CreateItem(entities.Item) (entities.Item, error)

	// CreateItem ???
	CreateItemBatchByCollectionId(collectionId string, items []entities.Item) ([]entities.Item, error)

	// GetItem ???
	GetItem(collectionId string, itemId string) (entities.Item, error)

	// UpdateItem ???
	UpdateItem(collectionId string, itemId string, item entities.Item) (entities.Item, error)

	// DeleteItem
	DeleteItem(collectionId string, itemId string) error

	// ????
	GetItemCountByCollection(collectionId string) int32
}
