package server

import (
	"context"
	"net"
	"net/http"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	"github.com/rexgreenway/collection-app/internal/gen/collection"
)

type collectionService struct {
	collection.UnimplementedCollectionServiceServer

	logger *zap.SugaredLogger
}

func (s *collectionService) GetCollection(
	ctx context.Context,
	req *collection.IdMessage,
) (*collection.Collection, error) {
	s.logger.Debugf("GetCollection called with id: %q", req.GetId())

	return &collection.Collection{
		Id:   req.Id,
		Name: "Test Name",
	}, nil
}

func StartServer(logger *zap.SugaredLogger) error {
	lis, err := net.Listen("tcp", "localhost:50100")
	if err != nil {
		logger.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	// Register reflection service on gRPC server.
	reflection.Register(grpcServer)

	collection.RegisterCollectionServiceServer(
		grpcServer,
		&collectionService{logger: logger},
	)

	logger.Info("starting grpc server on 50100")

	return grpcServer.Serve(lis)
}

func StartHTTPGateway(logger *zap.SugaredLogger) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	address := ":8089"

	mux := runtime.NewServeMux()

	collection.RegisterCollectionServiceHandlerServer(ctx, mux, &collectionService{logger: logger})

	s := &http.Server{
		Addr:    address,
		Handler: mux,
	}

	logger.Info("starting http gateway on 8089")

	return s.ListenAndServe()
}
