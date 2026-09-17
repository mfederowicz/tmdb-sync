package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
)

// TVVideosHandler handles `tv -a videos -i <series_id>`.
type TVVideosHandler struct {
	SeriesID int64
	Language string
}

// Handle fetches the videos (trailers, teasers, ...) for a TV series.
func (h TVVideosHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	videos, _, err := client.TV.GetVideos(ctx, h.SeriesID, h.Language)
	if err != nil {
		return nil, err
	}
	return videos, nil
}
