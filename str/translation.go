package str

// TranslationData is the translated-field payload nested inside a Translation.
type TranslationData struct {
	Name     string `json:"name"`
	Overview string `json:"overview"`
	Homepage string `json:"homepage"`
}

// Translation is a single language's translation, as returned by TMDB's
// per-resource translations endpoints.
type Translation struct {
	Iso31661    string          `json:"iso_3166_1"`
	Iso6391     string          `json:"iso_639_1"`
	Name        string          `json:"name"`
	EnglishName string          `json:"english_name"`
	Data        TranslationData `json:"data"`
}

// Translations is the response shape shared by TMDB's per-resource
// translations endpoints (e.g. collection/{id}/translations, movie/{id}/translations).
type Translations struct {
	ID           int64         `json:"id"`
	Translations []Translation `json:"translations"`
}
