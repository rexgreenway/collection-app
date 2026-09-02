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

func TestListItems(t *testing.T) {
	ctx := context.Background()

	t.Run("lists nothing when empty store", func(t *testing.T) {
		svc := newTestService(t)

		resp, err := svc.ListItems(ctx, &pb.ListItemsRequest{
			CollectionId: "no-items-col",
		})

		require.NoError(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.OK, st.Code())

		assert.Empty(t, resp.Data)

		assert.EqualValues(t, 1, resp.Pagination.Page)
		assert.EqualValues(t, 10, resp.Pagination.PageSize)
	})

	t.Run("lists all items in a collection", func(t *testing.T) {
		svc := newTestService(t)

		createColResp, err := svc.CreateCollection(ctx, &pb.CreateCollectionRequest{
			Collection: &pb.Collection{
				Name: "new-collection",
			},
		})
		require.NoError(t, err)

		collectionId := createColResp.Data.Id

		// CreateItems in Collection
		for i := range 3 {
			svc.CreateItem(ctx, &pb.CreateItemRequest{
				CollectionId: collectionId,
				Item: &pb.Item{
					Name: fmt.Sprintf("Test Col %d", i),
				},
			})
		}

		resp, err := svc.ListItems(ctx, &pb.ListItemsRequest{
			CollectionId: collectionId,
		})

		require.NoError(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.OK, st.Code())

		assert.Len(t, resp.Data, 3)
	})

	t.Run("lists items using pagination", func(t *testing.T) {
		svc := newTestService(t)

		createColResp, err := svc.CreateCollection(ctx, &pb.CreateCollectionRequest{
			Collection: &pb.Collection{
				Name: "new-collection",
			},
		})
		require.NoError(t, err)

		collectionId := createColResp.Data.Id

		// CreateItems in Collection
		for i := range 8 {
			svc.CreateItem(ctx, &pb.CreateItemRequest{
				CollectionId: collectionId,
				Item: &pb.Item{
					Name: fmt.Sprintf("Test Col %d", i),
				},
			})
		}

		page1, err := svc.ListItems(ctx, &pb.ListItemsRequest{
			CollectionId: collectionId,
			Pagination: &pb.PaginationParams{
				Page:     1,
				PageSize: 5,
			},
		})

		require.NoError(t, err)
		assert.Len(t, page1.Data, 5)

		page2, err := svc.ListItems(ctx, &pb.ListItemsRequest{
			CollectionId: collectionId,
			Pagination: &pb.PaginationParams{
				Page:     2,
				PageSize: 5,
			},
		})

		require.NoError(t, err)
		assert.Len(t, page2.Data, 3)

		st := status.Convert(err)
		assert.Equal(t, codes.OK, st.Code())

		// Assert that the two pages do not overlap
		assert.NotSubset(t, page1.Data, page2.Data)
		assert.NotSubset(t, page2.Data, page1.Data)
	})

	t.Run("list items internal store failure", func(t *testing.T) {
		svc, store := newMockedService(t)

		collectionId := "fake-col-id"

		// Force the dependency to fail.
		store.EXPECT().
			ListItemsByCollectionId(collectionId, mock.AnythingOfType("*entities.Pagination")).
			Return(nil, errors.New("boom"))

		resp, err := svc.ListItems(ctx, &pb.ListItemsRequest{
			CollectionId: collectionId,
		})

		require.Error(t, err)
		assert.Nil(t, resp)

		st := status.Convert(err)
		assert.Equal(t, codes.Internal, st.Code())
	})
}

