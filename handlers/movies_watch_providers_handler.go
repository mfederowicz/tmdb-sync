package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
)

// MoviesWatchProvidersHandler handles `movies -a watch-providers -i <movie_id>`.
type MoviesWatchProvidersHandler struct {
	MovieID int64
}

// Handle fetches the per-region watch providers for a movie.
func (h MoviesWatchProvidersHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	providers, _, err := client.Movies.GetWatchProviders(ctx, h.MovieID)
	if err != nil {
		return nil, err
	}
	return providers, nil
}
