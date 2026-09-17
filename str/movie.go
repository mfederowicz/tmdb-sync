package str

// Movie is the response of GET /movie/{movie_id}.
type Movie struct {
	ID               int64   `json:"id"`
	Title            string  `json:"title"`
	OriginalTitle    string  `json:"original_title"`
	Overview         string  `json:"overview"`
	ReleaseDate      string  `json:"release_date"`
	Runtime          int     `json:"runtime"`
	VoteAverage      float64 `json:"vote_average"`
	VoteCount        int     `json:"vote_count"`
	Popularity       float64 `json:"popularity"`
	Adult            bool    `json:"adult"`
	Video            bool    `json:"video"`
	OriginalLanguage string  `json:"original_language"`
}

// RatedMovie is a movie as returned by an account's rated-movies endpoint,
// which adds the account's own rating to the usual movie fields.
type RatedMovie struct {
	Movie
	Rating float64 `json:"rating"`
}

// RatedMovies is a paginated list of an account's rated movies.
type RatedMovies struct {
	Page         int          `json:"page"`
	Results      []RatedMovie `json:"results"`
	TotalPages   int          `json:"total_pages"`
	TotalResults int          `json:"total_results"`
}

// MovieRatingRequest is the body of POST /movie/{movie_id}/rating.
type MovieRatingRequest struct {
	Value float64 `json:"value"`
}

// MovieAccountStates is the response shape for GET /movie/{movie_id}/account_states.
// Rated is `any` because TMDB returns either `false` (not rated) or an object
// like {"value": 8} (rated) for the same field.
type MovieAccountStates struct {
	ID        int64 `json:"id"`
	Favorite  bool  `json:"favorite"`
	Rated     any   `json:"rated"`
	Watchlist bool  `json:"watchlist"`
}

// MovieAlternativeTitle is one alternative title, as returned inline on the
// movie alternative-titles endpoint.
type MovieAlternativeTitle struct {
	Iso31661 string `json:"iso_3166_1"`
	Title    string `json:"title"`
	Type     string `json:"type"`
}

// MovieAlternativeTitles is the response shape for
// GET /movie/{movie_id}/alternative_titles.
type MovieAlternativeTitles struct {
	ID     int64                   `json:"id"`
	Titles []MovieAlternativeTitle `json:"titles"`
}

// MovieCastMember is one cast entry, as returned inline on the movie
// credits endpoint.
type MovieCastMember struct {
	Adult              bool    `json:"adult"`
	Gender             int     `json:"gender"`
	ID                 int64   `json:"id"`
	KnownForDepartment string  `json:"known_for_department"`
	Name               string  `json:"name"`
	OriginalName       string  `json:"original_name"`
	Popularity         float64 `json:"popularity"`
	ProfilePath        string  `json:"profile_path"`
	CastID             int64   `json:"cast_id"`
	Character          string  `json:"character"`
	CreditID           string  `json:"credit_id"`
	Order              int     `json:"order"`
}

