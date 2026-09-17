package str

// List is the response shape for GET /list/{list_id}.
type List struct {
	ID            int64   `json:"id"`
	Name          string  `json:"name"`
	Description   string  `json:"description"`
	CreatedBy     string  `json:"created_by"`
	FavoriteCount int     `json:"favorite_count"`
	ItemCount     int     `json:"item_count"`
	ISO6391       string  `json:"iso_639_1"`
	PosterPath    string  `json:"poster_path"`
	Items         []Movie `json:"items"`
}

// ListItemStatus is the response shape for GET /list/{list_id}/item_status.
type ListItemStatus struct {
	ID          int64 `json:"id"`
	ItemPresent bool  `json:"item_present"`
}
