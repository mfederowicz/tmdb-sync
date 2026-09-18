package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
)

// TVEpisodesDetailsHandler handles `tv-episodes -a details -i <series_id> -s <season_number> -e <episode_number>`.
type TVEpisodesDetailsHandler struct {
	SeriesID      int64
	SeasonNumber  int
	EpisodeNumber int
}

// Handle fetches details for a single TV episode.
func (h TVEpisodesDetailsHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	episode, _, err := client.TVEpisodes.GetTVEpisode(ctx, h.SeriesID, h.SeasonNumber, h.EpisodeNumber)
	if err != nil {
		return nil, err
	}
	return episode, nil
}
