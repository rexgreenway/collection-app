package logger

import (
	"fmt"

	"github.com/rexgreenway/collection-app/internal/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.SugaredLogger

// FromConfig returns sugared logger provided with a configuration.
func FromConfig(cfg *config.Config) (*zap.SugaredLogger, error) {
	var l *zap.Logger

	switch cfg.Environment {

	case config.DEVELOPMENT:
		config := zap.NewDevelopmentConfig()
		config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		l = zap.Must(config.Build())

	case config.PRODUCTION:
		l = zap.Must(zap.NewProduction())

	default:
		return nil, fmt.Errorf("no such environment %q", cfg.Environment)
	}

	logger = l.Sugar()

	return logger, nil
}
