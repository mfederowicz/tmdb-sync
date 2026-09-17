package str

// SearchKeywordResult is a single keyword entry within a search/keyword response.
type SearchKeywordResult struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

// SearchKeywords is a paginated list of keywords, as returned by GET /search/keyword.
type SearchKeywords struct {
	Page         int                   `json:"page"`
	Results      []SearchKeywordResult `json:"results"`
	TotalPages   int                   `json:"total_pages"`
	TotalResults int                   `json:"total_results"`
}
