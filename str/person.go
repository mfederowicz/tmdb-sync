package str

// PersonDetails is the response shape for the person details endpoint.
type PersonDetails struct {
	Adult              bool     `json:"adult"`
	AlsoKnownAs        []string `json:"also_known_as"`
	Biography          string   `json:"biography"`
	Birthday           string   `json:"birthday"`
	Deathday           string   `json:"deathday"`
	Gender             int      `json:"gender"`
	Homepage           string   `json:"homepage"`
	ID                 int64    `json:"id"`
	ImdbID             string   `json:"imdb_id"`
	KnownForDepartment string   `json:"known_for_department"`
	Name               string   `json:"name"`
	PlaceOfBirth       string   `json:"place_of_birth"`
	Popularity         float64  `json:"popularity"`
	ProfilePath        string   `json:"profile_path"`
}

// PersonCombinedCredit is a single cast/crew entry in a person's combined
// credits (movie or TV, distinguished by MediaType).
type PersonCombinedCredit struct {
	ID            int64   `json:"id"`
	MediaType     string  `json:"media_type"`
	Adult         bool    `json:"adult"`
	Title         string  `json:"title"`
	OriginalTitle string  `json:"original_title"`
	Name          string  `json:"name"`
	OriginalName  string  `json:"original_name"`
	Overview      string  `json:"overview"`
	PosterPath    string  `json:"poster_path"`
	ReleaseDate   string  `json:"release_date"`
	FirstAirDate  string  `json:"first_air_date"`
	Popularity    float64 `json:"popularity"`
	VoteAverage   float64 `json:"vote_average"`
	VoteCount     int64   `json:"vote_count"`
	Character     string  `json:"character"`
	CreditID      string  `json:"credit_id"`
	Department    string  `json:"department"`
	Job           string  `json:"job"`
	Order         int     `json:"order"`
	EpisodeCount  int     `json:"episode_count"`
}

// PersonCombinedCredits is the response shape for
// GET /person/{person_id}/combined_credits.
type PersonCombinedCredits struct {
	ID   int64                  `json:"id"`
	Cast []PersonCombinedCredit `json:"cast"`
	Crew []PersonCombinedCredit `json:"crew"`
}
