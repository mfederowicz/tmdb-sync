package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
)

// TVDeleteRatingHandler handles `tv -a delete-rating -i <series_id>`.
type TVDeleteRatingHandler struct {
	SeriesID  int64
	SessionID string
}

// Handle removes the session's account's rating for a TV series.
func (h TVDeleteRatingHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	status, _, err := client.TV.DeleteRating(ctx, h.SeriesID, h.SessionID)
	if err != nil {
		return nil, err
	}
	return status, nil
}
