package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
)

// TVEpisodesAddRatingHandler handles `tv-episodes -a add-rating -i <series_id> -s <season_number> -e <episode_number> -value <rating>`.
type TVEpisodesAddRatingHandler struct {
	SeriesID       int64
	SeasonNumber   int
	EpisodeNumber  int
	SessionID      string
	GuestSessionID string
	Value          float64
}

// Handle rates a TV episode on behalf of the session's account, or a guest
// session when GuestSessionID is set instead of SessionID.
func (h TVEpisodesAddRatingHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	status, _, err := client.TVEpisodes.AddRating(ctx, h.SeriesID, h.SeasonNumber, h.EpisodeNumber, h.SessionID, h.GuestSessionID, h.Value)
	if err != nil {
		return nil, err
	}
	return status, nil
}
