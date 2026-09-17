package str

// Keyword is the response shape for the keyword details endpoint.
type Keyword struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}
