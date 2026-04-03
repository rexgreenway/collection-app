package entities

type Item struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// IDEAS:
// - ItemType to indicate a specific schema for an item (film, comic, book, etc.)
// - Info/Extra field to hold metadata
