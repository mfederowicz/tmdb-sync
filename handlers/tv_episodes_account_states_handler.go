package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
)

// TVEpisodesAccountStatesHandler handles `tv-episodes -a account-states -i <series_id> -s <season_number> -e <episode_number>`.
type TVEpisodesAccountStatesHandler struct {
	SeriesID      int64
	SeasonNumber  int
	EpisodeNumber int
	SessionID     string
}

// Handle fetches an account's rated status for a single TV episode.
func (h TVEpisodesAccountStatesHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	states, _, err := client.TVEpisodes.GetAccountStates(ctx, h.SeriesID, h.SeasonNumber, h.EpisodeNumber, h.SessionID)
	if err != nil {
		return nil, err
	}
	return states, nil
}
