package internal

import (
	"context"
	"fmt"
	"net/http"

	"github.com/mfederowicz/tmdb-sync/str"
)

// TVService handles communication with the /tv endpoints of the TMDB API.
type TVService Service

// GetTV fetches details for a single TV series by TMDB id.
//
// Api docs: https://developer.themoviedb.org/reference/tv-series-details
func (s *TVService) GetTV(ctx context.Context, seriesID int64) (*str.TV, *str.Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("tv/%d", seriesID), nil)
	if err != nil {
		return nil, nil, err
	}

	tv := new(str.TV)
	resp, err := s.client.Do(ctx, req, tv)
	if err != nil {
		return nil, resp, err
	}

	return tv, resp, nil
}
