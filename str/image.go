package str

// Image is a single poster/backdrop/logo entry, as returned by TMDB's
// per-resource images endpoints.
type Image struct {
	AspectRatio float64 `json:"aspect_ratio"`
	Height      int     `json:"height"`
	Iso6391     string  `json:"iso_639_1"`
	FilePath    string  `json:"file_path"`
	VoteAverage float64 `json:"vote_average"`
	VoteCount   int64   `json:"vote_count"`
	Width       int     `json:"width"`
}

// Images is the response shape shared by TMDB's per-resource images
// endpoints (e.g. collection/{id}/images, movie/{id}/images).
type Images struct {
	ID        int64   `json:"id"`
	Backdrops []Image `json:"backdrops"`
	Posters   []Image `json:"posters"`
	Logos     []Image `json:"logos,omitempty"`
}
