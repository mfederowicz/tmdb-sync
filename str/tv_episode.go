package str

// TVEpisode is the response shape for
// GET /tv/{series_id}/season/{season_number}/episode/{episode_number}.
type TVEpisode struct {
	AirDate        string         `json:"air_date"`
	Crew           []TVCrewMember `json:"crew"`
	EpisodeNumber  int            `json:"episode_number"`
	GuestStars     []TVCastMember `json:"guest_stars"`
	ID             int64          `json:"id"`
	Name           string         `json:"name"`
	Overview       string         `json:"overview"`
	ProductionCode string         `json:"production_code"`
	Runtime        int            `json:"runtime"`
	SeasonNumber   int            `json:"season_number"`
	StillPath      string         `json:"still_path"`
	VoteAverage    float64        `json:"vote_average"`
	VoteCount      int64          `json:"vote_count"`
}

// TVEpisodeAccountStates is the response shape for
// GET .../episode/{episode_number}/account_states.
type TVEpisodeAccountStates struct {
	ID    int64 `json:"id"`
	Rated any   `json:"rated"`
}
