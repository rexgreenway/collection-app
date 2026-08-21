package entities

// Collection ???
type Collection struct {
	Id   string `json:"id"`
	Name string `json:"name"`

	Metadata Metadata `json:"metadata"`
}

// CollectionUpdate holds the mutable fields of a Collection.
// A nil field is left unchanged.
type CollectionUpdate struct {
	Name *string `json:"name,omitempty"`
}

// Update applies the non-nil fields of u to the collection.
func (c *Collection) Update(u CollectionUpdate) {
	if u.Name != nil {
		c.Name = *u.Name
	}
}
