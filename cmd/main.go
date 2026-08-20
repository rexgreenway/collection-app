package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/rexgreenway/collection-app/internal/config"
	"github.com/rexgreenway/collection-app/internal/logger"
	"github.com/rexgreenway/collection-app/internal/server"
	"github.com/rexgreenway/collection-app/internal/services/collection"
	"github.com/rexgreenway/collection-app/internal/storage"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
)

func main() {
	// bootLog is an ephemeral logger used during application configuration &
	//  startup before the main logger is configured.
	bootLog := zap.Must(zap.NewDevelopment()).Sugar()

	cfg := config.FromEnv(bootLog)

	logger, err := logger.FromConfig(&cfg)
	if err != nil {
		bootLog.Fatalf("Failed to initialise logger: %v\n", err)
	}
	bootLog.Sync() // flush bootLog upon successful set up of main logger
	defer logger.Sync()

	// Set if Gin is Prod or not?
	// THIS ONLY NEEDS TO BE SET IF GIN IS ACTUALLY EVEN BEING RUN
	if cfg.Environment == config.PRODUCTION {
		gin.SetMode(gin.ReleaseMode)
	}

	store, err := storage.StorageFactory(cfg.Store, logger)
	if err != nil {
		logger.Fatalf("Failed to initialise %q type Store: %v\n", cfg.Store, err)
	}

	// Instantiate collection service
	collectionServiceType := collection.GRPC
	// Implementations should be config / environment driven
	collectionService, err := collection.CollectionServiceFactory(collectionServiceType, logger, store)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to initialise %q type Collection Service: %v\n", storeType, err)
		os.Exit(1)
	}

	// Establish Context & Err Group
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	g, ctx := errgroup.WithContext(ctx)

	// Start Servers inside go routines
	// If a server is started should be driven from config!!
	g.Go(func() error {
		return server.StartGrpcServer(ctx, logger, store, collectionService)
	})

	g.Go(func() error {
		return server.StartHTTPServer(ctx, logger, store, collectionService)
	})

	if err := g.Wait(); err != nil {
		logger.Errorf("Application exiting due to error: %v", err)
	}

	logger.Info("Exiting Collection Application")
}
