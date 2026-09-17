package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
)

// TVRecommendationsHandler handles `tv -a recommendations -i <series_id>`.
type TVRecommendationsHandler struct {
	SeriesID   int64
	Language   string
	PagesLimit int
}

// Handle fetches TV series recommended off a single series.
func (h TVRecommendationsHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	shows, err := client.TV.GetRecommendations(ctx, h.SeriesID, h.Language, h.PagesLimit)
	if err != nil {
		return nil, err
	}
	return shows, nil
}
