package storage

import (
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

func TestInMemoryStorage(t *testing.T) {
	newStore := func(t *testing.T) Store {
		store, err := newInMemoryStorage(zap.NewNop().Sugar())
		require.NoError(t, err)
		return store
	}

	t.Run("Collections", func(t *testing.T) { runCollectionCRUDTests(t, newStore) })
	// t.Run("Items", func(t *testing.T) { runItemCRUDTests(t, store) })
}
