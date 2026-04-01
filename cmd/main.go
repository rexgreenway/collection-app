package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

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

	storeType := storage.InMemory
	store, err := storage.StorageManager(storeType)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to initialise %q type store: %v\n", storeType, err)
		os.Exit(1)
	}

	// Establish Context
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	// Setup Err Group
	g, ctx := errgroup.WithContext(ctx)

	// Start gRPC Server inside a GO Routine
	g.Go(func() error {
		return server.StartGrpcServer(ctx, logger, store)
	})

	g.Go(func() error {
		return server.StartHTTPServer(ctx, logger, store)
	})

	// Wait for signal or error
	if err := g.Wait(); err != nil {
		logger.Errorf("Application exiting due to error: %v", err)
	}

	logger.Info("Exiting Collection Application")
}