// MovieCrewMember is one crew entry, as returned inline on the movie
// credits endpoint.
type MovieCrewMember struct {
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

// MovieCredits is the response shape for GET /movie/{movie_id}/credits.
type MovieCredits struct {
	ID   int64             `json:"id"`
	Cast []MovieCastMember `json:"cast"`
	Crew []MovieCrewMember `json:"crew"`
}

// MovieLists is a paginated page of the lists a movie belongs to, as
// returned by GET /movie/{movie_id}/lists.
type MovieLists struct {
	ID           int64         `json:"id"`
	Page         int           `json:"page"`
	Results      []AccountList `json:"results"`
	TotalPages   int           `json:"total_pages"`
	TotalResults int           `json:"total_results"`
}

// MovieKeywords is the response shape for GET /movie/{movie_id}/keywords.
type MovieKeywords struct {
	ID       int64     `json:"id"`
	Keywords []Keyword `json:"keywords"`
}

// MovieReleaseDate is one release date entry within a country's release
// dates, as returned inline on the movie release-dates endpoint.
type MovieReleaseDate struct {
	Certification string   `json:"certification"`
	Descriptors   []string `json:"descriptors"`
	Iso6391       string   `json:"iso_639_1"`
	Note          string   `json:"note"`
	ReleaseDate   string   `json:"release_date"`
	Type          int      `json:"type"`
}

// MovieReleaseDatesResult is a single country's release dates, as returned
// inline on the movie release-dates endpoint.
type MovieReleaseDatesResult struct {
	Iso31661     string             `json:"iso_3166_1"`
	ReleaseDates []MovieReleaseDate `json:"release_dates"`
}

// MovieReleaseDates is the response shape for
// GET /movie/{movie_id}/release_dates.
type MovieReleaseDates struct {
	ID      int64                     `json:"id"`
	Results []MovieReleaseDatesResult `json:"results"`
}

// MovieReviewAuthorDetails is the author metadata nested inside a review.
type MovieReviewAuthorDetails struct {
	Name       string  `json:"name"`
	Username   string  `json:"username"`
	AvatarPath string  `json:"avatar_path"`
	Rating     float64 `json:"rating"`
}

// MovieReview is a single review, as returned by
// GET /movie/{movie_id}/reviews.
type MovieReview struct {
	Author        string                   `json:"author"`
	AuthorDetails MovieReviewAuthorDetails `json:"author_details"`
	Content       string                   `json:"content"`
	CreatedAt     string                   `json:"created_at"`
	ID            string                   `json:"id"`
	UpdatedAt     string                   `json:"updated_at"`
	URL           string                   `json:"url"`
}

// MovieReviews is a paginated page of a movie's reviews, as returned by
// GET /movie/{movie_id}/reviews.
type MovieReviews struct {
	ID           int64         `json:"id"`
	Page         int           `json:"page"`
	Results      []MovieReview `json:"results"`
	TotalPages   int           `json:"total_pages"`
	TotalResults int           `json:"total_results"`
}

// MovieVideo is a single video entry, as returned by
// GET /movie/{movie_id}/videos.
type MovieVideo struct {
	Iso6391     string `json:"iso_639_1"`
	Iso31661    string `json:"iso_3166_1"`
	Name        string `json:"name"`
	Key         string `json:"key"`
	Site        string `json:"site"`
	Size        int    `json:"size"`
	Type        string `json:"type"`
	Official    bool   `json:"official"`
	PublishedAt string `json:"published_at"`
	ID          string `json:"id"`
}

// MovieVideos is the response shape for GET /movie/{movie_id}/videos.
type MovieVideos struct {
	ID      int64        `json:"id"`
	Results []MovieVideo `json:"results"`
}

// WatchProvider is a single streaming/rental/purchase provider entry, as
// returned inline on TMDB's per-resource watch-providers endpoints (e.g.
// movie/{id}/watch/providers, tv/{id}/watch/providers).
type WatchProvider struct {
	LogoPath        string `json:"logo_path"`
	ProviderID      int64  `json:"provider_id"`
	ProviderName    string `json:"provider_name"`
	DisplayPriority int    `json:"display_priority"`
}

// WatchProviderRegion is one country's watch-provider offers, keyed by
// ISO 3166-1 country code in the parent Results map.
type WatchProviderRegion struct {
	Link     string          `json:"link"`
	Flatrate []WatchProvider `json:"flatrate,omitempty"`
	Rent     []WatchProvider `json:"rent,omitempty"`
	Buy      []WatchProvider `json:"buy,omitempty"`
}

// MovieWatchProviders is the response shape for
// GET /movie/{movie_id}/watch/providers.
type MovieWatchProviders struct {
	ID      int64                          `json:"id"`
	Results map[string]WatchProviderRegion `json:"results"`
}

// MovieExternalIDs is the response shape for
// GET /movie/{movie_id}/external_ids.
type MovieExternalIDs struct {
	ID          int64  `json:"id"`
	ImdbID      string `json:"imdb_id"`
	WikidataID  string `json:"wikidata_id"`
	FacebookID  string `json:"facebook_id"`
	InstagramID string `json:"instagram_id"`
	TwitterID   string `json:"twitter_id"`
}
