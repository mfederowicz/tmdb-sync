package internal

import (
	"context"
	"fmt"
	"net/http"

	"github.com/mfederowicz/tmdb-sync/str"
	"github.com/mfederowicz/tmdb-sync/uri"
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

// GetAccountStates fetches an account's favorite/rated/watchlist status for
// a single TV series.
//
// Api docs: https://developer.themoviedb.org/reference/tv-series-account-states
func (s *TVService) GetAccountStates(ctx context.Context, seriesID int64, sessionID string) (*str.TVAccountStates, *str.Response, error) {
	urlStr, err := uri.AddQuery(fmt.Sprintf("tv/%d/account_states", seriesID), &uri.AccountOptions{SessionID: sessionID})
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(http.MethodGet, urlStr, nil)
	if err != nil {
		return nil, nil, err
	}

	states := new(str.TVAccountStates)
	resp, err := s.client.Do(ctx, req, states)
	if err != nil {
		return nil, resp, err
	}

	return states, resp, nil
}

// GetAggregateCredits fetches the cast and crew for a single TV series,
// with roles/jobs aggregated across every episode.
//
// Api docs: https://developer.themoviedb.org/reference/tv-series-aggregate-credits
func (s *TVService) GetAggregateCredits(ctx context.Context, seriesID int64, language string) (*str.TVAggregateCredits, *str.Response, error) {
	urlStr, err := uri.AddQuery(fmt.Sprintf("tv/%d/aggregate_credits", seriesID), &uri.LanguageOptions{Language: language})
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

// GetAlternativeTitles fetches the alternative titles for a single TV
// series.
//
// Api docs: https://developer.themoviedb.org/reference/tv-series-alternative-titles
func (s *TVService) GetAlternativeTitles(ctx context.Context, seriesID int64) (*str.TVAlternativeTitles, *str.Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("tv/%d/alternative_titles", seriesID), nil)
	if err != nil {
		return nil, nil, err
	}

	titles := new(str.TVAlternativeTitles)
	resp, err := s.client.Do(ctx, req, titles)
	if err != nil {
		return nil, resp, err
	}

	return titles, resp, nil
}
