package internal

import (
	"context"
	"fmt"
	"net/http"

	"github.com/mfederowicz/tmdb-sync/str"
)

// TVSeasonsService handles communication with the /tv/{series_id}/season
// endpoints of the TMDB API.
type TVSeasonsService Service

// GetTVSeason fetches details for a single TV season by series id and
// season number.
//
// Api docs: https://developer.themoviedb.org/reference/tv-season-details
func (s *TVSeasonsService) GetTVSeason(ctx context.Context, seriesID int64, seasonNumber int) (*str.TVSeason, *str.Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("tv/%d/season/%d", seriesID, seasonNumber), nil)
	if err != nil {
		return nil, nil, err
	}

	season := new(str.TVSeason)
	resp, err := s.client.Do(ctx, req, season)
	if err != nil {
		return nil, resp, err
	}

	return season, resp, nil
}
