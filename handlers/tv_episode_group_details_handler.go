package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
)

// TVEpisodeGroupDetailsHandler handles `tv-episode-groups -a details -i <episode_group_id>`.
type TVEpisodeGroupDetailsHandler struct {
	ID string
}

// Handle fetches details for a single TV episode group.
func (h TVEpisodeGroupDetailsHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	group, _, err := client.TVEpisodeGroup.GetTVEpisodeGroup(ctx, h.ID)
	if err != nil {
		return nil, err
	}
	return group, nil
}
