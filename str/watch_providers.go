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
