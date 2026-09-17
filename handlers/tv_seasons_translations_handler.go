package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
)

// TVSeasonsTranslationsHandler handles `tv-seasons -a translations -i <series_id> -s <season_number>`.
type TVSeasonsTranslationsHandler struct {
	SeriesID     int64
	SeasonNumber int
}

// Handle fetches the translations for a TV season.
func (h TVSeasonsTranslationsHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	translations, _, err := client.TVSeasons.GetTranslations(ctx, h.SeriesID, h.SeasonNumber)
	if err != nil {
		return nil, err
	}
	return translations, nil
}
