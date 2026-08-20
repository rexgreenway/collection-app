package collection

import (
	"context"
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	pb "github.com/rexgreenway/collection-app/internal/gen/v1/collection"
	"github.com/rexgreenway/collection-app/internal/storage"
)

// ------- Collection Methods -------

// ListCollections
func (s *CollectionService) ListCollections(
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
func (s *CollectionService) CreateCollection(
	ctx context.Context,
	req *pb.CreateCollectionRequest,
) (*pb.GetCollectionResponse, error) {
	id := s.utils.NewId()

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
func (s *CollectionService) GetCollection(
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
			Id:        collection.Id,
			Name:      collection.Name,
			ItemCount: int32(itemCount),
		},
	}, nil
}

// UpdateCollection ???
func (s *CollectionService) UpdateCollection(
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
func (s *CollectionService) DeleteCollection(
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
