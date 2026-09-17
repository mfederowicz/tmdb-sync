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

// TVCastRole is a single role a cast member played across episodes, as
// returned inline on the tv aggregate-credits endpoint's cast entries.
type TVCastRole struct {
	CreditID     string `json:"credit_id"`
	Character    string `json:"character"`
	EpisodeCount int    `json:"episode_count"`
}

// TVCrewJob is a single job a crew member held across episodes, as returned
// inline on the tv aggregate-credits endpoint's crew entries.
type TVCrewJob struct {
	CreditID     string `json:"credit_id"`
	Job          string `json:"job"`
	EpisodeCount int    `json:"episode_count"`
}

// TVAggregateCastMember is one cast entry, as returned inline on the tv
// aggregate-credits endpoint. Unlike the plain credits endpoint, a cast
// member's roles are aggregated across every episode they appeared in.
type TVAggregateCastMember struct {
	Adult              bool         `json:"adult"`
	Gender             int          `json:"gender"`
	ID                 int64        `json:"id"`
	KnownForDepartment string       `json:"known_for_department"`
	Name               string       `json:"name"`
	OriginalName       string       `json:"original_name"`
	Popularity         float64      `json:"popularity"`
	ProfilePath        string       `json:"profile_path"`
	Roles              []TVCastRole `json:"roles"`
	TotalEpisodeCount  int          `json:"total_episode_count"`
	Order              int          `json:"order"`
}

// TVAggregateCrewMember is one crew entry, as returned inline on the tv
// aggregate-credits endpoint. Unlike the plain credits endpoint, a crew
// member's jobs are aggregated across every episode they worked on.
type TVAggregateCrewMember struct {
	Adult              bool        `json:"adult"`
	Gender             int         `json:"gender"`
	ID                 int64       `json:"id"`
	KnownForDepartment string      `json:"known_for_department"`
	Name               string      `json:"name"`
	OriginalName       string      `json:"original_name"`
	Popularity         float64     `json:"popularity"`
	ProfilePath        string      `json:"profile_path"`
	Jobs               []TVCrewJob `json:"jobs"`
	Department         string      `json:"department"`
	TotalEpisodeCount  int         `json:"total_episode_count"`
}

// TVAggregateCredits is the response shape for
// GET /tv/{series_id}/aggregate_credits.
type TVAggregateCredits struct {
	ID   int64                   `json:"id"`
	Cast []TVAggregateCastMember `json:"cast"`
	Crew []TVAggregateCrewMember `json:"crew"`
}

// TVAlternativeTitle is one alternative title, as returned inline on the tv
// alternative-titles endpoint.
type TVAlternativeTitle struct {
	Iso31661 string `json:"iso_3166_1"`
	Title    string `json:"title"`
	Type     string `json:"type"`
}

// TVAlternativeTitles is the response shape for
// GET /tv/{series_id}/alternative_titles.
type TVAlternativeTitles struct {
	ID      int64                `json:"id"`
	Results []TVAlternativeTitle `json:"results"`
}

// TVContentRating is one country's content rating (certification), as
// returned inline on the tv content-ratings endpoint.
type TVContentRating struct {
	Iso31661    string   `json:"iso_3166_1"`
	Rating      string   `json:"rating"`
	Descriptors []string `json:"descriptors"`
}

// TVContentRatings is the response shape for
// GET /tv/{series_id}/content_ratings.
type TVContentRatings struct {
	ID      int64             `json:"id"`
	Results []TVContentRating `json:"results"`
}

// TVCastMember is one cast entry, as returned inline on the tv credits
// endpoint.
type TVCastMember struct {
	Adult              bool    `json:"adult"`
	Gender             int     `json:"gender"`
	ID                 int64   `json:"id"`
	KnownForDepartment string  `json:"known_for_department"`
	Name               string  `json:"name"`
	OriginalName       string  `json:"original_name"`
	Popularity         float64 `json:"popularity"`
	ProfilePath        string  `json:"profile_path"`
	Character          string  `json:"character"`
	CreditID           string  `json:"credit_id"`
	Order              int     `json:"order"`
}

