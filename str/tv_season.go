package str

// TVSeasonEpisode is one episode entry nested inside a TV season's details.
type TVSeasonEpisode struct {
	AirDate        string         `json:"air_date"`
	EpisodeNumber  int            `json:"episode_number"`
	ID             int64          `json:"id"`
	Name           string         `json:"name"`
	Overview       string         `json:"overview"`
	ProductionCode string         `json:"production_code"`
	Runtime        int            `json:"runtime"`
	SeasonNumber   int            `json:"season_number"`
	ShowID         int64          `json:"show_id"`
	StillPath      string         `json:"still_path"`
	VoteAverage    float64        `json:"vote_average"`
	VoteCount      int64          `json:"vote_count"`
	Crew           []TVCrewMember `json:"crew"`
	GuestStars     []TVCastMember `json:"guest_stars"`
}

// TVSeasonAccountState is a single episode's account state, as returned by
// the tv season account_states endpoint.
type TVSeasonAccountState struct {
	ID            int64 `json:"id"`
	EpisodeNumber int   `json:"episode_number"`
	Rated         any   `json:"rated"`
}

// TVSeasonAccountStates is the response shape for
// GET /tv/{series_id}/season/{season_number}/account_states.
type TVSeasonAccountStates struct {
	ID      int64                  `json:"id"`
	Results []TVSeasonAccountState `json:"results"`
}

// TVSeasonExternalIDs is the response shape for
// GET /tv/{series_id}/season/{season_number}/external_ids.
type TVSeasonExternalIDs struct {
	ID          int64  `json:"id"`
	FreebaseMID string `json:"freebase_mid"`
	FreebaseID  string `json:"freebase_id"`
	TvdbID      int64  `json:"tvdb_id"`
	TvrageID    int64  `json:"tvrage_id"`
	WikidataID  string `json:"wikidata_id"`
}

// TVSeason is the response shape for GET /tv/{series_id}/season/{season_number}.
type TVSeason struct {
	ID           int64             `json:"id"`
	AirDate      string            `json:"air_date"`
	Episodes     []TVSeasonEpisode `json:"episodes"`
	Name         string            `json:"name"`
	Overview     string            `json:"overview"`
	PosterPath   string            `json:"poster_path"`
	SeasonNumber int               `json:"season_number"`
	VoteAverage  float64           `json:"vote_average"`
}
