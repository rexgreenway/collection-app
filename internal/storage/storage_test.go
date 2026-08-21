package storage

import (
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

	// newTestStore returns a store with the above defined predicable utility
	// functions for testing.
	newTestStore := func() Store {
		return newStore(
			t,
			WithNowUTCFunc(testNowUTCFunc),
		)
	}

	// sharedTestStore is single shared store for use in tests that do not
	// require an empty store to assert correct returns.
	sharedTestStore := newStore(t)

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

	// LIST

	t.Run("list when nothing exists returns empty array", func(t *testing.T) {
		// Create Empty Store
		store := newStore(t)

		collections, err := store.ListCollections(nil)

		require.NoError(t, err)
		assert.Empty(t, collections)
	})

	t.Run("list with no pagination succeeds", func(t *testing.T) {
		// Create Empty Store
		store := newStore(t)

		// Create multiple collections
		for i := 0; i < 3; i++ {
			_, err := store.CreateCollection(
				entities.Collection{
					Id:   "col" + string(rune(i)),
					Name: "Collection " + string(rune(i)),
				},
			)
			require.NoError(t, err)
		}

		collections, err := store.ListCollections(nil)

		require.NoError(t, err)
		assert.Equal(t, 3, len(collections))
	})

	t.Run("list with pagination succeeds", func(t *testing.T) {
		// Create Empty Store
		store := newStore(t)

		// Create multiple collections
		for i := 0; i < 8; i++ {
			_, err := store.CreateCollection(
				entities.Collection{
					Id:   "col" + string(rune(i)),
					Name: "Collection " + string(rune(i)),
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
			assert.NotContains(t, page2, item)
		}
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

	t.Run("updating a collection succeeds", func(t *testing.T) {
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

func runItemCRUDTests(t *testing.T, store Store) { /* ... */ }