func TestCreateItems(t *testing.T) {
	ctx := context.Background()

	t.Run("creates default item successfully", func(t *testing.T) {
		svc := newTestService(t)

		createColResp, err := svc.CreateCollection(ctx, &pb.CreateCollectionRequest{
			Collection: &pb.Collection{
				Name: "new-collection",
			},
		})
		require.NoError(t, err)

		resp, err := svc.CreateItem(ctx, &pb.CreateItemRequest{
			CollectionId: createColResp.Data.Id,
		})
		require.NoError(t, err)

		// Check success code
		st := status.Convert(err)
		assert.Equal(t, codes.OK, st.Code())

		// Check content of repsonse
		assert.NotEmpty(t, resp.Data.Id)
		assert.Equal(t, entities.DEFAULT_ITEM_NAME, resp.Data.Name)
		assert.Equal(t, createColResp.Data.Id, resp.Data.CollectionId)
	})

	t.Run("creates defined item successfully", func(t *testing.T) {
		svc := newTestService(t)

		createColResp, err := svc.CreateCollection(ctx, &pb.CreateCollectionRequest{
			Collection: &pb.Collection{
				Name: "new-collection",
			},
		})
		require.NoError(t, err)

		collectionId := createColResp.Data.Id

		newItem := pb.Item{
			Name: "Test Item",
		}

		resp, err := svc.CreateItem(ctx, &pb.CreateItemRequest{
			CollectionId: collectionId,
			Item:         &newItem,
		})
		require.NoError(t, err)

		// Check success code
		st := status.Convert(err)
		assert.Equal(t, codes.OK, st.Code())

		// Check content of response
		assert.NotEmpty(t, resp.Data.Id)
		assert.Equal(t, newItem.Name, resp.Data.Name)
		assert.Equal(t, collectionId, resp.Data.CollectionId)
	})

	t.Run("fails when item id already exists", func(t *testing.T) {
		// Create a test service where same id is always generated.
		svc := newTestService(
			t,
			WithNewIdFunc(func() string { return "same-test-id" }),
		)

		// First call succeeds.
		_, err := svc.CreateItem(ctx, &pb.CreateItemRequest{
			Item: &pb.Item{Name: "First"},
		})

		// First call succeeds.
		_, err = svc.CreateItem(ctx, &pb.CreateItemRequest{
			Item: &pb.Item{Name: "Second"},
		})

		st := status.Convert(err)
		assert.Equal(t, codes.AlreadyExists, st.Code())
		assert.Contains(t, st.Message(), "already exists")
	})

	t.Run("create item internal store failure", func(t *testing.T) {
		svc, store := newMockedService(t)

		// Force the dependency to fail.
		store.EXPECT().
			CreateItem(mock.AnythingOfType("entities.Item")).
			Return(entities.Item{}, errors.New("boom"))

		resp, err := svc.CreateItem(ctx, &pb.CreateItemRequest{})

		require.Error(t, err)
		assert.Nil(t, resp)

		st := status.Convert(err)
		assert.Equal(t, codes.Internal, st.Code())
	})
}

func TestGetItem(t *testing.T) {
	ctx := context.Background()

	t.Run("get item successfully", func(t *testing.T) {
		svc := newTestService(t)

		collectionId := "test-col-id"

		newItem := pb.Item{
			Name: "Test Item",
		}

		createResp, err := svc.CreateItem(ctx, &pb.CreateItemRequest{
			CollectionId: collectionId,
			Item:         &newItem,
		})
		require.NoError(t, err)

		resp, err := svc.GetItem(ctx, &pb.CollectionItemId{
			Id:           createResp.Data.Id,
			CollectionId: collectionId,
		})

		st := status.Convert(err)
		assert.Equal(t, codes.OK, st.Code())

		// Check content of response
		assert.NotEmpty(t, resp.Data.Id)
		assert.Equal(t, newItem.Name, resp.Data.Name)
		assert.Equal(t, collectionId, resp.Data.CollectionId)
	})

	t.Run("get item doesn't exists", func(t *testing.T) {
		svc, store := newMockedService(t)

		// Return the sentinel error the service knows how to classify.
		store.EXPECT().
			GetItem("missing-id").
			Return(entities.Item{}, storage.ErrItemNotFound)

		resp, err := svc.GetItem(ctx, &pb.CollectionItemId{Id: "missing-id", CollectionId: "parent-col-id"})

		require.Error(t, err)
		assert.Nil(t, resp)

		st := status.Convert(err)
		assert.Equal(t, codes.NotFound, st.Code())
	})

	t.Run("get item internal store failure", func(t *testing.T) {
		svc, store := newMockedService(t)

		// Force the dependency to fail.
		store.EXPECT().
			GetItem(mock.AnythingOfType("string")).
			Return(entities.Item{}, errors.New("boom"))

		resp, err := svc.GetItem(ctx, &pb.CollectionItemId{})

		require.Error(t, err)
		assert.Nil(t, resp)

		st := status.Convert(err)
		assert.Equal(t, codes.Internal, st.Code())
	})
}

