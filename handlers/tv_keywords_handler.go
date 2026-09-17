package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
)

// TVKeywordsHandler handles `tv -a keywords -i <series_id>`.
type TVKeywordsHandler struct {
	SeriesID int64
}

// Handle fetches the keywords for a TV series.
func (h TVKeywordsHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	keywords, _, err := client.TV.GetKeywords(ctx, h.SeriesID)
	if err != nil {
		return nil, err
	}
	return keywords, nil
}
