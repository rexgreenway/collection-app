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
	return &pb.PaginationParams{
		Page:     p.Page,
		PageSize: p.PageSize,
	}
}

// ------- Collection Transformers -------

// collectionToProto ???
func collectionToProto(collection entities.Collection) *pb.Collection {
	return &pb.Collection{
		Id:   collection.Id,
		Name: collection.Name,
	}
}

// protoToCollection ???
func protoToCollection(collection *pb.Collection) entities.Collection {
	return entities.Collection{
		Id:   collection.GetId(),
		Name: collection.GetName(),
	}
}

// ------- Item Transformers -------

// itemToProto
func itemToProto(item entities.Item) *pb.Item {
	return &pb.Item{
		Id:           item.Id,
		Name:         item.Name,
		CollectionId: item.CollectionId,
	}
}

// protoToItem ???
func protoToItem(item *pb.Item) entities.Item {
	return entities.Item{
		Id:           item.GetId(),
		Name:         item.GetName(),
		CollectionId: item.GetCollectionId(),
	}
}
