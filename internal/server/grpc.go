package server

import (
	"context"
	"fmt"
	"net"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "github.com/rexgreenway/collection-app/internal/gen/v1/collection"
	"github.com/rexgreenway/collection-app/internal/services/collection"
)

const GRPC ServerType = "grpc"

// GrpcServer ???
type GrpcServer struct {
	logger            *zap.SugaredLogger
	addr              string
	collectionService *collection.CollectionService
}

// NewGrpcServer ???
func NewGrpcServer(
	logger *zap.SugaredLogger,
	addr string,
	service *collection.CollectionService,
) *GrpcServer {
	return &GrpcServer{logger: logger, addr: addr, collectionService: service}
}

func (s *GrpcServer) Name() ServerType { return GRPC }

// Start ???
func (s *GrpcServer) Start(ctx context.Context) error {
	lis, err := net.Listen("tcp", s.addr)
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	// Register reflection service on gRPC server.
	reflection.Register(grpcServer)

	// Register CollectionService with the gRPC server
	// This registers the implementation of the pb.CollectionServiceServer with the
	// grpcServer created above.
	// This means that if I want to create a new implementation of the Collection Server
	// Maybe create a new version in the future... I can create a whole new implementation
	// swap it out!
	pb.RegisterCollectionServiceServer(grpcServer, s.collectionService)

	// Goroutine watches for context cancellation
	go func() {
		<-ctx.Done()
		s.logger.Info("shutting down grpc server...")
		grpcServer.GracefulStop() // stops accepting new RPCs, waits for in-flight ones to finish
	}()

	s.logger.Infof("starting grpc server on %s", s.addr)

	return grpcServer.Serve(lis)
}
