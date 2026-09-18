package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
)

// WatchProvidersMovieProvidersHandler handles `watch-providers -a movie-providers`.
type WatchProvidersMovieProvidersHandler struct {
	Language    string
	WatchRegion string
}

// Handle fetches the list of watch providers for movies.
func (h WatchProvidersMovieProvidersHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	providers, _, err := client.WatchProviders.GetMovieProviders(ctx, h.Language, h.WatchRegion)
	if err != nil {
		return nil, err
	}
	return providers, nil
}