func TestUpdateItem(t *testing.T) {
	ctx := context.Background()

	t.Run("update item doesn't exists", func(t *testing.T) {
		svc, store := newMockedService(t)

		// Return the sentinel error the service knows how to classify.
		store.EXPECT().
			UpdateItem(mock.AnythingOfType("string"), mock.AnythingOfType("entities.ItemUpdate")).
			Return(entities.Item{}, storage.ErrItemNotFound)

		resp, err := svc.UpdateItem(ctx, &pb.UpdateItemRequest{})

		require.Error(t, err)
		assert.Nil(t, resp)

		st := status.Convert(err)
		assert.Equal(t, codes.NotFound, st.Code())
	})

	t.Run("update item internal store failure", func(t *testing.T) {
		svc, store := newMockedService(t)

		// Force the dependency to fail.
		store.EXPECT().
			UpdateItem(
				mock.AnythingOfType("string"),
				mock.AnythingOfType("entities.ItemUpdate"),
			).
			Return(entities.Item{}, errors.New("boom"))

		resp, err := svc.UpdateItem(ctx, &pb.UpdateItemRequest{})

		require.Error(t, err)
		assert.Nil(t, resp)

		st := status.Convert(err)
		assert.Equal(t, codes.Internal, st.Code())
	})
}

func TestDeleteItem(t *testing.T) {
	ctx := context.Background()

	t.Run("delete item successfully", func(t *testing.T) {
		svc := newTestService(t)

		collectionId := "test-col-id"

		newItem := pb.Item{
			Name: "Test Item",
		}

		createResp, err := svc.CreateItem(ctx, &pb.CreateItemRequest{
			CollectionId: collectionId,
			Item:         &newItem,
		})
		require.NoError(t, err)

		_, err = svc.GetItem(ctx, &pb.CollectionItemId{
			Id:           createResp.Data.Id,
			CollectionId: collectionId,
		})
		require.NoError(t, err)

		resp, err := svc.DeleteItem(ctx, &pb.CollectionItemId{
			Id:           createResp.Data.Id,
			CollectionId: collectionId,
		})

		st := status.Convert(err)
		assert.Equal(t, codes.OK, st.Code())
		assert.Nil(t, resp)

		_, err = svc.GetItem(ctx, &pb.CollectionItemId{
			Id:           createResp.Data.Id,
			CollectionId: collectionId,
		})

		assert.Nil(t, resp)

		st = status.Convert(err)
		assert.Equal(t, codes.NotFound, st.Code())
	})
	t.Run("delete fails when item not found", func(t *testing.T) {
		svc, store := newMockedService(t)

		// Return the sentinel error the service knows how to classify.
		store.EXPECT().
			DeleteItem(mock.AnythingOfType("string")).
			Return(storage.ErrItemNotFound)

		resp, err := svc.DeleteItem(ctx, &pb.CollectionItemId{})

		require.Error(t, err)
		assert.Nil(t, resp)

		st := status.Convert(err)
		assert.Equal(t, codes.NotFound, st.Code())
	})
}
