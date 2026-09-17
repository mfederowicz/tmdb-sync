package str

// ReviewAuthorDetails is the nested author metadata on a review.
type ReviewAuthorDetails struct {
	Name       string  `json:"name"`
	Username   string  `json:"username"`
	AvatarPath string  `json:"avatar_path"`
	Rating     float64 `json:"rating"`
}

// Review is the response shape for the review details endpoint.
type Review struct {
	ID            string              `json:"id"`
	Author        string              `json:"author"`
	AuthorDetails ReviewAuthorDetails `json:"author_details"`
	Content       string              `json:"content"`
	CreatedAt     string              `json:"created_at"`
	Iso6391       string              `json:"iso_639_1"`
	MediaID       int64               `json:"media_id"`
	MediaTitle    string              `json:"media_title"`
	MediaType     string              `json:"media_type"`
	UpdatedAt     string              `json:"updated_at"`
	URL           string              `json:"url"`
}
