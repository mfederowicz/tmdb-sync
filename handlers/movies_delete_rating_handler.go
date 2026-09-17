package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
)

// MoviesDeleteRatingHandler handles `movies -a delete-rating -i <movie_id>`.
type MoviesDeleteRatingHandler struct {
	MovieID   int64
	SessionID string
}

// Handle removes the session's account's rating for a movie.
func (h MoviesDeleteRatingHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	status, _, err := client.Movies.DeleteRating(ctx, h.MovieID, h.SessionID)
	if err != nil {
		return nil, err
	}
	return status, nil
}
