package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
)

// MoviesUpcomingHandler handles `movies -a upcoming`.
type MoviesUpcomingHandler struct {
	PagesLimit int
}

// Handle fetches the current upcoming-movies list.
func (h MoviesUpcomingHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	movies, err := client.Movies.GetUpcomingMovies(ctx, h.PagesLimit)
	if err != nil {
		return nil, err
	}
	return movies, nil
}
