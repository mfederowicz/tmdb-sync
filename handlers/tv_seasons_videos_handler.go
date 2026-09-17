package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
)

// TVSeasonsVideosHandler handles `tv-seasons -a videos -i <series_id> -s <season_number>`.
type TVSeasonsVideosHandler struct {
	SeriesID     int64
	SeasonNumber int
	Language     string
}

// Handle fetches the videos (trailers, teasers, ...) for a TV season.
func (h TVSeasonsVideosHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	videos, _, err := client.TVSeasons.GetVideos(ctx, h.SeriesID, h.SeasonNumber, h.Language)
	if err != nil {
		return nil, err
	}
	return videos, nil
}
