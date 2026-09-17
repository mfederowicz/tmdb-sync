package str

// ChangeItem is one changed id, as returned by the movie/tv/person change
// list endpoints.
type ChangeItem struct {
	ID    int64 `json:"id"`
	Adult bool  `json:"adult"`
}

// Changes is a paginated list of changed ids.
type Changes struct {
	Page         int          `json:"page"`
	Results      []ChangeItem `json:"results"`
	TotalPages   int          `json:"total_pages"`
	TotalResults int          `json:"total_results"`
}
