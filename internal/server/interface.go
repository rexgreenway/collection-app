package server

import (
	"context"
)

type ServerType string

// Server is a runnable network server.
type Server interface {
	// Name identifies the server for logging.
	Name() ServerType

	// Start runs the server, blocking until ctx is cancelled or it fails.
	Start(ctx context.Context) error
}
