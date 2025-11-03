package main

import (
	"os"

	"github.com/rexgreenway/collection-app/internal/logger"
	"github.com/rexgreenway/collection-app/internal/server"
)

func main() {
	// Establish Application Environment
	environment := os.Getenv("ENVIRONMENT")

	// Establish Logging
	// Use logger through dependency injection!!!
	logger := logger.FromConfig(&logger.Config{Environment: environment})

	logger.Info("starting collection application")

	// - Start gRPC Server
	go server.StartServer(logger)

	// Start HTTP Gateway
	server.StartHTTPGateway(logger)

	logger.Info("exiting collection application")
}
