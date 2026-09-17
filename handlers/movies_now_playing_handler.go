package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
)

// MoviesNowPlayingHandler handles `movies -a now-playing`.
type MoviesNowPlayingHandler struct {
	PagesLimit int
}

// Handle fetches the current now-playing-movies list.
func (h MoviesNowPlayingHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	movies, err := client.Movies.GetNowPlayingMovies(ctx, h.PagesLimit)
	if err != nil {
		return nil, err
	}
	return movies, nil
}
