package server

import (
	"net"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "github.com/rexgreenway/collection-app/internal/gen/collection"
	"github.com/rexgreenway/collection-app/internal/services/collection"
	"github.com/rexgreenway/collection-app/internal/storage"
)

// START GRPC SERVER
func StartGrpcServer(logger *zap.SugaredLogger, store storage.Store) error {
	lis, err := net.Listen("tcp", "localhost:50100")
	if err != nil {
		logger.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	// Register reflection service on gRPC server.
	reflection.Register(grpcServer)

	// Register CollectionService with the gRPC server
	pb.RegisterCollectionServiceServer(grpcServer, collection.NewServer(logger, store))

	logger.Info("starting grpc server on 50100")

	return grpcServer.Serve(lis)
}
