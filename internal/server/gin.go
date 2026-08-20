package server

import (
	"context"
	"net/http"
	"time"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.uber.org/zap"

	pb "github.com/rexgreenway/collection-app/internal/gen/v1/collection"
	"github.com/rexgreenway/collection-app/internal/services/collection"
)

const GIN ServerType = "grpc"

type GinServer struct {
	logger  *zap.SugaredLogger
	addr    string
	service *collection.CollectionService
}

func NewGinServer(
	logger *zap.SugaredLogger,
	addr string,
	service *collection.CollectionService,
) *GinServer {
	return &GinServer{logger: logger, addr: addr, service: service}
}

func (s *GinServer) Name() ServerType { return GIN }

func (s *GinServer) Start(ctx context.Context) error {
	// Gateway Mux from grpc-gateway
	gwMux := runtime.NewServeMux()

	// Register the collection service with the gateway
	if err := pb.RegisterCollectionServiceHandlerServer(ctx, gwMux, s.service); err != nil {
		return err
	}

	// Create GIN router
	router := gin.New()

	// zapLogger is the desugared zap.Logger required by the ginzap middleware.
	zapLogger := s.logger.Desugar()
	router.Use(ginzap.Ginzap(zapLogger, time.RFC3339, true))
	router.Use(ginzap.RecoveryWithZap(zapLogger, true))

	// v1 is the group of endpoints prefixed with `/v1`
	v1 := router.Group("/v1")
	{
		// gRPC-gateway handles collection routes (prefix is stripped for sending to gwMux)
		handler := gin.WrapH(http.StripPrefix("/v1", gwMux))
		v1.Any("/collections", handler)
		v1.Any("/collections/*path", handler)

		// Add a /v1/docs route that serves the API docs
		// v1.GET("/docs", swaggerHandler)
	}

	// srv is the HTTP server that serves the GIN router.
	srv := &http.Server{
		Addr:    s.addr,
		Handler: router,
	}

	// Goroutine watches for context cancellation
	go func() {
		<-ctx.Done()
		s.logger.Info("shutting down http server...")
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := srv.Shutdown(shutdownCtx); err != nil {
			s.logger.Errorf("http server forced shutdown: %v", err)
		}
	}()

	s.logger.Infof("starting http gateway on %s", s.addr)

	// ListenAndServe returns http.ErrServerClosed after Shutdown() completes — that's expected
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err
	}
	return nil
}
