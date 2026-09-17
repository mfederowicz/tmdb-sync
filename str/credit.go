package str

// CreditMedia is the media (movie or TV show/episode) a credit is attached to.
type CreditMedia struct {
	ID            int64   `json:"id"`
	Name          string  `json:"name"`
	OriginalName  string  `json:"original_name"`
	Title         string  `json:"title"`
	OriginalTitle string  `json:"original_title"`
	Overview      string  `json:"overview"`
	PosterPath    string  `json:"poster_path"`
	Character     string  `json:"character"`
	Episodes      []int64 `json:"episodes"`
	Seasons       []int64 `json:"seasons"`
}

// CreditPerson is the person a credit is attached to.
type CreditPerson struct {
	ID           int64  `json:"id"`
	Name         string `json:"name"`
	OriginalName string `json:"original_name"`
	ProfilePath  string `json:"profile_path"`
}

// Credit is the response shape for the credit details endpoint.
type Credit struct {
	ID         string       `json:"id"`
	CreditType string       `json:"credit_type"`
	Department string       `json:"department"`
	Job        string       `json:"job"`
	MediaType  string       `json:"media_type"`
	Media      CreditMedia  `json:"media"`
	Person     CreditPerson `json:"person"`
}
