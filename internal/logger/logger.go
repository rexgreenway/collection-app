package logger

import (
	"fmt"

	"go.uber.org/zap"
)

var logger *zap.SugaredLogger

type Config struct {
	Environment string
}

// FromConfig returns sugared logger provided with a configuration.
func FromConfig(config *Config) (*zap.SugaredLogger, error) {
	var l *zap.Logger

	switch config.Environment {
	case "dev":
		l = zap.Must(zap.NewDevelopment())
	case "prod":
		l = zap.Must(zap.NewProduction())
	default:
		return nil, fmt.Errorf("no such environment %q", config.Environment)
	}

	logger = l.Sugar()

	return logger, nil
}
