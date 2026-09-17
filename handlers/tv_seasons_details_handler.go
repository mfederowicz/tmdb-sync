package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
)

// TVSeasonsDetailsHandler handles `tv-seasons -a details -i <series_id> -s <season_number>`.
type TVSeasonsDetailsHandler struct {
	SeriesID     int64
	SeasonNumber int
}

// Handle fetches details for a single TV season.
func (h TVSeasonsDetailsHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	season, _, err := client.TVSeasons.GetTVSeason(ctx, h.SeriesID, h.SeasonNumber)
	if err != nil {
		return nil, err
	}
	return season, nil
}
