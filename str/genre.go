package str

// Genre is a single TMDB genre.
type Genre struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

// GenreList is the response shape for the genre list endpoints.
type GenreList struct {
	Genres []*Genre `json:"genres"`
}
