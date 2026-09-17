package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
)

// TVSeasonsExternalIDsHandler handles `tv-seasons -a external-ids -i <series_id> -s <season_number>`.
type TVSeasonsExternalIDsHandler struct {
	SeriesID     int64
	SeasonNumber int
}

// Handle fetches the external ids for a TV season.
func (h TVSeasonsExternalIDsHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	ids, _, err := client.TVSeasons.GetExternalIDs(ctx, h.SeriesID, h.SeasonNumber)
	if err != nil {
		return nil, err
	}
	return ids, nil
}
