package main

import (
	"os"

	"github.com/rexgreenway/collection-app/internal/logger"
	"github.com/rexgreenway/collection-app/internal/server"
	"github.com/rexgreenway/collection-app/internal/storage"
)

func main() {
	environment := os.Getenv("ENVIRONMENT")

	logger, err := logger.FromConfig(&logger.Config{Environment: environment})
	if err != nil {
		panic(err)
	}

	logger.Info("Starting Collection Application")

	storeType := storage.InMemory
	store, err := storage.StorageManager(storeType)
	if err != nil {
		logger.Panicf("Failed to initialise %s Store %w", storeType, err)
	}

	// Start gRPC Server inside a GO Routine
	go func() {
		err := server.StartGrpcServer(logger, store)
		if err != nil {
			logger.Panicf("Failed to start gRPC Server %w", err)
		}
	}()

	// Start the HTTP Server which wraps a gRPC Gateway
	err = server.StartHTTPServer(logger, store)
	if err != nil {
		logger.Panicf("Failed to start HTTP Gateway %w", err)
	}

	logger.Info("Exiting Collection Application")
}
