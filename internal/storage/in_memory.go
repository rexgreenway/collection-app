package storage

import (
	"maps"
	"slices"
	"sort"

	"go.uber.org/zap"

	"github.com/rexgreenway/collection-app/internal/entities"
)

const IN_MEMORY StorageType = "in_memory"

// inMemoryStorage ???
type inMemoryStorage struct {
	logger *zap.SugaredLogger

	utils Utils

	collections map[string]entities.Collection
	items       map[string]entities.Item
}

// newInMemoryStorage ???
func newInMemoryStorage(logger *zap.SugaredLogger, utils Utils) (*inMemoryStorage, error) {
	return &inMemoryStorage{
		logger: logger,

		utils: utils,

		collections: map[string]entities.Collection{},

		items: map[string]entities.Item{},
	}, nil
}

// ------- Collection Storage Methods -------

// ListCollections ???
func (s inMemoryStorage) ListCollections(pagination *entities.Pagination) ([]entities.Collection, error) {
	total := int32(len(s.collections))

	paginationBounds := resolvePaginationBounds(pagination, total)

	collections := slices.Collect(maps.Values(s.collections))

	// Sort Collections by CreatedAt in descending order (newest first, tie breaking using id)
	// Could turn this sorting into a util that can be uniformly used across stores / resources with equivalent metadata
	sort.Slice(collections, func(i, j int) bool {
		a, b := collections[i], collections[j]
		if a.Metadata.CreatedAt.Equal(b.Metadata.CreatedAt) {
			return a.Id < b.Id
		}
		return a.Metadata.CreatedAt.After(b.Metadata.CreatedAt)
	})

	return collections[paginationBounds.Start:paginationBounds.End], nil
}

// CreateCollection ???
func (s inMemoryStorage) CreateCollection(collection entities.Collection) (entities.Collection, error) {
	if _, ok := s.collections[collection.Id]; ok {
		return entities.Collection{}, ErrCollectionAlreadyExists
	}

	collection.StampCreatedAt(s.utils.NowUTC())

	s.collections[collection.Id] = collection

	return collection, nil
}

// GetCollection ???
func (s inMemoryStorage) GetCollection(collectionId string) (entities.Collection, error) {
	collection, ok := s.collections[collectionId]
	if !ok {
		return entities.Collection{}, ErrCollectionNotFound
	}

	return collection, nil
}

// UpdateCollection ??? (Change this to a diff method??)
func (s *inMemoryStorage) UpdateCollection(id string, update entities.CollectionUpdate) (entities.Collection, error) {
	collection, ok := s.collections[id]
	if !ok {
		return entities.Collection{}, ErrCollectionNotFound
	}

	collection.Update(update)

	s.collections[id] = collection

	return collection, nil
}

// DeleteCollection ???
func (s *inMemoryStorage) DeleteCollection(collectionId string) error {
	if _, ok := s.collections[collectionId]; !ok {
		return ErrCollectionNotFound
	}

	delete(s.collections, collectionId)

	// Remove the CollectionId from items.
	for _, item := range s.items {
		if item.CollectionId == collectionId {
			item.CollectionId = entities.MISSING_COLLECTION_ID
			s.items[item.Id] = item
		}
	}

	return nil
}

// ------- Item Storage Methods -------

// CreateItem ???
func (s *inMemoryStorage) CreateItem(item entities.Item) (entities.Item, error) {
	if _, ok := s.items[item.Id]; ok {
		return entities.Item{}, ErrItemAlreadyExists
	}

	item.StampCreatedAt(s.utils.NowUTC())

	s.items[item.Id] = item

	return item, nil
}

// CreateItemBatchByCollectionId ??? -- Does this method need a upper limit ???
func (s *inMemoryStorage) CreateItemBatch(items []entities.Item) ([]entities.Item, error) {
	for _, item := range items {
		if _, ok := s.items[item.Id]; ok {
			s.logger.Warnf("Item %q already exists, skipping creation.", item.Id)
		} else {
			item.StampCreatedAt(s.utils.NowUTC())
			s.items[item.Id] = item
		}
	}

	return items, nil
}

// GetItem ???
func (s *inMemoryStorage) GetItem(itemId string) (entities.Item, error) {
	item, ok := s.items[itemId]
	if !ok {
		return entities.Item{}, ErrItemNotFound
	}

	return item, nil
}

func (s *inMemoryStorage) UpdateItem(itemId string, update entities.ItemUpdate) (entities.Item, error) {
	item, ok := s.items[itemId]
	if !ok {
		return entities.Item{}, ErrItemNotFound
	}

	item.Update(update)

	s.items[itemId] = item

	return item, nil
}

// DeleteItem ???
func (s *inMemoryStorage) DeleteItem(itemId string) error {
	if _, ok := s.items[itemId]; !ok {
		return ErrItemNotFound
	}

	delete(s.items, itemId)

	return nil
}

// ListItemsByCollectionId ???
func (s *inMemoryStorage) ListItemsByCollectionId(
	collectionId string,
	pagination *entities.Pagination,
) ([]entities.Item, error) {
	var collectionItems []entities.Item
	for _, item := range s.items {
		if item.CollectionId == collectionId {
			collectionItems = append(collectionItems, item)
		}
	}

	total := int32(len(collectionItems))

	paginationBounds := resolvePaginationBounds(pagination, total)

	// Sort Items by CreatedAt in descending order (newest first, tie breaking using id)
	sort.Slice(collectionItems, func(i, j int) bool {
		a, b := collectionItems[i], collectionItems[j]
		if a.Metadata.CreatedAt.Equal(b.Metadata.CreatedAt) {
			return a.Id < b.Id
		}
		return a.Metadata.CreatedAt.After(b.Metadata.CreatedAt)
	})

	return collectionItems[paginationBounds.Start:paginationBounds.End], nil
}

// GetItemCountByCollection ???
func (s *inMemoryStorage) GetItemCountByCollectionId(collectionId string) int32 {
	var collectionItems []entities.Item
	for _, item := range s.items {
		if item.CollectionId == collectionId {
			collectionItems = append(collectionItems, item)
		}
	}

	return int32(len(collectionItems))
}
