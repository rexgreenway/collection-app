package entities

import "time"

// Metadata holds store-managed, read-only bookkeeping fields.
type Metadata struct {
	CreatedAt time.Time `json:"created_at"`
}
