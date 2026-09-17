package internal

import (
	"context"
	"fmt"
	"net/http"

	"github.com/mfederowicz/tmdb-sync/str"
)

// ReviewsService handles communication with the /review endpoints of the TMDB API.
type ReviewsService Service

// GetReview fetches details for a single review by TMDB review id.
//
// Api docs: https://developer.themoviedb.org/reference/review-details
func (s *ReviewsService) GetReview(ctx context.Context, reviewID string) (*str.Review, *str.Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("review/%s", reviewID), nil)
	if err != nil {
		return nil, nil, err
	}

	review := new(str.Review)
	resp, err := s.client.Do(ctx, req, review)
	if err != nil {
		return nil, resp, err
	}

	return review, resp, nil
}
