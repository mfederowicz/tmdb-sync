package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
)

// TVEpisodeGroupsHandler handles `tv -a episode-groups -i <series_id>`.
type TVEpisodeGroupsHandler struct {
	SeriesID int64
}

// Handle fetches the episode groups for a TV series.
func (h TVEpisodeGroupsHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	groups, _, err := client.TV.GetEpisodeGroups(ctx, h.SeriesID)
	if err != nil {
		return nil, err
	}
	return groups, nil
}
