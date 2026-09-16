package str

// Configuration is the response of GET /configuration.
type Configuration struct {
	Images     ImagesConfiguration `json:"images"`
	ChangeKeys []string            `json:"change_keys"`
}

// ImagesConfiguration describes TMDB's image CDN base URLs and sizes.
type ImagesConfiguration struct {
	BaseURL       string   `json:"base_url"`
	SecureBaseURL string   `json:"secure_base_url"`
	BackdropSizes []string `json:"backdrop_sizes"`
	LogoSizes     []string `json:"logo_sizes"`
	PosterSizes   []string `json:"poster_sizes"`
	ProfileSizes  []string `json:"profile_sizes"`
	StillSizes    []string `json:"still_sizes"`
}
