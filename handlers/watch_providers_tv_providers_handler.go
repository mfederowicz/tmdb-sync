package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
)

// WatchProvidersTVProvidersHandler handles `watch-providers -a tv-providers`.
type WatchProvidersTVProvidersHandler struct {
	Language    string
	WatchRegion string
}

// Handle fetches the list of watch providers for TV series.
func (h WatchProvidersTVProvidersHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	providers, _, err := client.WatchProviders.GetTVProviders(ctx, h.Language, h.WatchRegion)
	if err != nil {
		return nil, err
	}
	return providers, nil
}
