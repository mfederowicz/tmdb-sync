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

// ListCreatorV4 is the created_by object of a v4 list.
type ListCreatorV4 struct {
	GravatarHash string `json:"gravatar_hash"`
	ID           string `json:"id"`
	Name         string `json:"name"`
	Username     string `json:"username"`
}

// ListItemV4 is one entry of a v4 list. TMDB sends movie fields (title,
// release_date) or TV fields (name, first_air_date) depending on MediaType.
type ListItemV4 struct {
	Adult            bool     `json:"adult"`
	BackdropPath     string   `json:"backdrop_path"`
	FirstAirDate     string   `json:"first_air_date,omitempty"`
	GenreIDs         []int    `json:"genre_ids"`
	ID               int64    `json:"id"`
	MediaType        string   `json:"media_type"`
	Name             string   `json:"name,omitempty"`
	OriginCountry    []string `json:"origin_country,omitempty"`
	OriginalLanguage string   `json:"original_language"`
	OriginalName     string   `json:"original_name,omitempty"`
	OriginalTitle    string   `json:"original_title,omitempty"`
	Overview         string   `json:"overview"`
	Popularity       float64  `json:"popularity"`
	PosterPath       string   `json:"poster_path"`
	ReleaseDate      string   `json:"release_date,omitempty"`
	Title            string   `json:"title,omitempty"`
	Video            bool     `json:"video,omitempty"`
	VoteAverage      float64  `json:"vote_average"`
	VoteCount        int      `json:"vote_count"`
}

// ListV4 is the response shape for GET /4/list/{list_id}. Results holds the
// items of every fetched page.
type ListV4 struct {
	AverageRating float64        `json:"average_rating"`
	BackdropPath  string         `json:"backdrop_path"`
	Comments      map[string]any `json:"comments"`
	CreatedBy     ListCreatorV4  `json:"created_by"`
	Description   string         `json:"description"`
	ID            int64          `json:"id"`
	ISO31661      string         `json:"iso_3166_1"`
	ISO6391       string         `json:"iso_639_1"`
	ItemCount     int            `json:"item_count"`
	Name          string         `json:"name"`
	Page          int            `json:"page"`
	PosterPath    string         `json:"poster_path"`
	Public        int            `json:"public"`
	Results       []ListItemV4   `json:"results"`
	Revenue       any            `json:"revenue"`
	Runtime       int            `json:"runtime"`
	SortBy        any            `json:"sort_by"`
	TotalPages    int            `json:"total_pages"`
	TotalResults  int            `json:"total_results"`
}

// ListCreateRequestV4 is the body of POST /4/list.
type ListCreateRequestV4 struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	ISO6391     string `json:"iso_639_1"`
	ISO31661    string `json:"iso_3166_1"`
	Public      bool   `json:"public"`
}

// ListCreateResponseV4 is the response of POST /4/list.
type ListCreateResponseV4 struct {
	ID            int64  `json:"id"`
	StatusCode    int    `json:"status_code"`
	StatusMessage string `json:"status_message"`
	Success       bool   `json:"success"`
}

// ListUpdateRequestV4 is the body of PUT /4/list/{list_id}; only the fields
// that are set are sent.
type ListUpdateRequestV4 struct {
	BackdropPath string `json:"backdrop_path,omitempty"`
	Description  string `json:"description,omitempty"`
	Name         string `json:"name,omitempty"`
	Public       *bool  `json:"public,omitempty"`
	SortBy       string `json:"sort_by,omitempty"`
}

// ListStatusV4 is the response of the v4 list mutation endpoints that only
// report success (update, delete, items, clear).
type ListStatusV4 struct {
	StatusCode    int    `json:"status_code"`
	StatusMessage string `json:"status_message"`
	Success       bool   `json:"success"`
}

// ListMediaV4 is a media entry sent to the v4 items endpoints.
type ListMediaV4 struct {
	MediaType string `json:"media_type"`
	MediaID   int64  `json:"media_id"`
	Comment   string `json:"comment,omitempty"`
}

// ListItemsRequestV4 is the body of POST and PUT /4/list/{list_id}/items;
// comment is only meaningful for PUT.
type ListItemsRequestV4 struct {
	Items []ListMediaV4 `json:"items"`
}

// ListItemResultV4 is the per-item outcome reported by the v4 items endpoints.
type ListItemResultV4 struct {
	MediaID   int64  `json:"media_id"`
	MediaType string `json:"media_type"`
	Success   bool   `json:"success"`
	Error     any    `json:"error,omitempty"`
}

// ListItemsResponseV4 is the response of the v4 items endpoints.
type ListItemsResponseV4 struct {
	Results       []ListItemResultV4 `json:"results"`
	StatusCode    int                `json:"status_code"`
	StatusMessage string             `json:"status_message"`
	Success       bool               `json:"success"`
}
