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

// Country is an entry of GET /configuration/countries.
type Country struct {
	ISO3166_1   string `json:"iso_3166_1"`
	EnglishName string `json:"english_name"`
	NativeName  string `json:"native_name"`
}

// JobDepartment is an entry of GET /configuration/jobs.
type JobDepartment struct {
	Department string   `json:"department"`
	Jobs       []string `json:"jobs"`
}

// Language is an entry of GET /configuration/languages.
type Language struct {
	ISO639_1    string `json:"iso_639_1"`
	EnglishName string `json:"english_name"`
	Name        string `json:"name"`
}

// TimezoneRegion is an entry of GET /configuration/timezones.
type TimezoneRegion struct {
	ISO3166_1 string   `json:"iso_3166_1"`
	Zones     []string `json:"zones"`
}
