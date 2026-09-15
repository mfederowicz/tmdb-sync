package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
)

// MoviesPopularHandler handles `movies -a popular -p <page>`.
type MoviesPopularHandler struct {
	Page int
}

// Handle fetches a page of the popular-movies list.
func (h MoviesPopularHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	movies, _, err := client.Movies.GetPopularMovies(ctx, h.Page)
	if err != nil {
		return nil, err
	}
	return movies, nil
}
