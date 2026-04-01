package collection

import (
	"github.com/rexgreenway/collection-app/internal/entities"
	pb "github.com/rexgreenway/collection-app/internal/gen/collection"
)

func collectionToProto(c entities.Collection) *pb.Collection {
	return &pb.Collection{
		Id:   c.ID,
		Name: c.Name,
	}
}

func protoToCollection(c *pb.Collection) entities.Collection {
	return entities.Collection{
		ID:   c.GetId(),
		Name: c.GetName(),
	}
}
