package storage

import (
	"fmt"
	"testing"
	"time"

	"github.com/rexgreenway/collection-app/internal/entities"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func runCollectionCRUDTests(
	t *testing.T,
	newStore func(t *testing.T, opts ...Option) Store,
) {
	testNow := time.Now().UTC()
	testNowUTCFunc := func() time.Time { return testNow }

	// newTestStore returns a store with predicable utility functions for testing.
	newTestStore := func() Store {
		return newStore(
			t,
			WithNowUTCFunc(testNowUTCFunc),
		)
	}

	// sharedTestStore is single shared store for use in tests that do not
	// require an empty store to assert correct returns.
	sharedTestStore := newStore(t)

	// LIST

	t.Run("list when nothing exists returns empty array", func(t *testing.T) {
		store := newStore(t)

		collections, err := store.ListCollections(nil)

		require.NoError(t, err)
		assert.Empty(t, collections)
	})

	t.Run("list with default pagination succeeds", func(t *testing.T) {
		store := newStore(t)

		// Create multiple collections
		for i := 0; i < 3; i++ {
			_, err := store.CreateCollection(
				entities.Collection{
					Id:   fmt.Sprintf("col-%v", i),
					Name: fmt.Sprintf("Collection %v", i),
				},
			)
			require.NoError(t, err)
		}

		collections, err := store.ListCollections(nil)

		require.NoError(t, err)
		assert.Equal(t, 3, len(collections))
	})

	t.Run("list pagination returns distinct results", func(t *testing.T) {
		store := newStore(t)

		// Create multiple collections
		for i := 0; i < 8; i++ {
			_, err := store.CreateCollection(
				entities.Collection{
					Id:   fmt.Sprintf("col-%v", i),
					Name: fmt.Sprintf("Collection %v", i),
				},
			)
			require.NoError(t, err)
		}

		page1, err := store.ListCollections(
			&entities.Pagination{
				Page:     1,
				PageSize: 5,
			},
		)
		require.NoError(t, err)
		assert.Equal(t, 5, len(page1))

		page2, err := store.ListCollections(
			&entities.Pagination{
				Page:     2,
				PageSize: 5,
			},
		)
		require.NoError(t, err)
		assert.Equal(t, 3, len(page2))

		// Check that the returned lists are distinct
		for _, item := range page1 {
			assert.NotContains(
				t,
				page2,
				item,
				fmt.Sprintf("%s should not appear in page 2: %v", item.Id, page2),
			)
		}
	})

	// CREATE

	t.Run("create & populate metadata succeeds", func(t *testing.T) {
		store := newTestStore()

		col := entities.Collection{Id: "create-test-id", Name: "Test"}
		created, err := store.CreateCollection(col)

		require.NoError(t, err)
		assert.Equal(
			t,
			entities.Collection{
				Id:   col.Id,
				Name: col.Name,
				Metadata: entities.Metadata{
					CreatedAt: testNow,
				},
			},
			created,
		)
	})

	t.Run("create duplicate fails", func(t *testing.T) {
		col := entities.Collection{Id: "create-dupe-id", Name: "Dupe"}

		_, _ = sharedTestStore.CreateCollection(col)
		_, err := sharedTestStore.CreateCollection(col)

		assert.ErrorIs(t, err, ErrCollectionAlreadyExists)
	})

	// FETCH

	t.Run("fetching a new collection succeeds", func(t *testing.T) {
		store := newTestStore()

		colId := "collection-test-id"
		col := entities.Collection{Id: colId, Name: "Test"}

		_, err := store.CreateCollection(col)
		require.NoError(t, err)

		fetched, err := store.GetCollection(colId)

		require.NoError(t, err)
		assert.Equal(
			t,
			entities.Collection{
				Id:   col.Id,
				Name: col.Name,
				Metadata: entities.Metadata{
					CreatedAt: testNow,
				},
			},
			fetched,
		)
	})

	t.Run("fetching a non-existent collection fails", func(t *testing.T) {
		_, err := sharedTestStore.GetCollection("no-such-collection")

		assert.ErrorIs(t, err, ErrCollectionNotFound)
	})

	// UPDATE

	t.Run("updating collection name succeeds", func(t *testing.T) {
		store := newTestStore()

		colId := "update-test-id"
		col := entities.Collection{Id: colId, Name: "First Name"}

		_, err := store.CreateCollection(col)
		require.NoError(t, err)

		updatedName := "New Name"

		_, err = store.UpdateCollection(
			colId,
			entities.CollectionUpdate{Name: &updatedName},
		)
		require.NoError(t, err)

		updated, err := store.GetCollection(colId)

		require.NoError(t, err)
		assert.Equal(
			t,
			entities.Collection{
				Id:   col.Id,
				Name: updatedName,
				Metadata: entities.Metadata{
					CreatedAt: testNow,
				},
			},
			updated,
		)
	})

	t.Run("updating a non-existent collection fails", func(t *testing.T) {
		newName := "New Name"
		_, err := sharedTestStore.UpdateCollection(
			"no-such-collection",
			entities.CollectionUpdate{Name: &newName},
		)

		assert.ErrorIs(t, err, ErrCollectionNotFound)
	})

	// DELETE

	t.Run("deleting a collection succeeds", func(t *testing.T) {
		colId := "delete-test-id"
		col := entities.Collection{Id: colId, Name: "First Name"}

		_, err := sharedTestStore.CreateCollection(col)
		require.NoError(t, err)

		err = sharedTestStore.DeleteCollection(colId)
		require.NoError(t, err)

		_, err = sharedTestStore.GetCollection(colId)

		assert.ErrorIs(t, err, ErrCollectionNotFound)
	})

	t.Run("deleting a non-existent collection fails", func(t *testing.T) {
		err := sharedTestStore.DeleteCollection("no-such-collection")

		assert.ErrorIs(t, err, ErrCollectionNotFound)
	})
}

func runItemCRUDTests(
	t *testing.T,
	newStore func(t *testing.T, opts ...Option) Store,
) {
	testNow := time.Now().UTC()
	testNowUTCFunc := func() time.Time { return testNow }

	// newTestStore returns a store with predicable utility functions for testing.
	newTestStore := func() Store {
		return newStore(
			t,
			WithNowUTCFunc(testNowUTCFunc),
		)
	}

	// sharedTestStore is single shared store for use in tests that do not
	// require an empty store to assert correct returns.
	sharedTestStore := newStore(t)

	// LIST

	t.Run("list when nothing exists returns empty array", func(t *testing.T) {
		store := newTestStore()

		items, err := store.ListItemsByCollectionId("no-such-collection", nil)

		require.NoError(t, err)
		assert.Empty(t, items)
	})

	t.Run("list with default pagination succeeds", func(t *testing.T) {
		store := newStore(t)

		testCollectionId := "list-test-collection"

		// Create items in our test collection
		for i := 0; i < 3; i++ {
			_, err := store.CreateItem(
				entities.Item{
					Id:           fmt.Sprintf("primary-%v", i),
					Name:         fmt.Sprintf("PrimItem %v", i),
					CollectionId: testCollectionId,
				},
			)
			require.NoError(t, err)
		}

		// Create items in our other collection
		for i := 0; i < 3; i++ {
			_, err := store.CreateItem(
				entities.Item{
					Id:           fmt.Sprintf("secondary-%v", i),
					Name:         fmt.Sprintf("SecItem %v", i),
					CollectionId: "other-collection",
				},
			)
			require.NoError(t, err)
		}

		items, err := store.ListItemsByCollectionId(testCollectionId, nil)

		require.NoError(t, err)
		assert.Equal(t, 3, len(items))

		// Check that all items listed are from the correct collection
		for _, item := range items {
			assert.Equal(
				t,
				testCollectionId,
				item.CollectionId,
			)
		}
	})

	t.Run("list pagination returns distinct results", func(t *testing.T) {
		store := newStore(t)

		testCollectionId := "list-test-collection"

		// Create multiple items
		for i := 0; i < 8; i++ {
			_, err := store.CreateItem(
				entities.Item{
					Id:           fmt.Sprintf("primary-%v", i),
					Name:         fmt.Sprintf("PrimItem %v", i),
					CollectionId: testCollectionId,
				},
			)
			require.NoError(t, err)
		}

		// Create items in our other collection
		for i := 0; i < 3; i++ {
			_, err := store.CreateItem(
				entities.Item{
					Id:           fmt.Sprintf("secondary-%v", i),
					Name:         fmt.Sprintf("SecItem %v", i),
					CollectionId: "other-collection",
				},
			)
			require.NoError(t, err)
		}

		page1, err := store.ListItemsByCollectionId(
			testCollectionId,
			&entities.Pagination{
				Page:     1,
				PageSize: 5,
			},
		)
		require.NoError(t, err)
		assert.Equal(t, 5, len(page1))

		page2, err := store.ListItemsByCollectionId(
			testCollectionId,
			&entities.Pagination{
				Page:     2,
				PageSize: 5,
			},
		)
		require.NoError(t, err)
		assert.Equal(t, 3, len(page2))

		// Check that the returned lists are distinct
		for _, item := range page1 {
			assert.NotContains(
				t,
				page2,
				item,
				fmt.Sprintf("%s should not appear in page 2: %v", item.Id, page2),
			)
		}
	})

	// CREATE

	t.Run("create & populate metadata succeeds", func(t *testing.T) {
		store := newTestStore()

		item := entities.Item{
			Id:           "create-item-test-id",
			Name:         "Test Item",
			CollectionId: "create-col-test-id",
		}
		created, err := store.CreateItem(item)

		require.NoError(t, err)
		assert.Equal(
			t,
			entities.Item{
				Id:           item.Id,
				Name:         item.Name,
				CollectionId: item.CollectionId,
				Metadata: entities.Metadata{
					CreatedAt: testNow,
				},
			},
			created,
		)
	})

	t.Run("create dupe fails", func(t *testing.T) {
		item := entities.Item{
			Id:           "create-in-missing-id",
			Name:         "Dupe Test",
			CollectionId: "dupe-parent-collection",
		}

		_, _ = sharedTestStore.CreateItem(item)
		_, err := sharedTestStore.CreateItem(item)

		assert.ErrorIs(t, err, ErrItemAlreadyExists)
	})

	// CREATE (BATCH)

	t.Run("create batch succeeds", func(t *testing.T) {
		store := newTestStore()

		collectionId := "create-batch-test-id"

		// Build multiple items
		var items []entities.Item
		for i := 0; i < 4; i++ {
			items = append(
				items,
				entities.Item{
					Id:           fmt.Sprintf("item-%v", i),
					Name:         "Test Item",
					CollectionId: collectionId,
				},
			)
		}

		_, err := store.CreateItemBatch(items)

		require.NoError(t, err)

		listedItems, err := store.ListItemsByCollectionId(collectionId, nil)

		assert.Len(t, listedItems, 4)
	})

	t.Run("create batch skips existing items with warning", func(t *testing.T) {
		store := newTestStore()

		collectionId := "create-batch-test-id"

		// Create an initial item
		_, err := store.CreateItem(entities.Item{
			Id:           "item-1",
			Name:         "Repeated In Batch",
			CollectionId: collectionId,
		})

		// Build multiple items & create as batch
		var items []entities.Item
		for i := 0; i < 4; i++ {
			items = append(
				items,
				entities.Item{
					Id:           fmt.Sprintf("item-%v", i),
					Name:         "Test Item",
					CollectionId: collectionId,
				},
			)
		}

		_, err = store.CreateItemBatch(items)

		require.NoError(t, err)

		listedItems, err := store.ListItemsByCollectionId(collectionId, nil)

		fmt.Println(listedItems)

		assert.Len(t, listedItems, 4)
	})

	// FETCH

	t.Run("fetching a new item succeeds", func(t *testing.T) {
		store := newTestStore()

		itemId := "item-test-id"
		item := entities.Item{Id: itemId, Name: "Test", CollectionId: "parent-collection"}

		_, err := store.CreateItem(item)
		require.NoError(t, err)

		fetched, err := store.GetItem(itemId)

		require.NoError(t, err)
		assert.Equal(
			t,
			entities.Item{
				Id:           item.Id,
				Name:         item.Name,
				CollectionId: item.CollectionId,
				Metadata: entities.Metadata{
					CreatedAt: testNow,
				},
			},
			fetched,
		)
	})

	t.Run("fetching a missing items fails", func(t *testing.T) {
		_, err := sharedTestStore.GetItem("no-such-item")

		assert.ErrorIs(t, err, ErrItemNotFound)
	})

	// UPDATE

	t.Run("updating item name succeeds", func(t *testing.T) {
		store := newTestStore()

		itemId := "update-test-id"
		item := entities.Item{Id: itemId, Name: "First Name", CollectionId: "update-parent-collection"}

		_, err := store.CreateItem(item)
		require.NoError(t, err)

		updatedName := "New Name"

		_, err = store.UpdateItem(
			itemId,
			entities.ItemUpdate{Name: &updatedName},
		)
		require.NoError(t, err)

		updated, err := store.GetItem(itemId)

		require.NoError(t, err)
		assert.Equal(
			t,
			entities.Item{
				Id:           itemId,
				Name:         updatedName,
				CollectionId: item.CollectionId,
				Metadata: entities.Metadata{
					CreatedAt: testNow,
				},
			},
			updated,
		)
	})

	t.Run("updating item collectionId moves item to other collection", func(t *testing.T) {
		store := newTestStore()

		itemId := "update-test-id"
		firstCollection := "first-parent-collection"
		item := entities.Item{Id: itemId, Name: "First Name", CollectionId: firstCollection}

		_, err := store.CreateItem(item)
		require.NoError(t, err)

		updatedCollection := "new-parent-collection"

		_, err = store.UpdateItem(
			itemId,
			entities.ItemUpdate{CollectionId: &updatedCollection},
		)
		require.NoError(t, err)

		updated, err := store.GetItem(itemId)

		expected := entities.Item{
			Id:           itemId,
			Name:         item.Name,
			CollectionId: updatedCollection,
			Metadata: entities.Metadata{
				CreatedAt: testNow,
			},
		}

		require.NoError(t, err)
		assert.Equal(t, expected, updated)

		firstColItems, err := store.ListItemsByCollectionId(firstCollection, nil)
		require.NoError(t, err)
		assert.NotContains(t, firstColItems, item)

		updatedColItems, err := store.ListItemsByCollectionId(updatedCollection, nil)
		require.NoError(t, err)
		assert.Contains(t, updatedColItems, expected)
	})

	t.Run("updating a missing item fails", func(t *testing.T) {
		newName := "New Name"
		_, err := sharedTestStore.UpdateItem(
			"no-such-item",
			entities.ItemUpdate{Name: &newName},
		)

		assert.ErrorIs(t, err, ErrItemNotFound)
	})

	// DELETE

	t.Run("delete item succeeds", func(t *testing.T) {
		itemId := "delete-test-id"
		item := entities.Item{Id: itemId, Name: "First Name", CollectionId: "collection-id"}

		_, err := sharedTestStore.CreateItem(item)
		require.NoError(t, err)

		err = sharedTestStore.DeleteItem(itemId)
		require.NoError(t, err)

		_, err = sharedTestStore.GetItem(itemId)

		assert.ErrorIs(t, err, ErrItemNotFound)
	})

	t.Run("deleting a missing item fails", func(t *testing.T) {
		err := sharedTestStore.DeleteItem("no-such-item")

		assert.ErrorIs(t, err, ErrItemNotFound)
	})
}

func runCollectionAndItemTests(
	t *testing.T,
	newStore func(t *testing.T, opts ...Option) Store,
) {
	t.Run("test getting item count by collection", func(t *testing.T) {})

	t.Run("delete collection unmarks collectionId from child items", func(t *testing.T) {})
}
