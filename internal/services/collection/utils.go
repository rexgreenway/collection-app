package collection

import (
	"github.com/rexgreenway/collection-app/internal/entities"
	pb "github.com/rexgreenway/collection-app/internal/gen/v1/collection"
)

// ----- TRANSFORMERS ----

// protoToPagination ???
func protoToPagination(pbPagination *pb.PaginationParams) entities.Pagination {
	var page, pageSize int32
	if pbPagination != nil {
		page = pbPagination.GetPage()
		pageSize = pbPagination.GetPageSize()
	}

	return entities.Pagination{
		Page:     page,
		PageSize: pageSize,
	}
}

// paginationToProto ???
func paginationToProto(p entities.Pagination) *pb.PaginationParams {
	foo := &pb.PaginationParams{
		Page:     p.Page,
		PageSize: p.PageSize,
	}

	return foo
}

// collectionToProto ???
func collectionToProto(c entities.Collection) *pb.Collection {
	return &pb.Collection{
		Id:   c.ID,
		Name: c.Name,
	}
}

// protoToCollection ???
func protoToCollection(c *pb.Collection) entities.Collection {
	return entities.Collection{
		ID:   c.GetId(),
		Name: c.GetName(),
	}
}
