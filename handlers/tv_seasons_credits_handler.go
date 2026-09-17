package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
)

// TVSeasonsCreditsHandler handles `tv-seasons -a credits -i <series_id> -s <season_number>`.
type TVSeasonsCreditsHandler struct {
	SeriesID     int64
	SeasonNumber int
	Language     string
}

// Handle fetches the cast and crew for a TV season.
func (h TVSeasonsCreditsHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	credits, _, err := client.TVSeasons.GetCredits(ctx, h.SeriesID, h.SeasonNumber, h.Language)
	if err != nil {
		return nil, err
	}
	return credits, nil
}
