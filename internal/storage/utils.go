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

// applyPaginationDefaults populates default Page and PageSize if pagination or
// pagination values are nil or are invalid (i.e. negative values).
func applyPaginationDefaults(pagination *entities.Pagination) *entities.Pagination {
	if pagination == nil {
		pagination = &entities.Pagination{}
	}

	if pagination.Page <= 0 {
		pagination.Page = DefaultPage
	}
	if pagination.PageSize <= 0 {
		pagination.PageSize = DefaultPageSize
	}

	return pagination
}

// resolvePaginationBounds ???
func resolvePaginationBounds(pagination *entities.Pagination, total int32) entities.PaginationBounds {
	pagination = applyPaginationDefaults(pagination)

	var maxPage float64
	if total == 0 {
		maxPage = 1
	} else {
		maxPage = math.Ceil(float64(total) / float64(pagination.PageSize))
	}

	pagination.Page = int32(math.Min(float64(pagination.Page), maxPage))

	start := pagination.PageSize * (pagination.Page - 1)
	end := int32(math.Min(float64(pagination.PageSize*pagination.Page), float64(total)))

	return entities.PaginationBounds{
		Start: start,
		End:   end,
	}
}
