package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
)

// MoviesPopularHandler handles `movies -a popular`, fetching pages up to
// PagesLimit (0 = unlimited, bounded by TMDB's total_pages).
type MoviesPopularHandler struct {
	PagesLimit int
}

// Handle fetches the popular-movies list.
func (h MoviesPopularHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	movies, err := client.Movies.GetPopularMovies(ctx, h.PagesLimit)
	if err != nil {
		return nil, err
	}
	return movies, nil
}
