package str

// Movie is the response of GET /movie/{movie_id}.
type Movie struct {
	ID               int64   `json:"id"`
	Title            string  `json:"title"`
	OriginalTitle    string  `json:"original_title"`
	Overview         string  `json:"overview"`
	ReleaseDate      string  `json:"release_date"`
	Runtime          int     `json:"runtime"`
	VoteAverage      float64 `json:"vote_average"`
	VoteCount        int     `json:"vote_count"`
	Popularity       float64 `json:"popularity"`
	Adult            bool    `json:"adult"`
	Video            bool    `json:"video"`
	OriginalLanguage string  `json:"original_language"`
}

// RatedMovie is a movie as returned by an account's rated-movies endpoint,
// which adds the account's own rating to the usual movie fields.
type RatedMovie struct {
	Movie
	Rating float64 `json:"rating"`
}

// RatedMovies is a paginated list of an account's rated movies.
type RatedMovies struct {
	Page         int          `json:"page"`
	Results      []RatedMovie `json:"results"`
	TotalPages   int          `json:"total_pages"`
	TotalResults int          `json:"total_results"`
}

// MovieAccountStates is the response shape for GET /movie/{movie_id}/account_states.
// Rated is `any` because TMDB returns either `false` (not rated) or an object
// like {"value": 8} (rated) for the same field.
type MovieAccountStates struct {
	ID        int64 `json:"id"`
	Favorite  bool  `json:"favorite"`
	Rated     any   `json:"rated"`
	Watchlist bool  `json:"watchlist"`
}

// MovieAlternativeTitle is one alternative title, as returned inline on the
// movie alternative-titles endpoint.
type MovieAlternativeTitle struct {
	Iso31661 string `json:"iso_3166_1"`
	Title    string `json:"title"`
	Type     string `json:"type"`
}

// MovieAlternativeTitles is the response shape for
// GET /movie/{movie_id}/alternative_titles.
type MovieAlternativeTitles struct {
	ID     int64                   `json:"id"`
	Titles []MovieAlternativeTitle `json:"titles"`
}
