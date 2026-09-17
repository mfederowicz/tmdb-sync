package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
)

// TVAggregateCreditsHandler handles `tv -a aggregate-credits -i <series_id>`.
type TVAggregateCreditsHandler struct {
	SeriesID int64
	Language string
}

// Handle fetches the aggregated cast and crew for a TV series.
func (h TVAggregateCreditsHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	credits, _, err := client.TV.GetAggregateCredits(ctx, h.SeriesID, h.Language)
	if err != nil {
		return nil, err
	}
	return credits, nil
}
