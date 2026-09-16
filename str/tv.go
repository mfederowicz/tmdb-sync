package str

// TV is a TV show as returned by TMDB TV endpoints (e.g. account favorite/rated TV).
type TV struct {
	ID               int64   `json:"id"`
	Name             string  `json:"name"`
	OriginalName     string  `json:"original_name"`
	Overview         string  `json:"overview"`
	FirstAirDate     string  `json:"first_air_date"`
	VoteAverage      float64 `json:"vote_average"`
	VoteCount        int     `json:"vote_count"`
	Popularity       float64 `json:"popularity"`
	Adult            bool    `json:"adult"`
	OriginalLanguage string  `json:"original_language"`
}

// TVShows is a paginated list of TV shows.
type TVShows struct {
	Page         int  `json:"page"`
	Results      []TV `json:"results"`
	TotalPages   int  `json:"total_pages"`
	TotalResults int  `json:"total_results"`
}

// RatedTV is a TV show as returned by an account's rated-tv endpoint, which
// adds the account's own rating to the usual TV fields.
type RatedTV struct {
	TV
	Rating float64 `json:"rating"`
}

// RatedTVShows is a paginated list of an account's rated TV shows.
type RatedTVShows struct {
	Page         int       `json:"page"`
	Results      []RatedTV `json:"results"`
	TotalPages   int       `json:"total_pages"`
	TotalResults int       `json:"total_results"`
}
