package str

// Trending is a paginated list of mixed movie/tv/person results, as returned
// by GET /trending/all/{time_window}.
type Trending struct {
	Page         int                 `json:"page"`
	Results      []SearchMultiResult `json:"results"`
	TotalPages   int                 `json:"total_pages"`
	TotalResults int                 `json:"total_results"`
}
