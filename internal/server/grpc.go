package server

import (
	"context"
	"fmt"
	"net"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "github.com/rexgreenway/collection-app/internal/gen/collection"
	"github.com/rexgreenway/collection-app/internal/services/collection"
	"github.com/rexgreenway/collection-app/internal/storage"
)

// StartGrpcServer
func StartGrpcServer(ctx context.Context, logger *zap.SugaredLogger, store storage.Store) error {
	lis, err := net.Listen("tcp", "localhost:50100")
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	// Register reflection service on gRPC server.
	reflection.Register(grpcServer)

	// Register CollectionService with the gRPC server
	pb.RegisterCollectionServiceServer(grpcServer, collection.NewServer(logger, store))

	// Goroutine watches for context cancellation
	go func() {
		<-ctx.Done()
		logger.Info("shutting down grpc server...")
		grpcServer.GracefulStop() // stops accepting new RPCs, waits for in-flight ones to finish
	}()

	logger.Info("starting grpc server on :50100")

	return grpcServer.Serve(lis)
}
