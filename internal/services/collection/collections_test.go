package collection

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/rexgreenway/collection-app/internal/gen/v1/collection"
	"github.com/rexgreenway/collection-app/internal/storage"
)

func newTestService(t *testing.T, opts ...Option) *CollectionService {
	t.Helper()

	logger := zap.NewNop().Sugar()

	store, err := storage.StorageFactory(storage.IN_MEMORY, logger)
	if err != nil {
		t.Fatalf("failed to create store: %v", err)
	}

	return NewService(logger, store, opts...)
}

func TestListCollections(t *testing.T) {
	ctx := context.Background()

	t.Run("lists nothing when empty store", func(t *testing.T) {
		svc := newTestService(t)

		resp, err := svc.ListCollections(ctx, &pb.ListCollectionsRequest{})

		require.NoError(t, err)

		assert.Empty(t, resp.Data)

		assert.EqualValues(t, 1, resp.Pagination.Page)
		assert.EqualValues(t, 10, resp.Pagination.PageSize)
	})

	t.Run("lists all created collections", func(t *testing.T) {
		svc := newTestService(t)

		for i := range 3 {
			svc.CreateCollection(ctx, &pb.CreateCollectionRequest{
				Collection: &pb.Collection{
					Name: fmt.Sprintf("Test Col %d", i),
				},
			})
		}

		resp, err := svc.ListCollections(ctx, &pb.ListCollectionsRequest{})

		require.NoError(t, err)
		assert.Len(t, resp.Data, 3)
	})

	t.Run("lists collections using pagination", func(t *testing.T) {
		svc := newTestService(t)

		for i := range 8 {
			svc.CreateCollection(ctx, &pb.CreateCollectionRequest{
				Collection: &pb.Collection{
					Name: fmt.Sprintf("Test Col %d", i),
				},
			})
		}

		// Get first page of 5 collections.
		page1, err := svc.ListCollections(ctx, &pb.ListCollectionsRequest{
			Pagination: &pb.PaginationParams{
				Page:     1,
				PageSize: 5,
			},
		})

		require.NoError(t, err)
		assert.Len(t, page1.Data, 5)

		// Get second page, with rest of the collections
		page2, err := svc.ListCollections(ctx, &pb.ListCollectionsRequest{
			Pagination: &pb.PaginationParams{
				Page:     2,
				PageSize: 5,
			},
		})

		require.NoError(t, err)
		assert.Len(t, page2.Data, 3)

		// Assert that the two pages do not overlap
		assert.NotSubset(t, page1.Data, page2.Data)
		assert.NotSubset(t, page2.Data, page1.Data)
	})

	// TESTs:
	// - test pagination
	// -
}

func TestCreateCollections(t *testing.T) {
	ctx := context.Background()

	t.Run("creates successfully", func(t *testing.T) {
		svc := newTestService(t)

		resp, err := svc.CreateCollection(ctx, &pb.CreateCollectionRequest{
			Collection: &pb.Collection{Name: "Test"},
		})

		// Require no returned exception
		require.NoError(t, err)

		// Check success code
		st := status.Convert(err)
		assert.Equal(t, codes.OK, st.Code())

		// Check content of repsonse
		assert.Equal(t, "Test", resp.Data.Name)
		assert.NotEmpty(t, resp.Data.Id)
	})

	t.Run("fails when collection id already exists", func(t *testing.T) {
		// Create a test service where same id is always generated.
		svc := newTestService(
			t,
			WithNewIdFunc(func() string { return "same-test-id" }),
		)

		// First call succeeds.
		_, err := svc.CreateCollection(ctx, &pb.CreateCollectionRequest{
			Collection: &pb.Collection{Name: "First"},
		})

		// Second call hits the same ID in store error.
		_, err = svc.CreateCollection(ctx, &pb.CreateCollectionRequest{
			Collection: &pb.Collection{Name: "Second"},
		})

		st := status.Convert(err)
		assert.Equal(t, codes.AlreadyExists, st.Code())
		assert.Contains(t, st.Message(), "already exists")
	})
}

// func TestListCollections(t *testing.T) {
// 	svc := newTestService(t)
// 	ctx := context.Background()

// 	t.Run("list when there is no collections", func(t *testing.T) {
// 		// svc.ListCollections
// 	})
// }
