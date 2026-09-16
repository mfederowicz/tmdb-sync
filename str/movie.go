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
