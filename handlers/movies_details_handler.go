package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
)

// MoviesDetailsHandler handles `movies -a details -i <movie_id>`.
type MoviesDetailsHandler struct {
	MovieID int64
}

// Handle fetches details for a single movie.
func (h MoviesDetailsHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	movie, _, err := client.Movies.GetMovie(ctx, h.MovieID)
	if err != nil {
		return nil, err
	}
	return movie, nil
}
