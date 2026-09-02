package collection

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/rexgreenway/collection-app/internal/entities"
	pb "github.com/rexgreenway/collection-app/internal/gen/v1/collection"
	"github.com/rexgreenway/collection-app/internal/storage"
)

func TestListCollections(t *testing.T) {
	ctx := context.Background()

	t.Run("lists nothing when empty store", func(t *testing.T) {
		svc := newTestService(t)

		resp, err := svc.ListCollections(ctx, &pb.ListCollectionsRequest{})

		require.NoError(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.OK, st.Code())

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

		st := status.Convert(err)
		assert.Equal(t, codes.OK, st.Code())

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

		st := status.Convert(err)
		assert.Equal(t, codes.OK, st.Code())

		assert.Len(t, page2.Data, 3)

		// Assert that the two pages do not overlap
		assert.NotSubset(t, page1.Data, page2.Data)
		assert.NotSubset(t, page2.Data, page1.Data)
	})

	t.Run("list collections internal store failure", func(t *testing.T) {
		svc, store := newMockedService(t)

		// Force the dependency to fail.
		store.EXPECT().
			ListCollections(mock.AnythingOfType("*entities.Pagination")).
			Return(nil, errors.New("boom"))

		resp, err := svc.ListCollections(ctx, &pb.ListCollectionsRequest{})

		require.Error(t, err)
		assert.Nil(t, resp)

		st := status.Convert(err)
		assert.Equal(t, codes.Internal, st.Code())
	})
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

	t.Run("create collection internal store failure", func(t *testing.T) {
		svc, store := newMockedService(t)

		// Force the dependency to fail.
		store.EXPECT().
			CreateCollection(mock.AnythingOfType("entities.Collection")).
			Return(entities.Collection{}, errors.New("boom"))

		resp, err := svc.CreateCollection(ctx, &pb.CreateCollectionRequest{})

		require.Error(t, err)
		assert.Nil(t, resp)

		st := status.Convert(err)
		assert.Equal(t, codes.Internal, st.Code())
	})
}

func TestGetCollection(t *testing.T) {
	ctx := context.Background()

	t.Run("get collection successfully", func(t *testing.T) {
		svc := newTestService(t)

		newCol := pb.Collection{
			Name: "Test Collection",
		}

		createResp, err := svc.CreateCollection(ctx, &pb.CreateCollectionRequest{
			Collection: &newCol,
		})
		require.NoError(t, err)

		resp, err := svc.GetCollection(ctx, &pb.CollectionId{Id: createResp.Data.Id})

		st := status.Convert(err)
		assert.Equal(t, codes.OK, st.Code())

		// Check content of response
		assert.NotEmpty(t, resp.Data.Id)
		assert.Equal(t, newCol.Name, resp.Data.Name)
		assert.EqualValues(t, 0, resp.Data.ItemCount)
	})

	t.Run("get collection doesn't exists", func(t *testing.T) {
		svc, store := newMockedService(t)

		// Return the sentinel error the service knows how to classify.
		store.EXPECT().
			GetCollection("missing-id").
			Return(entities.Collection{}, storage.ErrCollectionNotFound)

		resp, err := svc.GetCollection(ctx, &pb.CollectionId{Id: "missing-id"})

		require.Error(t, err)
		assert.Nil(t, resp)

		st := status.Convert(err)
		assert.Equal(t, codes.NotFound, st.Code())

		// GetItemCountByCollection must not be reached once GetCollection fails.
		store.AssertNotCalled(t, "GetItemCountByCollection", "missing-id")
	})

	t.Run("get collection internal store failure", func(t *testing.T) {
		svc, store := newMockedService(t)

		// Force the dependency to fail.
		store.EXPECT().
			GetCollection(mock.AnythingOfType("string")).
			Return(entities.Collection{}, errors.New("boom"))

		resp, err := svc.GetCollection(ctx, &pb.CollectionId{})

		require.Error(t, err)
		assert.Nil(t, resp)

		st := status.Convert(err)
		assert.Equal(t, codes.Internal, st.Code())
	})
}

func TestUpdateCollection(t *testing.T) {
	ctx := context.Background()

	t.Run("update collection doesn't exists", func(t *testing.T) {
		svc, store := newMockedService(t)

		id := "missing-id"

		// Return the sentinel error the service knows how to classify.
		store.EXPECT().
			UpdateCollection(id, mock.AnythingOfType("entities.CollectionUpdate")).
			Return(entities.Collection{}, storage.ErrCollectionNotFound)

		resp, err := svc.UpdateCollection(
			ctx,
			&pb.UpdateCollectionRequest{
				Id: id,
				Collection: &pb.Collection{
					Name: "New Name",
				},
			},
		)

		require.Error(t, err)
		assert.Nil(t, resp)

		st := status.Convert(err)
		assert.Equal(t, codes.NotFound, st.Code())

		// GetItemCountByCollection must not be reached once GetCollection fails.
		store.AssertNotCalled(t, "GetItemCountByCollection", "missing-id")
	})

	t.Run("update collection internal store failure", func(t *testing.T) {
		svc, store := newMockedService(t)

		// Force the dependency to fail.
		store.EXPECT().
			UpdateCollection(
				mock.AnythingOfType("string"),
				mock.AnythingOfType("entities.CollectionUpdate"),
			).
			Return(entities.Collection{}, errors.New("boom"))

		resp, err := svc.UpdateCollection(
			ctx,
			&pb.UpdateCollectionRequest{
				Collection: &pb.Collection{},
			},
		)

		require.Error(t, err)
		assert.Nil(t, resp)

		st := status.Convert(err)
		assert.Equal(t, codes.Internal, st.Code())
	})
}

func TestDeleteCollection(t *testing.T) {
	ctx := context.Background()

	t.Run("delete collection successfully", func(t *testing.T) {
		svc := newTestService(t)

		createResp, err := svc.CreateCollection(ctx, &pb.CreateCollectionRequest{
			Collection: &pb.Collection{
				Name: "Test Collection",
			},
		})
		require.NoError(t, err)

		_, err = svc.GetCollection(ctx, &pb.CollectionId{Id: createResp.Data.Id})
		require.NoError(t, err)

		resp, err := svc.DeleteCollection(ctx, &pb.CollectionId{Id: createResp.Data.Id})

		st := status.Convert(err)
		assert.Equal(t, codes.OK, st.Code())
		assert.Nil(t, resp)

		_, err = svc.GetCollection(ctx, &pb.CollectionId{Id: createResp.Data.Id})

		assert.Nil(t, resp)

		st = status.Convert(err)
		assert.Equal(t, codes.NotFound, st.Code())
	})
	t.Run("delete fails when collection not found", func(t *testing.T) {
		svc, store := newMockedService(t)

		// Return the sentinel error the service knows how to classify.
		store.EXPECT().
			DeleteCollection(mock.AnythingOfType("string")).
			Return(storage.ErrCollectionNotFound)

		resp, err := svc.DeleteCollection(ctx, &pb.CollectionId{})

		require.Error(t, err)
		assert.Nil(t, resp)

		st := status.Convert(err)
		assert.Equal(t, codes.NotFound, st.Code())
	})
}
