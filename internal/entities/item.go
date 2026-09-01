package entities

// Item Defaults / Fallbacks
const (
	DEFAULT_ITEM_NAME = "New Item"

	NO_COLLECTION_ID = "no-collection-id"
)

// Item ???
type Item struct {
	Id   string `json:"id"`
	Name string `json:"name"`

	CollectionId string `json:"collection_id"`

	Metadata `json:"metadata"`
}

// ItemUpdate holds the mutable fields of an Item.
// A nil field is left unchanged.
type ItemUpdate struct {
	Name *string `json:"name,omitempty"`

	CollectionId *string `json:"collection_id,omitempty"`
}

// Update applies the non-nil fields of u to the item.
func (c *Item) Update(u ItemUpdate) {
	if u.Name != nil {
		c.Name = *u.Name
	}
	if u.CollectionId != nil {
		c.CollectionId = *u.CollectionId
	}
}
