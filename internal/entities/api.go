package entities

// Pagination ???
type Pagination struct {
	Page     int32 `json:"page"`
	PageSize int32 `json:"page_size"`
}

// PaginationBounds ???
type PaginationBounds struct {
	Start int32 `json:"start"`
	End   int32 `json:"end"`
}
