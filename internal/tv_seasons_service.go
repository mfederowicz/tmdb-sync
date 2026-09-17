package internal

import (
	"context"
	"fmt"
	"net/http"

	"github.com/mfederowicz/tmdb-sync/str"
	"github.com/mfederowicz/tmdb-sync/uri"
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

// GetAccountStates fetches an account's rated status for every episode in a
// single TV season.
//
// Api docs: https://developer.themoviedb.org/reference/tv-season-account-states
func (s *TVSeasonsService) GetAccountStates(ctx context.Context, seriesID int64, seasonNumber int, sessionID string) (*str.TVSeasonAccountStates, *str.Response, error) {
	urlStr, err := uri.AddQuery(fmt.Sprintf("tv/%d/season/%d/account_states", seriesID, seasonNumber), &uri.AccountOptions{SessionID: sessionID})
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(http.MethodGet, urlStr, nil)
	if err != nil {
		return nil, nil, err
	}

	states := new(str.TVSeasonAccountStates)
	resp, err := s.client.Do(ctx, req, states)
	if err != nil {
		return nil, resp, err
	}

	return states, resp, nil
}

// GetAggregateCredits fetches the cast and crew for a single TV season,
// with roles/jobs aggregated across every episode.
//
// Api docs: https://developer.themoviedb.org/reference/tv-season-aggregate-credits
func (s *TVSeasonsService) GetAggregateCredits(ctx context.Context, seriesID int64, seasonNumber int, language string) (*str.TVAggregateCredits, *str.Response, error) {
	urlStr, err := uri.AddQuery(fmt.Sprintf("tv/%d/season/%d/aggregate_credits", seriesID, seasonNumber), &uri.LanguageOptions{Language: language})
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(http.MethodGet, urlStr, nil)
	if err != nil {
		return nil, nil, err
	}

	credits := new(str.TVAggregateCredits)
	resp, err := s.client.Do(ctx, req, credits)
	if err != nil {
		return nil, resp, err
	}

	return credits, resp, nil
}
