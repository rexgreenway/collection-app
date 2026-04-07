package entities

type Item struct {
	Id   string `json:"id"`
	Name string `json:"name"`

	CollectionId string `json:"collection_id"`
}

// IDEAS:
// - ItemType to indicate a specific schema for an item (film, comic, book, etc.)
// - Info/Extra field to hold metadata
