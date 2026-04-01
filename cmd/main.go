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
	storeType := storage.InMemory
	store, err := storage.StorageManager(storeType)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to initialise %q type store: %v\n", storeType, err)
		os.Exit(1)
	}

	// Establish Context & Err Group
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	g, ctx := errgroup.WithContext(ctx)

	// Start Servers inside go routines
	g.Go(func() error {
		return server.StartGrpcServer(ctx, logger, store)
	})

	g.Go(func() error {
		return server.StartHTTPServer(ctx, logger, store)
	})

	if err := g.Wait(); err != nil {
		logger.Errorf("Application exiting due to error: %v", err)
	}

	logger.Info("Exiting Collection Application")
}
