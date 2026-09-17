package internal

import (
	"context"
	"fmt"
	"net/http"

	"github.com/mfederowicz/tmdb-sync/str"
)

// CreditsService handles communication with the /credit endpoints of the TMDB API.
type CreditsService Service

// GetCredit fetches details for a single credit by TMDB credit id.
//
// Api docs: https://developer.themoviedb.org/reference/credit-details
func (s *CreditsService) GetCredit(ctx context.Context, creditID string) (*str.Credit, *str.Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("credit/%s", creditID), nil)
	if err != nil {
		return nil, nil, err
	}

	credit := new(str.Credit)
	resp, err := s.client.Do(ctx, req, credit)
	if err != nil {
		return nil, resp, err
	}

	return credit, resp, nil
}
