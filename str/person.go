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

// PersonExternalIDs is the response shape for
// GET /person/{person_id}/external_ids.
type PersonExternalIDs struct {
	ID          int64  `json:"id"`
	FreebaseMid string `json:"freebase_mid"`
	FreebaseID  string `json:"freebase_id"`
	ImdbID      string `json:"imdb_id"`
	TvrageID    int64  `json:"tvrage_id"`
	WikidataID  string `json:"wikidata_id"`
	FacebookID  string `json:"facebook_id"`
	InstagramID string `json:"instagram_id"`
	TiktokID    string `json:"tiktok_id"`
	TwitterID   string `json:"twitter_id"`
	YoutubeID   string `json:"youtube_id"`
}

// PersonImages is the response shape for GET /person/{person_id}/images.
type PersonImages struct {
	ID       int64   `json:"id"`
	Profiles []Image `json:"profiles"`
}

// PersonMovieCredit is a single cast/crew entry in a person's movie credits.
type PersonMovieCredit struct {
	ID            int64   `json:"id"`
	Adult         bool    `json:"adult"`
	Title         string  `json:"title"`
	OriginalTitle string  `json:"original_title"`
	Overview      string  `json:"overview"`
	PosterPath    string  `json:"poster_path"`
	ReleaseDate   string  `json:"release_date"`
	Popularity    float64 `json:"popularity"`
	VoteAverage   float64 `json:"vote_average"`
	VoteCount     int64   `json:"vote_count"`
	Character     string  `json:"character"`
	CreditID      string  `json:"credit_id"`
	Department    string  `json:"department"`
	Job           string  `json:"job"`
	Order         int     `json:"order"`
}

// PersonMovieCredits is the response shape for
// GET /person/{person_id}/movie_credits.
type PersonMovieCredits struct {
	ID   int64               `json:"id"`
	Cast []PersonMovieCredit `json:"cast"`
	Crew []PersonMovieCredit `json:"crew"`
}

// PopularPerson is a single entry in the popular-people list.
type PopularPerson struct {
	ID                 int64                  `json:"id"`
	Adult              bool                   `json:"adult"`
	Gender             int                    `json:"gender"`
	KnownForDepartment string                 `json:"known_for_department"`
	Name               string                 `json:"name"`
	OriginalName       string                 `json:"original_name"`
	Popularity         float64                `json:"popularity"`
	ProfilePath        string                 `json:"profile_path"`
	KnownFor           []PersonCombinedCredit `json:"known_for"`
}

// PopularPersons is a paginated list of people, as returned by
// GET /person/popular.
type PopularPersons struct {
	Page         int             `json:"page"`
	Results      []PopularPerson `json:"results"`
	TotalPages   int             `json:"total_pages"`
	TotalResults int             `json:"total_results"`
}

// PersonTVCredit is a single cast/crew entry in a person's TV credits.
type PersonTVCredit struct {
	ID           int64   `json:"id"`
	Name         string  `json:"name"`
	OriginalName string  `json:"original_name"`
	Overview     string  `json:"overview"`
	PosterPath   string  `json:"poster_path"`
	FirstAirDate string  `json:"first_air_date"`
	Popularity   float64 `json:"popularity"`
	VoteAverage  float64 `json:"vote_average"`
	VoteCount    int64   `json:"vote_count"`
	Character    string  `json:"character"`
	CreditID     string  `json:"credit_id"`
	Department   string  `json:"department"`
	Job          string  `json:"job"`
	EpisodeCount int     `json:"episode_count"`
}

// PersonTVCredits is the response shape for
// GET /person/{person_id}/tv_credits.
type PersonTVCredits struct {
	ID   int64            `json:"id"`
	Cast []PersonTVCredit `json:"cast"`
	Crew []PersonTVCredit `json:"crew"`
}

// PersonTranslationData is the translated-field payload nested inside a
// PersonTranslation.
type PersonTranslationData struct {
	Biography string `json:"biography"`
}

// PersonTranslation is a single language's translation, as returned by
// GET /person/{person_id}/translations.
type PersonTranslation struct {
	Iso31661    string                `json:"iso_3166_1"`
	Iso6391     string                `json:"iso_639_1"`
	Name        string                `json:"name"`
	EnglishName string                `json:"english_name"`
	Data        PersonTranslationData `json:"data"`
}

// PersonTranslations is the response shape for
// GET /person/{person_id}/translations.
type PersonTranslations struct {
	ID           int64               `json:"id"`
	Translations []PersonTranslation `json:"translations"`
}

// PersonCombinedCredits is the response shape for
// GET /person/{person_id}/combined_credits.
type PersonCombinedCredits struct {
	ID   int64                  `json:"id"`
	Cast []PersonCombinedCredit `json:"cast"`
	Crew []PersonCombinedCredit `json:"crew"`
}
