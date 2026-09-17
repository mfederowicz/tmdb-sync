package str

// Person is a minimal person shape as returned by the find endpoint's
// person_results.
type Person struct {
	ID           int64   `json:"id"`
	Name         string  `json:"name"`
	OriginalName string  `json:"original_name"`
	Popularity   float64 `json:"popularity"`
	Adult        bool    `json:"adult"`
	Gender       int     `json:"gender"`
	KnownForDept string  `json:"known_for_department"`
	ProfilePath  string  `json:"profile_path"`
}

// TVEpisodeResult is a TV episode as returned by the find endpoint's
// tv_episode_results.
type TVEpisodeResult struct {
	ID            int64   `json:"id"`
	Name          string  `json:"name"`
	Overview      string  `json:"overview"`
	AirDate       string  `json:"air_date"`
	EpisodeNumber int     `json:"episode_number"`
	SeasonNumber  int     `json:"season_number"`
	ShowID        int64   `json:"show_id"`
	StillPath     string  `json:"still_path"`
	VoteAverage   float64 `json:"vote_average"`
	VoteCount     int     `json:"vote_count"`
}

// TVSeasonResult is a TV season as returned by the find endpoint's
// tv_season_results.
type TVSeasonResult struct {
	ID           int64  `json:"id"`
	Name         string `json:"name"`
	Overview     string `json:"overview"`
	AirDate      string `json:"air_date"`
	SeasonNumber int    `json:"season_number"`
	ShowID       int64  `json:"show_id"`
	PosterPath   string `json:"poster_path"`
}

// FindResults is the response of GET /find/{external_id}.
type FindResults struct {
	MovieResults     []Movie           `json:"movie_results"`
	PersonResults    []Person          `json:"person_results"`
	TVResults        []TV              `json:"tv_results"`
	TVEpisodeResults []TVEpisodeResult `json:"tv_episode_results"`
	TVSeasonResults  []TVSeasonResult  `json:"tv_season_results"`
}
