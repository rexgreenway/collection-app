package storage

import (
	"math"

	"github.com/rexgreenway/collection-app/internal/entities"
)

const (
	// DefaultPage is the default page to return if none is specified.
	DefaultPage = int32(1)
	// DefaultPageSize is the default page size to return if none is specified.
	DefaultPageSize = int32(10)
)

// validateTransformPagination ???
func validateTransformPagination(pagination *entities.Pagination, total int32) {
	if total == 0 {
		pagination.Page = int32(1)
	} else {
		// Transform to defaults in the storage package as related to interaction with store specifically
		// In the future could move this to be something each implementation enacts itself, with limits/maxes etc.
		// As various Storage tools might have different requirements (i.e. DataStore 30???)
		if pagination.Page == 0 {
			pagination.Page = DefaultPage
		}
		if pagination.PageSize == 0 {
			pagination.PageSize = DefaultPageSize
		}

		// Find max page
		max_page := math.Ceil(float64(total) / float64(pagination.PageSize))
		pagination.Page = int32(math.Min(float64(pagination.Page), max_page))

		pagination.Start = pagination.PageSize * (pagination.Page - 1)
		pagination.End = int32(math.Min(float64(pagination.PageSize*pagination.Page), float64(total)))
	}
}
