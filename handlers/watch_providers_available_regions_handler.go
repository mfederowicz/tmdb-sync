package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
)

// WatchProvidersAvailableRegionsHandler handles `watch-providers -a available-regions`.
type WatchProvidersAvailableRegionsHandler struct {
	Language string
}

// Handle fetches the list of regions TMDB has watch provider data for.
func (h WatchProvidersAvailableRegionsHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	regions, _, err := client.WatchProviders.GetAvailableRegions(ctx, h.Language)
	if err != nil {
		return nil, err
	}
	return regions, nil
}
