package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
)

// TVDetailsHandler handles `tv -a details -i <series_id>`.
type TVDetailsHandler struct {
	SeriesID int64
}

// Handle fetches details for a single TV series.
func (h TVDetailsHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	tv, _, err := client.TV.GetTV(ctx, h.SeriesID)
	if err != nil {
		return nil, err
	}
	return tv, nil
}
