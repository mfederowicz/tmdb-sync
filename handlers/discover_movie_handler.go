package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
	"github.com/mfederowicz/tmdb-sync/uri"
)

// DiscoverMovieHandler handles `discover -a movie [filters]`, fetching pages
// up to PagesLimit (0 = unlimited, bounded by TMDB's total_pages).
type DiscoverMovieHandler struct {
	Options    uri.DiscoverMovieOptions
	PagesLimit int
}

// Handle fetches the discover-movie list matching Options.
func (h DiscoverMovieHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	movies, err := client.Discover.GetDiscoverMovies(ctx, h.Options, h.PagesLimit)
	if err != nil {
		return nil, err
	}
	return movies, nil
}