// TVCrewMember is one crew entry, as returned inline on the tv credits
// endpoint.
type TVCrewMember struct {
	Adult              bool    `json:"adult"`
	Gender             int     `json:"gender"`
	ID                 int64   `json:"id"`
	KnownForDepartment string  `json:"known_for_department"`
	Name               string  `json:"name"`
	OriginalName       string  `json:"original_name"`
	Popularity         float64 `json:"popularity"`
	ProfilePath        string  `json:"profile_path"`
	CreditID           string  `json:"credit_id"`
	Department         string  `json:"department"`
	Job                string  `json:"job"`
}

// TVCredits is the response shape for GET /tv/{series_id}/credits.
type TVCredits struct {
	ID   int64          `json:"id"`
	Cast []TVCastMember `json:"cast"`
	Crew []TVCrewMember `json:"crew"`
}

// TVEpisodeGroupNetwork is the network associated with an episode group, as
// returned inline on the tv episode-groups endpoint.
type TVEpisodeGroupNetwork struct {
	ID            int64  `json:"id"`
	LogoPath      string `json:"logo_path"`
	Name          string `json:"name"`
	OriginCountry string `json:"origin_country"`
}

// TVEpisodeGroup is one episode-group entry, as returned by the tv
// episode-groups endpoint.
type TVEpisodeGroup struct {
	Description  string                `json:"description"`
	EpisodeCount int                   `json:"episode_count"`
	GroupCount   int                   `json:"group_count"`
	ID           string                `json:"id"`
	Name         string                `json:"name"`
	Network      TVEpisodeGroupNetwork `json:"network"`
	Type         int                   `json:"type"`
}

// TVEpisodeGroups is the response shape for GET /tv/{series_id}/episode_groups.
type TVEpisodeGroups struct {
	Results []TVEpisodeGroup `json:"results"`
	ID      int64            `json:"id"`
}

// TVExternalIDs is the response shape for GET /tv/{series_id}/external_ids.
type TVExternalIDs struct {
	ID          int64  `json:"id"`
	ImdbID      string `json:"imdb_id"`
	FreebaseMID string `json:"freebase_mid"`
	FreebaseID  string `json:"freebase_id"`
	TvdbID      int64  `json:"tvdb_id"`
	TvrageID    int64  `json:"tvrage_id"`
	WikidataID  string `json:"wikidata_id"`
	FacebookID  string `json:"facebook_id"`
	InstagramID string `json:"instagram_id"`
	TwitterID   string `json:"twitter_id"`
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

// TVKeywords is the response shape for GET /tv/{series_id}/keywords.
type TVKeywords struct {
	ID      int64     `json:"id"`
	Results []Keyword `json:"results"`
}

// TVLists is a paginated page of the lists a TV series belongs to, as
// returned by GET /tv/{series_id}/lists.
type TVLists struct {
	ID           int64         `json:"id"`
	Page         int           `json:"page"`
	Results      []AccountList `json:"results"`
	TotalPages   int           `json:"total_pages"`
	TotalResults int           `json:"total_results"`
}

// TVReviewAuthorDetails is the author metadata nested inside a TV review.
type TVReviewAuthorDetails struct {
	Name       string  `json:"name"`
	Username   string  `json:"username"`
	AvatarPath string  `json:"avatar_path"`
	Rating     float64 `json:"rating"`
}

// TVReview is a single review, as returned by
// GET /tv/{series_id}/reviews.
type TVReview struct {
	Author        string                `json:"author"`
	AuthorDetails TVReviewAuthorDetails `json:"author_details"`
	Content       string                `json:"content"`
	CreatedAt     string                `json:"created_at"`
	ID            string                `json:"id"`
	UpdatedAt     string                `json:"updated_at"`
	URL           string                `json:"url"`
}

// TVReviews is a paginated page of a TV series' reviews, as returned by
// GET /tv/{series_id}/reviews.
type TVReviews struct {
	ID           int64      `json:"id"`
	Page         int        `json:"page"`
	Results      []TVReview `json:"results"`
	TotalPages   int        `json:"total_pages"`
	TotalResults int        `json:"total_results"`
}
