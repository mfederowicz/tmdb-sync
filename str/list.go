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

// ListCreateRequest is the body of POST /list.
type ListCreateRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Language    string `json:"language,omitempty"`
}

// ListCreateResponse is the response of POST /list.
type ListCreateResponse struct {
	Success       bool   `json:"success"`
	StatusCode    int    `json:"status_code"`
	StatusMessage string `json:"status_message"`
	ListID        int64  `json:"list_id"`
}

// ListItemRequest is the body of POST /list/{list_id}/add_item and
// POST /list/{list_id}/remove_item.
type ListItemRequest struct {
	MediaID int64 `json:"media_id"`
}
