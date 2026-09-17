package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
)

// MoviesCreditsHandler handles `movies -a credits -i <movie_id>`.
type MoviesCreditsHandler struct {
	MovieID  int64
	Language string
}

// Handle fetches the cast and crew for a movie.
func (h MoviesCreditsHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	credits, _, err := client.Movies.GetCredits(ctx, h.MovieID, h.Language)
	if err != nil {
		return nil, err
	}
	return credits, nil
}
