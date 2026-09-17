package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
)

// TVAccountStatesHandler handles `tv -a account-states -i <series_id>`.
type TVAccountStatesHandler struct {
	SeriesID  int64
	SessionID string
}

// Handle fetches an account's favorite/rated/watchlist status for a TV series.
func (h TVAccountStatesHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	states, _, err := client.TV.GetAccountStates(ctx, h.SeriesID, h.SessionID)
	if err != nil {
		return nil, err
	}
	return states, nil
}
