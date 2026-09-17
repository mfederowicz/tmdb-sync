package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
)

// TVContentRatingsHandler handles `tv -a content-ratings -i <series_id>`.
type TVContentRatingsHandler struct {
	SeriesID int64
}

// Handle fetches the per-country content ratings for a TV series.
func (h TVContentRatingsHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	ratings, _, err := client.TV.GetContentRatings(ctx, h.SeriesID)
	if err != nil {
		return nil, err
	}
	return ratings, nil
}
