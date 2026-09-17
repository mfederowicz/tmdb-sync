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
