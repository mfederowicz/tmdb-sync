package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
)

// TVScreenedTheatricallyHandler handles `tv -a screened-theatrically -i <series_id>`.
type TVScreenedTheatricallyHandler struct {
	SeriesID int64
}

// Handle fetches the episodes of a TV series that had a theatrical screening.
func (h TVScreenedTheatricallyHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	screened, _, err := client.TV.GetScreenedTheatrically(ctx, h.SeriesID)
	if err != nil {
		return nil, err
	}
	return screened, nil
}
