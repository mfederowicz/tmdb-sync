package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
)

// TVEpisodesExternalIDsHandler handles `tv-episodes -a external-ids -i <series_id> -s <season_number> -e <episode_number>`.
type TVEpisodesExternalIDsHandler struct {
	SeriesID      int64
	SeasonNumber  int
	EpisodeNumber int
}

// Handle fetches the external ids for a single TV episode.
func (h TVEpisodesExternalIDsHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	ids, _, err := client.TVEpisodes.GetExternalIDs(ctx, h.SeriesID, h.SeasonNumber, h.EpisodeNumber)
	if err != nil {
		return nil, err
	}
	return ids, nil
}
