package str

// WatchProviderRegionInfo is one entry in the available-regions endpoint's
// results, describing a watch-provider region TMDB supports.
type WatchProviderRegionInfo struct {
	ISO31661    string `json:"iso_3166_1"`
	EnglishName string `json:"english_name"`
	NativeName  string `json:"native_name"`
}

// WatchProviderRegions is the response shape for GET /watch/providers/regions.
type WatchProviderRegions struct {
	Results []WatchProviderRegionInfo `json:"results"`
}

// WatchProviderListEntry is one provider entry in the movie/tv watch
// providers list endpoints, including the per-region display priorities
// alongside the provider's own identity fields.
type WatchProviderListEntry struct {
	DisplayPriorities map[string]int `json:"display_priorities"`
	DisplayPriority   int            `json:"display_priority"`
	LogoPath          string         `json:"logo_path"`
	ProviderName      string         `json:"provider_name"`
	ProviderID        int64          `json:"provider_id"`
}

// MovieWatchProviderList is the response shape for GET /watch/providers/movie.
type MovieWatchProviderList struct {
	Results []WatchProviderListEntry `json:"results"`
}
