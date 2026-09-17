package str

// SearchMultiResult is a single entry within a search/multi response. TMDB
// returns movie, tv, and person shapes interleaved, distinguished by
// MediaType; fields not relevant to a given media type are left zero.
type SearchMultiResult struct {
	ID               int64    `json:"id"`
	MediaType        string   `json:"media_type"`
	Adult            bool     `json:"adult"`
	Popularity       float64  `json:"popularity"`
	BackdropPath     string   `json:"backdrop_path"`
	GenreIDs         []int64  `json:"genre_ids"`
	OriginalLanguage string   `json:"original_language"`
	PosterPath       string   `json:"poster_path"`
	Overview         string   `json:"overview"`
	VoteAverage      float64  `json:"vote_average"`
	VoteCount        int      `json:"vote_count"`
	Title            string   `json:"title"`
	OriginalTitle    string   `json:"original_title"`
	ReleaseDate      string   `json:"release_date"`
	Video            bool     `json:"video"`
	Name             string   `json:"name"`
	OriginalName     string   `json:"original_name"`
	FirstAirDate     string   `json:"first_air_date"`
	OriginCountry    []string `json:"origin_country"`
	ProfilePath      string   `json:"profile_path"`
	KnownForDept     string   `json:"known_for_department"`
	Gender           int      `json:"gender"`
}

// SearchMulti is a paginated list of mixed movie/tv/person results, as
// returned by GET /search/multi.
type SearchMulti struct {
	Page         int                 `json:"page"`
	Results      []SearchMultiResult `json:"results"`
	TotalPages   int                 `json:"total_pages"`
	TotalResults int                 `json:"total_results"`
}
