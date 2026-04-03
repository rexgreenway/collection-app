package collection

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	pb "github.com/rexgreenway/collection-app/internal/gen/v1/collection"
	"github.com/rexgreenway/collection-app/internal/storage"
)

type collectionServer struct {
	// This adds forward compatibility to this
	// implementation of the server
	pb.UnimplementedCollectionServiceServer

	logger *zap.SugaredLogger

	store storage.Store

	cancel context.CancelFunc
}

func (s *collectionServer) ListCollections(
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

func (s *collectionServer) CreateCollection(
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

func (s *collectionServer) GetCollection(
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

	result := collectionToProto(collection)

	return &pb.GetCollectionResponse{
		Data: result,
	}, nil
}

func (s *collectionServer) UpdateCollection(
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

func (s *collectionServer) DeleteCollection(
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

func NewServer(logger *zap.SugaredLogger, store storage.Store) *collectionServer {
	return &collectionServer{logger: logger, store: store}
}
