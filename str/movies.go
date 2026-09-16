package str

// Movies is a paginated list of movies, as returned by e.g. GET /movie/popular.
type Movies struct {
	Page         int     `json:"page"`
	Results      []Movie `json:"results"`
	TotalPages   int     `json:"total_pages"`
	TotalResults int     `json:"total_results"`
}
