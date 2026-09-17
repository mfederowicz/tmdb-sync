package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
)

// TVExternalIDsHandler handles `tv -a external-ids -i <series_id>`.
type TVExternalIDsHandler struct {
	SeriesID int64
}

// Handle fetches the external ids for a TV series.
func (h TVExternalIDsHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	ids, _, err := client.TV.GetExternalIDs(ctx, h.SeriesID)
	if err != nil {
		return nil, err
	}
	return ids, nil
}
