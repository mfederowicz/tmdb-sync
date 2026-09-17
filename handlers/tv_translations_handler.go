package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
)

// TVTranslationsHandler handles `tv -a translations -i <series_id>`.
type TVTranslationsHandler struct {
	SeriesID int64
}

// Handle fetches the translated fields for a single TV series.
func (h TVTranslationsHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	translations, _, err := client.TV.GetTranslations(ctx, h.SeriesID)
	if err != nil {
		return nil, err
	}
	return translations, nil
}
