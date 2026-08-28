package entities

import (
	"fmt"
	"time"
)

// Metadata holds store-managed, read-only bookkeeping fields.
type Metadata struct {
	CreatedAt time.Time `json:"created_at"`
}

// StampCreatedAt sets the Metadata field CreatedAt with the given time if unset.
func (m *Metadata) StampCreatedAt(time time.Time) error {
	if !m.CreatedAt.IsZero() {
		return fmt.Errorf("Metadata field CreateAt is already set.")
	}
	m.CreatedAt = time
	return nil
}
