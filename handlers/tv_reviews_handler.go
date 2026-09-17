package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
)

// TVReviewsHandler handles `tv -a reviews -i <series_id>`.
type TVReviewsHandler struct {
	SeriesID   int64
	Language   string
	PagesLimit int
}

// Handle fetches a TV series' reviews.
func (h TVReviewsHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	reviews, err := client.TV.GetReviews(ctx, h.SeriesID, h.Language, h.PagesLimit)
	if err != nil {
		return nil, err
	}
	return reviews, nil
}
