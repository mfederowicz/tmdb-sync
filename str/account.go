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
