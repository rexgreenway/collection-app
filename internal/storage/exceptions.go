package storage

import (
	"errors"
)

// EXCEPTIONS
var (
	ErrCollectionNotFound      = errors.New("collection not found")
	ErrCollectionAlreadyExists = errors.New("collection already exists")

	ErrItemNotFound      = errors.New("item not found")
	ErrItemAlreadyExists = errors.New("item already exists")
)
