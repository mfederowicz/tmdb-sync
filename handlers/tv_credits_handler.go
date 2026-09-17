package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
)

// TVCreditsHandler handles `tv -a credits -i <series_id>`.
type TVCreditsHandler struct {
	SeriesID int64
	Language string
}

// Handle fetches the cast and crew for a TV series.
func (h TVCreditsHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	credits, _, err := client.TV.GetCredits(ctx, h.SeriesID, h.Language)
	if err != nil {
		return nil, err
	}
	return credits, nil
}
