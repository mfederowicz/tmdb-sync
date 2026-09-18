package str

// TVEpisodeGroupDetailsEpisode is one episode inside a group entry, as
// returned by the tv episode group details endpoint.
type TVEpisodeGroupDetailsEpisode struct {
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
	Order          int     `json:"order"`
}

// TVEpisodeGroupDetailsGroup is one group entry (a "season" in the group's
// own custom ordering), as returned by the tv episode group details endpoint.
type TVEpisodeGroupDetailsGroup struct {
	ID       string                         `json:"id"`
	Name     string                         `json:"name"`
	Order    int                            `json:"order"`
	Locked   bool                           `json:"locked"`
	Episodes []TVEpisodeGroupDetailsEpisode `json:"episodes"`
}

// TVEpisodeGroupDetails is the response shape for GET /tv/episode_group/{id}.
type TVEpisodeGroupDetails struct {
	Description  string                       `json:"description"`
	EpisodeCount int                          `json:"episode_count"`
	GroupCount   int                          `json:"group_count"`
	Groups       []TVEpisodeGroupDetailsGroup `json:"groups"`
	ID           string                       `json:"id"`
	Name         string                       `json:"name"`
	Network      TVEpisodeGroupNetwork        `json:"network"`
	Type         int                          `json:"type"`
}
