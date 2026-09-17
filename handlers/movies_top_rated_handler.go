package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
)

// MoviesTopRatedHandler handles `movies -a top-rated`.
type MoviesTopRatedHandler struct {
	PagesLimit int
}

// Handle fetches the current top-rated-movies list.
func (h MoviesTopRatedHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	movies, err := client.Movies.GetTopRatedMovies(ctx, h.PagesLimit)
	if err != nil {
		return nil, err
	}
	return movies, nil
}
