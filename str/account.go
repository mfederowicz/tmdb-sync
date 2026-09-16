package str

// Account is the response of GET /account/{account_id}.
type Account struct {
	Avatar       *AccountAvatar `json:"avatar"`
	ID           int64          `json:"id"`
	ISO6391      string         `json:"iso_639_1"`
	ISO31661     string         `json:"iso_3166_1"`
	Name         string         `json:"name"`
	IncludeAdult bool           `json:"include_adult"`
	Username     string         `json:"username"`
}

// AccountAvatar holds an account's avatar images.
type AccountAvatar struct {
	Gravatar *AccountGravatar `json:"gravatar"`
	Tmdb     *AccountTmdb     `json:"tmdb"`
}

// AccountGravatar is a Gravatar-backed avatar.
type AccountGravatar struct {
	Hash string `json:"hash"`
}

// AccountTmdb is a TMDB-hosted avatar.
type AccountTmdb struct {
	AvatarPath *string `json:"avatar_path"`
}

// AccountWatchlistRequest is the body of POST /account/{account_id}/watchlist.
type AccountWatchlistRequest struct {
	MediaType string `json:"media_type"`
	MediaID   int64  `json:"media_id"`
	Watchlist bool   `json:"watchlist"`
}

// AccountFavoriteRequest is the body of POST /account/{account_id}/favorite.
type AccountFavoriteRequest struct {
	MediaType string `json:"media_type"`
	MediaID   int64  `json:"media_id"`
	Favorite  bool   `json:"favorite"`
}

// AccountList is one of an account's custom lists, as returned by
// GET /account/{account_id}/lists.
type AccountList struct {
	Description   string `json:"description"`
	FavoriteCount int    `json:"favorite_count"`
	ID            int64  `json:"id"`
	ItemCount     int    `json:"item_count"`
	ISO6391       string `json:"iso_639_1"`
	ListType      string `json:"list_type"`
	Name          string `json:"name"`
	PosterPath    string `json:"poster_path"`
}

// AccountLists is a paginated list of an account's custom lists.
type AccountLists struct {
	Page         int           `json:"page"`
	Results      []AccountList `json:"results"`
	TotalPages   int           `json:"total_pages"`
	TotalResults int           `json:"total_results"`
}
