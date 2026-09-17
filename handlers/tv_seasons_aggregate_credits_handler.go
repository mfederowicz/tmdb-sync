package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
)

// TVSeasonsAggregateCreditsHandler handles `tv-seasons -a aggregate-credits -i <series_id> -s <season_number>`.
type TVSeasonsAggregateCreditsHandler struct {
	SeriesID     int64
	SeasonNumber int
	Language     string
}

// Handle fetches the cast and crew for a TV season, aggregated across every episode.
func (h TVSeasonsAggregateCreditsHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	credits, _, err := client.TVSeasons.GetAggregateCredits(ctx, h.SeriesID, h.SeasonNumber, h.Language)
	if err != nil {
		return nil, err
	}
	return credits, nil
}
