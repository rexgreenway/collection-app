package server

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"time"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
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

// StartHTTPServer ???
func StartHTTPServer(ctx context.Context, logger *zap.SugaredLogger, store storage.Store) error {
	// Gateway Mux from grpc-gateway
	gwMux := runtime.NewServeMux()

	// Register the collection service with the gateway
	pb.RegisterCollectionServiceHandlerServer(ctx, gwMux, collection.NewServer(logger, store))

	// Create GIN router with v1 prefix group and attach the
	router := gin.New()

	// Desugar back to *zap.Logger for the middleware
	zapLogger := logger.Desugar()
	router.Use(ginzap.Ginzap(zapLogger, time.RFC3339, true))
	router.Use(ginzap.RecoveryWithZap(zapLogger, true))

	v1 := router.Group("/v1")
	{
		// gRPC-gateway handles collection routes (prefix is stripped for sending to gwMux)
		handler := gin.WrapH(http.StripPrefix("/v1", gwMux))
		v1.Any("/collections", handler)
		v1.Any("/collections/*path", handler)

		// Add a /v1/docs route that serves the API docs
		// v1.GET("/docs", swaggerHandler)
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
