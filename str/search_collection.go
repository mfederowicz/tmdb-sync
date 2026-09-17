package str

// SearchCollectionResult is a single collection entry within a search/collection response.
type SearchCollectionResult struct {
	Adult            bool   `json:"adult"`
	BackdropPath     string `json:"backdrop_path"`
	ID               int64  `json:"id"`
	Name             string `json:"name"`
	OriginalLanguage string `json:"original_language"`
	OriginalName     string `json:"original_name"`
	Overview         string `json:"overview"`
	PosterPath       string `json:"poster_path"`
}

// SearchCollections is a paginated list of collections, as returned by GET /search/collection.
type SearchCollections struct {
	Page         int                      `json:"page"`
	Results      []SearchCollectionResult `json:"results"`
	TotalPages   int                      `json:"total_pages"`
	TotalResults int                      `json:"total_results"`
}
