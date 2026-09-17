package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
)

// TVSeasonsAccountStatesHandler handles `tv-seasons -a account-states -i <series_id> -s <season_number>`.
type TVSeasonsAccountStatesHandler struct {
	SeriesID     int64
	SeasonNumber int
	SessionID    string
}

// Handle fetches an account's rated status for every episode in a TV season.
func (h TVSeasonsAccountStatesHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	states, _, err := client.TVSeasons.GetAccountStates(ctx, h.SeriesID, h.SeasonNumber, h.SessionID)
	if err != nil {
		return nil, err
	}
	return states, nil
}
