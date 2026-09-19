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

// AccountListV4 is one of an account's lists, as returned by
// GET /4/account/{account_object_id}/lists. TMDB sends the flags (adult,
// featured, public) and sort_by as integers and revenue as a string.
type AccountListV4 struct {
	AccountObjectID string  `json:"account_object_id"`
	Adult           int     `json:"adult"`
	AverageRating   float64 `json:"average_rating"`
	BackdropPath    string  `json:"backdrop_path"`
	CreatedAt       string  `json:"created_at"`
	Description     string  `json:"description"`
	Featured        int     `json:"featured"`
	ID              int64   `json:"id"`
	ISO31661        string  `json:"iso_3166_1"`
	ISO6391         string  `json:"iso_639_1"`
	Name            string  `json:"name"`
	NumberOfItems   int     `json:"number_of_items"`
	PosterPath      string  `json:"poster_path"`
	Public          int     `json:"public"`
	Revenue         string  `json:"revenue"`
	Runtime         int     `json:"runtime"`
	SortBy          int     `json:"sort_by"`
	UpdatedAt       string  `json:"updated_at"`
}

// AccountRatingV4 is the user's rating on an item in the v4 rated endpoints.
type AccountRatingV4 struct {
	CreatedAt string  `json:"created_at"`
	Value     float64 `json:"value"`
}

// RatedMovieV4 is a movie in GET /4/account/{account_object_id}/movie/rated.
type RatedMovieV4 struct {
	Movie
	AccountRating *AccountRatingV4 `json:"account_rating"`
}

// RatedTVV4 is a TV show in GET /4/account/{account_object_id}/tv/rated.
type RatedTVV4 struct {
	TV
	AccountRating *AccountRatingV4 `json:"account_rating"`
}
