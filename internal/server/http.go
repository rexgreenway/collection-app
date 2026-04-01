package server

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.uber.org/zap"

	pb "github.com/rexgreenway/collection-app/internal/gen/collection"
	"github.com/rexgreenway/collection-app/internal/services/collection"
	"github.com/rexgreenway/collection-app/internal/storage"
)

// StartHTTPServer ???
func StartHTTPServer(logger *zap.SugaredLogger, store storage.Store) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Gateway Mux from grpc-gateway
	gwMux := runtime.NewServeMux()

	// Register the collection service with the gateway
	pb.RegisterCollectionServiceHandlerServer(ctx, gwMux, collection.NewServer(logger, store))

	// Create GIN router with v1 prefix group and attach the
	router := gin.Default()
	v1 := router.Group("/v1")
	{
		// gRPC-gateway handles collection routes (prefix is stripped for sending to gwMux)
		handler := gin.WrapH(http.StripPrefix("/v1", gwMux))
		v1.Any("/collections", handler)
		v1.Any("/collections/*path", handler)
	}

	logger.Info("starting http gateway on 8089")

	return router.Run(":8089")
}
