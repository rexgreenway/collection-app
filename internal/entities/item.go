package entities

type Item struct {
	Id   string `json:"id"`
	Name string `json:"name"`

	CollectionId string `json:"collection_id"`
}
