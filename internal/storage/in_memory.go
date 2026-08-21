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

	collections     map[string]entities.Collection
	collectionItems map[string]map[string]entities.Item
}

// newInMemoryStorage ???
func newInMemoryStorage(logger *zap.SugaredLogger, utils Utils) (*inMemoryStorage, error) {
	return &inMemoryStorage{
		logger: logger,

		utils: utils,

		collections: map[string]entities.Collection{},

		collectionItems: map[string]map[string]entities.Item{},
	}, nil
}

// ------- Store Utils -------

// stampCreatedAt sets the CreatedAt Metadata for a collection.
func (s inMemoryStorage) stampCreatedAt(collection *entities.Collection) {
	collection.Metadata.CreatedAt = s.utils.NowUTC()
}

// ------- Collection Storage Methods -------

// ListCollections ???
func (s inMemoryStorage) ListCollections(pagination *entities.Pagination) ([]entities.Collection, error) {
	total := int32(len(s.collections))

	paginationBounds := resolvePaginationBounds(pagination, total)

	collections := slices.Collect(maps.Values(s.collections))

	// Sort Collections by CreatedAt in descending order (newest first)
	sort.Slice(collections, func(i, j int) bool {
		return collections[i].Metadata.CreatedAt.After(collections[j].Metadata.CreatedAt)
	})

	return collections[paginationBounds.Start:paginationBounds.End], nil
}

// CreateCollection ???
func (s inMemoryStorage) CreateCollection(collection entities.Collection) (entities.Collection, error) {
	if _, ok := s.collections[collection.Id]; ok {
		return entities.Collection{}, ErrCollectionAlreadyExists
	}

	s.stampCreatedAt(&collection)

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

	return nil
}

// ------- Item Storage Methods -------

// ListItemsByCollectionId ???
func (s *inMemoryStorage) ListItemsByCollectionId(
	collectionId string,
	pagination *entities.Pagination,
) ([]entities.Item, error) {
	if _, ok := s.collections[collectionId]; !ok {
		return []entities.Item{}, ErrCollectionNotFound
	}

	items := s.collectionItems[collectionId]

	total := int32(len(items))

	paginationBounds := resolvePaginationBounds(pagination, total)

	result := slices.Collect(maps.Values(items))[paginationBounds.Start:paginationBounds.End]

	return result, nil
}

// CreateItem ???
func (s *inMemoryStorage) CreateItem(item entities.Item) (entities.Item, error) {
	// Check collection exists
	collectionItems, ok := s.collectionItems[item.CollectionId]
	if !ok {
		return entities.Item{}, ErrCollectionNotFound
	}

	if _, ok := collectionItems[item.Id]; ok {
		return entities.Item{}, ErrItemAlreadyExists
	}
	collectionItems[item.Id] = item

	return item, nil
}

// CreateItemBatchByCollectionId ???
func (s *inMemoryStorage) CreateItemBatchByCollectionId(
	collectionId string,
	items []entities.Item,
) ([]entities.Item, error) {
	// Check collection exists
	if _, ok := s.collections[collectionId]; !ok {
		return []entities.Item{}, ErrCollectionNotFound
	}

	collectionItems := s.collectionItems[collectionId]

	for _, item := range items {
		if _, ok := collectionItems[item.Id]; ok {
			s.logger.Warnf("Item %q already exists, skipping creation.", item.Id)
		} else {
			collectionItems[item.Id] = item
		}
	}

	return items, nil
}

// GetItem ???
func (s *inMemoryStorage) GetItem(
	collectionId string,
	itemId string,
) (entities.Item, error) {
	// Check collection exists
	collectionItems, ok := s.collectionItems[collectionId]
	if !ok {
		return entities.Item{}, ErrCollectionNotFound
	}

	item, ok := collectionItems[itemId]
	if !ok {
		return entities.Item{}, ErrItemNotFound
	}

	return item, nil
}

func (s *inMemoryStorage) UpdateItem(
	collectionId string,
	itemId string,
	item entities.Item,
) (entities.Item, error) {
	// Check collection exists
	collectionItems, ok := s.collectionItems[collectionId]
	if !ok {
		return entities.Item{}, ErrCollectionNotFound
	}

	// Check item exists
	if _, ok := collectionItems[itemId]; !ok {
		return entities.Item{}, ErrItemNotFound
	}

	collectionItems[itemId] = item

	return item, nil
}

// DeleteItem ???
func (s *inMemoryStorage) DeleteItem(
	collectionId string,
	itemId string,
) error {
	collectionItems, ok := s.collectionItems[collectionId]
	if !ok {
		return ErrCollectionNotFound
	}

	// Check item exists
	if _, ok := collectionItems[itemId]; !ok {
		return ErrItemNotFound
	}

	delete(collectionItems, itemId)

	return nil
}

// GetItemCountByCollection ???
func (s *inMemoryStorage) GetItemCountByCollection(collectionId string) int32 {
	if items, ok := s.collectionItems[collectionId]; !ok {
		return 0
	} else {
		return int32(len(items))
	}
}
