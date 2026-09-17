package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
)

// TVWatchProvidersHandler handles `tv -a watch-providers -i <series_id>`.
type TVWatchProvidersHandler struct {
	SeriesID int64
}

// Handle fetches the per-region watch providers for a TV series.
func (h TVWatchProvidersHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	providers, _, err := client.TV.GetWatchProviders(ctx, h.SeriesID)
	if err != nil {
		return nil, err
	}
	return providers, nil
}
