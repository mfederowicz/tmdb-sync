package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
)

// TVEpisodesCreditsHandler handles `tv-episodes -a credits -i <series_id> -s <season_number> -e <episode_number>`.
type TVEpisodesCreditsHandler struct {
	SeriesID      int64
	SeasonNumber  int
	EpisodeNumber int
	Language      string
}

// Handle fetches the cast and crew for a single TV episode.
func (h TVEpisodesCreditsHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	credits, _, err := client.TVEpisodes.GetCredits(ctx, h.SeriesID, h.SeasonNumber, h.EpisodeNumber, h.Language)
	if err != nil {
		return nil, err
	}
	return credits, nil
}
