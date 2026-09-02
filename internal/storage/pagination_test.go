package storage

import (
	"testing"

	"github.com/rexgreenway/collection-app/internal/entities"
	"github.com/stretchr/testify/assert"
)

func TestPaginationUtils(t *testing.T) {
	// APPLYING PAGINATION

	t.Run("default pagination when missing", func(t *testing.T) {
		pagination := applyPaginationDefaults(nil)

		assert.Equal(t, pagination, &entities.Pagination{
			Page:     DefaultPage,
			PageSize: DefaultPageSize,
		})
	})

	t.Run("default pagination when values missing", func(t *testing.T) {
		pagination := applyPaginationDefaults(&entities.Pagination{})

		assert.Equal(t, pagination, &entities.Pagination{
			Page:     DefaultPage,
			PageSize: DefaultPageSize,
		})
	})

	t.Run("return pagination unchanged when provided", func(t *testing.T) {
		pagination := applyPaginationDefaults(
			&entities.Pagination{
				Page:     4,
				PageSize: 23,
			},
		)

		assert.Equal(t, pagination, &entities.Pagination{
			Page:     4,
			PageSize: 23,
		})
	})

	t.Run("default pagination if invalid values", func(t *testing.T) {
		pagination := applyPaginationDefaults(
			&entities.Pagination{
				Page:     -4,
				PageSize: -23,
			},
		)

		assert.Equal(t, pagination, &entities.Pagination{
			Page:     DefaultPage,
			PageSize: DefaultPageSize,
		})
	})

	// RESOLVING PAGINATION

	t.Run("resolve pagination with no entities", func(t *testing.T) {
		bounds := resolvePaginationBounds(nil, 0)

		assert.Equal(t, bounds, entities.PaginationBounds{
			Start: 0,
			End:   0,
		})
	})

	t.Run("resolve pagination with fewer entities than page size", func(t *testing.T) {
		total := DefaultPageSize - 2
		bounds := resolvePaginationBounds(nil, total)

		assert.Equal(t, bounds, entities.PaginationBounds{
			Start: 0,
			End:   total,
		})
	})

	t.Run("resolve pagination with greater entities than page size", func(t *testing.T) {
		total := DefaultPageSize + 2
		bounds := resolvePaginationBounds(nil, total)

		assert.Equal(t, bounds, entities.PaginationBounds{
			Start: 0,
			End:   DefaultPageSize,
		})
	})

	t.Run("resolve pagination with greater entities than page size", func(t *testing.T) {
		total := DefaultPageSize + 2
		bounds := resolvePaginationBounds(
			&entities.Pagination{
				Page: 2,
			},
			total,
		)

		assert.Equal(t, bounds, entities.PaginationBounds{
			Start: DefaultPageSize,
			End:   total,
		})
	})

	t.Run("check pagination updates to maximum page", func(t *testing.T) {
		// Page 5 with 10 items per page => min. total entities = 51.
		// However total = 25 < 51 => max. 3 pages of size 10.
		pagination := &entities.Pagination{
			Page:     5,
			PageSize: 10,
		}
		bounds := resolvePaginationBounds(pagination, 25)

		assert.Equal(t, bounds, entities.PaginationBounds{
			Start: 20,
			End:   25,
		})
		assert.Equal(t, int32(3), pagination.Page)
	})
}
