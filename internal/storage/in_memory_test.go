package storage

import (
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

func TestInMemoryStorage(t *testing.T) {
	newStore := func(t *testing.T, opts ...Option) Store {
		t.Helper()

		store, err := StorageFactory(IN_MEMORY, zap.NewNop().Sugar(), opts...)
		require.NoError(t, err)
		return store
	}

	t.Run("Collections", func(t *testing.T) { runCollectionCRUDTests(t, newStore) })
	t.Run("Items", func(t *testing.T) { runItemCRUDTests(t, newStore) })
}
