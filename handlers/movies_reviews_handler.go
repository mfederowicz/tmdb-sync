package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
)

// MoviesReviewsHandler handles `movies -a reviews -i <movie_id>`.
type MoviesReviewsHandler struct {
	MovieID    int64
	Language   string
	PagesLimit int
}

// Handle fetches a movie's reviews.
func (h MoviesReviewsHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	reviews, err := client.Movies.GetReviews(ctx, h.MovieID, h.Language, h.PagesLimit)
	if err != nil {
		return nil, err
	}
	return reviews, nil
}
