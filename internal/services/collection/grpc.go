package collection

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/rexgreenway/collection-app/internal/entities"
	pb "github.com/rexgreenway/collection-app/internal/gen/v1/collection"
	"github.com/rexgreenway/collection-app/internal/storage"
)

// GRPC Service Type for the collection service.
const GRPC ServiceType = "grpc"

// collectionServer is the gRPC implementation of the Collection Service server.
type grpcServer struct {
	// This adds forward compatibility to this implementation of the server
	pb.UnimplementedCollectionServiceServer

	logger *zap.SugaredLogger

	store storage.Store
}

func newGrpcServer(logger *zap.SugaredLogger, store storage.Store) *grpcServer {
	return &grpcServer{logger: logger, store: store}
}

// ------- Collection Methods -------

// ListCollections
func (s *grpcServer) ListCollections(
	ctx context.Context,
	req *pb.ListCollectionsRequest,
) (*pb.ListCollectionsResponse, error) {
	pagination := protoToPagination(req.GetPagination())

	collections, err := s.store.ListCollections(&pagination)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "ListCollections failed: %v", err)
	}

	var result []*pb.Collection
	for _, c := range collections {
		result = append(result, collectionToProto(c))
	}

	return &pb.ListCollectionsResponse{
		Data:       result,
		Pagination: paginationToProto(pagination),
	}, nil
}

// CreateCollection ???
func (s *grpcServer) CreateCollection(
	ctx context.Context,
	req *pb.CreateCollectionRequest,
) (*pb.GetCollectionResponse, error) {
	id := uuid.NewString()

	protoCollection := &pb.Collection{
		Id:   id,
		Name: req.Collection.GetName(),
	}

	_, err := s.store.CreateCollection(protoToCollection(protoCollection))
	if err != nil {
		if errors.Is(err, storage.ErrCollectionAlreadyExists) {
			return nil, status.Errorf(codes.AlreadyExists, "Collection %q already exists", id)
		}
		return nil, status.Errorf(codes.Internal, "CreateCollection %q failed: %v", id, err)
	}

	return &pb.GetCollectionResponse{
		Data: protoCollection,
	}, nil
}

// GetCollection ???
func (s *grpcServer) GetCollection(
	ctx context.Context,
	req *pb.CollectionId,
) (*pb.GetCollectionResponse, error) {
	id := req.GetId()

	collection, err := s.store.GetCollection(id)
	if err != nil {
		if errors.Is(err, storage.ErrCollectionNotFound) {
			return nil, status.Errorf(codes.NotFound, "Collection %q not found", id)
		}
		return nil, status.Errorf(codes.Internal, "GetCollection %q failed: %v", id, err)
	}

	itemCount := s.store.GetItemCountByCollection(id)

	return &pb.GetCollectionResponse{
		Data: &pb.Collection{
			Id:        collection.ID,
			Name:      collection.Name,
			ItemCount: int32(itemCount),
		},
	}, nil
}

// UpdateCollection ???
func (s *grpcServer) UpdateCollection(
	ctx context.Context,
	req *pb.UpdateCollectionRequest,
) (*pb.GetCollectionResponse, error) {
	id := req.GetId()

	collection, err := s.store.UpdateCollection(id, protoToCollection(req.GetCollection()))
	if err != nil {
		if errors.Is(err, storage.ErrCollectionNotFound) {
			return nil, status.Errorf(codes.NotFound, "Collection %q not found", id)
		}
		return nil, status.Errorf(codes.Internal, "UpdateCollection %q failed: %v", id, err)
	}

	return &pb.GetCollectionResponse{
		Data: collectionToProto(collection),
	}, nil
}

// DeleteCollection ???
func (s *grpcServer) DeleteCollection(
	ctx context.Context,
	req *pb.CollectionId,
) (*emptypb.Empty, error) {
	id := req.GetId()

	err := s.store.DeleteCollection(id)
	if err != nil {
		if errors.Is(err, storage.ErrCollectionNotFound) {
			return nil, status.Errorf(codes.NotFound, "Collection %q not found", id)
		}
		return nil, status.Errorf(codes.Internal, "DeleteCollection %q failed: %v", id, err)
	}

	return nil, nil
}

// ------- Item Methods -------

// ListItems ???
func (s *grpcServer) ListItems(
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
func (s *grpcServer) CreateItem(
	ctx context.Context,
	req *pb.CreateItemRequest,
) (*pb.GetItemResponse, error) {
	id := uuid.NewString()
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
func (s *grpcServer) CreateItems(
	ctx context.Context,
	req *pb.CreateItemsRequest,
) (*pb.ListItemsResponse, error) {
	collectionId := req.GetCollectionId()

	var itemsToCreate []entities.Item
	for _, reqItem := range req.GetItems() {
		item := entities.Item{
			Id:           uuid.NewString(),
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
func (s *grpcServer) GetItem(
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
func (s *grpcServer) UpdateItem(
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
func (s *grpcServer) DeleteItem(
	ctx context.Context,
	req *pb.CollectionItemId,
) (*pb.GetItemResponse, error) {
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
