package storage

import (
	"github.com/rexgreenway/collection-app/internal/entities"
)

// Will probably have to add User related stuff to everything in the future...

// Storage defines the methodologies for fetching and editing collection data.
type Store interface {
	// ListCollections ???
	ListCollections(pagination *entities.Pagination) ([]entities.Collection, error)
	// CreateCollection adds the provided collection into storage.
	CreateCollection(entities.Collection) (entities.Collection, error)
	// GetCollection looks up and returns a collection from storage given an ID.
	GetCollection(collectionId string) (entities.Collection, error)
	// UpdateCollection ???
	UpdateCollection(id string, update entities.CollectionUpdate) (entities.Collection, error)
	// DeleteCollection ???
	DeleteCollection(collectionId string) error

	// CreateItem ???
	CreateItem(entities.Item) (entities.Item, error)
	// CreateItemBatch ???
	CreateItemBatch(items []entities.Item) ([]entities.Item, error)
	// GetItem ???
	GetItem(itemId string) (entities.Item, error)
	// UpdateItem ???
	UpdateItem(id string, update entities.ItemUpdate) (entities.Item, error)
	// DeleteItem
	DeleteItem(itemId string) error

	// ListItemsForCollection ???
	ListItemsByCollectionId(collectionId string, pagination *entities.Pagination) ([]entities.Item, error)
	// GetItemCountByCollection ??? RENAME???
	GetItemCountByCollectionId(collectionId string) int32
}
