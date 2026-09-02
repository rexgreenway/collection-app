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

	logger.Info("Collection Application Starting...")

	// Set if Gin is Prod or not?
	// THIS ONLY NEEDS TO BE SET IF GIN IS ACTUALLY EVEN BEING RUN
	if cfg.Environment == config.PRODUCTION {
		gin.SetMode(gin.ReleaseMode)
	}

	store, err := storage.StorageFactory(cfg.Store, logger)
	if err != nil {
		logger.Fatalf("Failed to initialise %q type Store: %v\n", cfg.Store, err)
	}

	collectionService := collection.NewService(logger, store)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	// servers is the configured set of servers you want to run.
	// THIS SHOULD BE DRIVEN BY CONFIG!
	servers := []server.Server{
		server.NewGrpcServer(logger, "localhost:50100", collectionService),
		server.NewGinServer(logger, ":8089", collectionService),
	}

	// Start servers in err group
	g, ctx := errgroup.WithContext(ctx)
	for _, s := range servers {
		s := s // capture (unnecessary on Go 1.22+)
		g.Go(func() error {
			// Checking for errors here allows us to wrap the error with more context!
			if err := s.Start(ctx); err != nil {
				return fmt.Errorf("%s server: %w", s.Name(), err)
			}
			return nil
		})
	}

	if err := g.Wait(); err != nil {
		logger.Errorf("Application exiting due to error: %v", err)
	}

	logger.Info("Collection Application Exiting...")
}
