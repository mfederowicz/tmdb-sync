package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
)

// TVAlternativeTitlesHandler handles `tv -a alternative-titles -i <series_id>`.
type TVAlternativeTitlesHandler struct {
	SeriesID int64
}

// Handle fetches the alternative titles for a TV series.
func (h TVAlternativeTitlesHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	titles, _, err := client.TV.GetAlternativeTitles(ctx, h.SeriesID)
	if err != nil {
		return nil, err
	}
	return titles, nil
}
