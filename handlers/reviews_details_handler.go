package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
)

// ReviewsDetailsHandler handles `reviews -a details -i <review_id>`.
type ReviewsDetailsHandler struct {
	ReviewID string
}

// Handle fetches details for a single review.
func (h ReviewsDetailsHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	review, _, err := client.Reviews.GetReview(ctx, h.ReviewID)
	if err != nil {
		return nil, err
	}
	return review, nil
}
