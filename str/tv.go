package str

// TV is the response of GET /tv/{series_id}.
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
	InProduction     bool    `json:"in_production"`
	NumberOfEpisodes int     `json:"number_of_episodes"`
	NumberOfSeasons  int     `json:"number_of_seasons"`
	OriginalLanguage string  `json:"original_language"`
}

// TVAccountStates is the response shape for GET /tv/{series_id}/account_states.
// Rated is `any` because TMDB returns either `false` (not rated) or an object
// like {"value": 8} (rated) for the same field.
type TVAccountStates struct {
	ID        int64 `json:"id"`
	Favorite  bool  `json:"favorite"`
	Rated     any   `json:"rated"`
	Watchlist bool  `json:"watchlist"`
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

// RatedTVEpisode is a TV episode as returned by an account's
// rated-tv-episodes endpoint, which adds the account's own rating to the
// usual episode fields.
type RatedTVEpisode struct {
	AirDate        string  `json:"air_date"`
	EpisodeNumber  int     `json:"episode_number"`
	ID             int64   `json:"id"`
	Name           string  `json:"name"`
	Overview       string  `json:"overview"`
	ProductionCode string  `json:"production_code"`
	Runtime        int     `json:"runtime"`
	SeasonNumber   int     `json:"season_number"`
	ShowID         int64   `json:"show_id"`
	StillPath      string  `json:"still_path"`
	VoteAverage    float64 `json:"vote_average"`
	VoteCount      int     `json:"vote_count"`
	Rating         float64 `json:"rating"`
}

// RatedTVEpisodes is a paginated list of an account's rated TV episodes.
type RatedTVEpisodes struct {
	Page         int              `json:"page"`
	Results      []RatedTVEpisode `json:"results"`
	TotalPages   int              `json:"total_pages"`
	TotalResults int              `json:"total_results"`
}
