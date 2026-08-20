package collection

import (
	"context"
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/rexgreenway/collection-app/internal/entities"
	pb "github.com/rexgreenway/collection-app/internal/gen/v1/collection"
	"github.com/rexgreenway/collection-app/internal/storage"
)

// ------- Item Methods -------

// ListItems ???
func (s *CollectionService) ListItems(
	ctx context.Context,
	req *pb.ListItemsRequest,
) (*pb.ListItemsResponse, error) {
	pagination := protoToPagination(req.GetPagination())

	collectionId := req.GetCollectionId()

	items, err := s.store.ListItemsByCollectionId(collectionId, &pagination)
	if err != nil {
		if errors.Is(err, storage.ErrCollectionNotFound) {
			return nil, status.Errorf(codes.NotFound, "Collection %q not found", collectionId)
		}
		return nil, status.Errorf(codes.Internal, "ListItemsByCollectionId failed: %v", err)
	}

	var result []*pb.Item
	for _, item := range items {
		result = append(result, itemToProto(item))
	}

	return &pb.ListItemsResponse{
		Data:       result,
		Pagination: paginationToProto(pagination),
	}, nil
}

// CreateItem ???
func (s *CollectionService) CreateItem(
	ctx context.Context,
	req *pb.CreateItemRequest,
) (*pb.GetItemResponse, error) {
	id := s.utils.NewId()
	collectionId := req.GetCollectionId()

	// Create the actual item
	item, err := s.store.CreateItem(entities.Item{
		Id:           id,
		Name:         req.Item.GetName(),
		CollectionId: collectionId,
	})
	if err != nil {
		if errors.Is(err, storage.ErrCollectionNotFound) {
			return nil, status.Errorf(codes.NotFound, "Collection %q not found", collectionId)
		}
		if errors.Is(err, storage.ErrItemAlreadyExists) {
			return nil, status.Errorf(codes.AlreadyExists, "Item %q already exists", id)
		}
		return nil, status.Errorf(codes.Internal, "CreateItem %q failed: %v", id, err)
	}

	return &pb.GetItemResponse{
		Data: itemToProto(item),
	}, nil
}

// CreateItem ???
func (s *CollectionService) CreateItems(
	ctx context.Context,
	req *pb.CreateItemsRequest,
) (*pb.ListItemsResponse, error) {
	collectionId := req.GetCollectionId()

	var itemsToCreate []entities.Item
	for _, reqItem := range req.GetItems() {
		item := entities.Item{
			Id:           s.utils.NewId(),
			Name:         reqItem.GetName(),
			CollectionId: collectionId,
		}

		itemsToCreate = append(itemsToCreate, item)
	}

	items, err := s.store.CreateItemBatchByCollectionId(collectionId, itemsToCreate)
	if err != nil {
		if errors.Is(err, storage.ErrCollectionNotFound) {
			return nil, status.Errorf(codes.NotFound, "Collection %q not found", collectionId)
		}
		return nil, status.Errorf(codes.Internal, "CreateItemBatchByCollectionId %q failed: %v", collectionId, err)
	}

	var result []*pb.Item
	for _, item := range items {
		result = append(result, itemToProto(item))
	}

	return &pb.ListItemsResponse{
		Data: result,
	}, nil
}

// GetItem ???
func (s *CollectionService) GetItem(
	ctx context.Context,
	req *pb.CollectionItemId,
) (*pb.GetItemResponse, error) {
	id := req.GetId()
	collectionId := req.GetCollectionId()

	item, err := s.store.GetItem(collectionId, id)
	if err != nil {
		if errors.Is(err, storage.ErrCollectionNotFound) {
			return nil, status.Errorf(codes.NotFound, "Collection %q not found", collectionId)
		}
		if errors.Is(err, storage.ErrItemNotFound) {
			return nil, status.Errorf(codes.NotFound, "Item %q not found", id)
		}
		return nil, status.Errorf(codes.Internal, "GetItem %q failed: %v", id, err)
	}

	return &pb.GetItemResponse{
		Data: itemToProto(item),
	}, nil
}

// UpdateItem ???
func (s *CollectionService) UpdateItem(
	ctx context.Context,
	req *pb.UpdateItemRequest,
) (*pb.GetItemResponse, error) {
	id := req.GetId()
	collectionId := req.GetCollectionId()

	item, err := s.store.UpdateItem(collectionId, id, protoToItem(req.GetItem()))
	if err != nil {
		if errors.Is(err, storage.ErrCollectionNotFound) {
			return nil, status.Errorf(codes.NotFound, "Collection %q not found", collectionId)
		}
		if errors.Is(err, storage.ErrItemNotFound) {
			return nil, status.Errorf(codes.NotFound, "Item %q not found", id)
		}
		return nil, status.Errorf(codes.Internal, "UpdateCollection %q failed: %v", id, err)
	}

	return &pb.GetItemResponse{
		Data: itemToProto(item),
	}, nil
}

// DeleteItem ???
func (s *CollectionService) DeleteItem(
	ctx context.Context,
	req *pb.CollectionItemId,
) (*emptypb.Empty, error) {
	id := req.GetId()
	collectionId := req.GetCollectionId()

	err := s.store.DeleteItem(collectionId, id)
	if err != nil {
		if errors.Is(err, storage.ErrCollectionNotFound) {
			return nil, status.Errorf(codes.NotFound, "Collection %q not found", id)
		}
		if errors.Is(err, storage.ErrItemNotFound) {
			return nil, status.Errorf(codes.NotFound, "Item %q not found", id)
		}
		return nil, status.Errorf(codes.Internal, "DeleteCollection %q failed: %v", id, err)
	}

	return nil, nil
}
