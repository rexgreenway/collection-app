package server

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.uber.org/zap"

	pb "github.com/rexgreenway/collection-app/internal/gen/collection"
	"github.com/rexgreenway/collection-app/internal/services/collection"
	"github.com/rexgreenway/collection-app/internal/storage"
)

// StartHTTPServer ???
func StartHTTPServer(ctx context.Context, logger *zap.SugaredLogger, store storage.Store) error {
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

	// Create an http.Server instead of using router.Run()
	srv := &http.Server{
		Addr:    ":8089",
		Handler: router,
	}

	// Goroutine watches for context cancellation
	go func() {
		<-ctx.Done()
		logger.Info("shutting down http server...")
		// Give in-flight requests 5 seconds to finish
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := srv.Shutdown(shutdownCtx); err != nil {
			logger.Errorf("http server forced shutdown: %v", err)
		}
	}()

	logger.Info("starting http gateway on :8089")

	// ListenAndServe returns http.ErrServerClosed after Shutdown() completes — that's expected
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err
	}
	return nil
}
