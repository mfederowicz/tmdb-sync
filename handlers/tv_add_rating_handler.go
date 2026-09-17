package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
)

// TVAddRatingHandler handles `tv -a add-rating -i <series_id> -value <rating>`.
type TVAddRatingHandler struct {
	SeriesID  int64
	SessionID string
	Value     float64
}

// Handle rates a TV series on behalf of the session's account.
func (h TVAddRatingHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	status, _, err := client.TV.AddRating(ctx, h.SeriesID, h.SessionID, h.Value)
	if err != nil {
		return nil, err
	}
	return status, nil
}
