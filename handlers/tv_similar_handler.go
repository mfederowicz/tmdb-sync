package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
)

// TVSimilarHandler handles `tv -a similar -i <series_id>`.
type TVSimilarHandler struct {
	SeriesID   int64
	Language   string
	PagesLimit int
}

// Handle fetches TV series similar to a single series.
func (h TVSimilarHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	shows, err := client.TV.GetSimilar(ctx, h.SeriesID, h.Language, h.PagesLimit)
	if err != nil {
		return nil, err
	}
	return shows, nil
}
