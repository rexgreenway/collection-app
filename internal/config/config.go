package config

import (
	"os"

	"github.com/rexgreenway/collection-app/internal/storage"
	"go.uber.org/zap"
)

// Environment Variants
const (
	PRODUCTION  string = "production"
	DEVELOPMENT string = "development"
)

// Environment Variable Constants
const (
	ENVIRONMENT string = "ENVIRONMENT"
	STORE       string = "STORE"
	SERVER      string = "SERVER"
)

// Config defines the configuration for the application upon startup.
type Config struct {
	Environment string

	Store storage.StorageType
}

// FromEnv loads the application configuration from environment variables.
// It sets default values if environment variables are not set.
func FromEnv(log *zap.SugaredLogger) Config {
	log.Debugf("Loading application config...")

	environment, ok := os.LookupEnv(ENVIRONMENT)
	if !ok {
		environment = PRODUCTION
		os.Setenv(ENVIRONMENT, environment)

		log.Debugf("ENVIRONMENT not set, defaulting to %q", environment)
	}

	storeType, ok := os.LookupEnv(STORE)
	if !ok {
		storeType = string(storage.IN_MEMORY)
		os.Setenv(STORE, storeType)

		log.Debugf("STORE not set, defaulting to %q", storeType)
	}

	// SERVERS
	// Not implemented yet, currently both

	return Config{
		Environment: environment,
		Store:       storage.StorageType(storeType),
	}
}
