package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/rexgreenway/collection-app/internal/logger"
	"github.com/rexgreenway/collection-app/internal/server"
	"github.com/rexgreenway/collection-app/internal/services/collection"
	"github.com/rexgreenway/collection-app/internal/storage"
	"golang.org/x/sync/errgroup"
)

func main() {
	environment := os.Getenv("ENVIRONMENT")

	logger, err := logger.FromConfig(&logger.Config{Environment: environment})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to initialise logger: %v\n", err)
		os.Exit(1)
	}
	defer logger.Sync() // now runs when main() exits — flushes any buffered log entries

	logger.Info("Starting Collection Application")

	logger.Info("Setting Up Environment")
	// Set if Gin is Prod or not?
	if environment == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Set up persistent Store
	// Read this from a env var at some point
	storeType := storage.IN_MEMORY
	store, err := storage.StorageFactory(storeType, logger)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to initialise %q type Store: %v\n", storeType, err)
		os.Exit(1)
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
