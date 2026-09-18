package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
)

// TVEpisodesVideosHandler handles `tv-episodes -a videos -i <series_id> -s <season_number> -e <episode_number>`.
type TVEpisodesVideosHandler struct {
	SeriesID      int64
	SeasonNumber  int
	EpisodeNumber int
	Language      string
}

// Handle fetches the videos for a single TV episode.
func (h TVEpisodesVideosHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	videos, _, err := client.TVEpisodes.GetVideos(ctx, h.SeriesID, h.SeasonNumber, h.EpisodeNumber, h.Language)
	if err != nil {
		return nil, err
	}
	return videos, nil
}
