package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
)

// TVEpisodesDeleteRatingHandler handles `tv-episodes -a delete-rating -i <series_id> -s <season_number> -e <episode_number>`.
type TVEpisodesDeleteRatingHandler struct {
	SeriesID      int64
	SeasonNumber  int
	EpisodeNumber int
	SessionID     string
}

// Handle removes the session's account's rating for a TV episode.
func (h TVEpisodesDeleteRatingHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	status, _, err := client.TVEpisodes.DeleteRating(ctx, h.SeriesID, h.SeasonNumber, h.EpisodeNumber, h.SessionID)
	if err != nil {
		return nil, err
	}
	return status, nil
}
