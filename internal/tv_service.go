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

// GetEpisodeGroups fetches the episode groups for a single TV series.
//
// Api docs: https://developer.themoviedb.org/reference/tv-series-episode-groups
func (s *TVService) GetEpisodeGroups(ctx context.Context, seriesID int64) (*str.TVEpisodeGroups, *str.Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("tv/%d/episode_groups", seriesID), nil)
	if err != nil {
		return nil, nil, err
	}

	groups := new(str.TVEpisodeGroups)
	resp, err := s.client.Do(ctx, req, groups)
	if err != nil {
		return nil, resp, err
	}

	return groups, resp, nil
}

// GetExternalIDs fetches a single TV series's external ids (IMDb, TVDB, ...).
//
// Api docs: https://developer.themoviedb.org/reference/tv-series-external-ids
func (s *TVService) GetExternalIDs(ctx context.Context, seriesID int64) (*str.TVExternalIDs, *str.Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("tv/%d/external_ids", seriesID), nil)
	if err != nil {
		return nil, nil, err
	}

	ids := new(str.TVExternalIDs)
	resp, err := s.client.Do(ctx, req, ids)
	if err != nil {
		return nil, resp, err
	}

	return ids, resp, nil
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

// GetContentRatings fetches the per-country content ratings for a single
// TV series.
//
// Api docs: https://developer.themoviedb.org/reference/tv-series-content-ratings
func (s *TVService) GetContentRatings(ctx context.Context, seriesID int64) (*str.TVContentRatings, *str.Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("tv/%d/content_ratings", seriesID), nil)
	if err != nil {
		return nil, nil, err
	}

	ratings := new(str.TVContentRatings)
	resp, err := s.client.Do(ctx, req, ratings)
	if err != nil {
		return nil, resp, err
	}

	return ratings, resp, nil
}

// GetCredits fetches the cast and crew for a single TV series.
//
// Api docs: https://developer.themoviedb.org/reference/tv-series-credits
func (s *TVService) GetCredits(ctx context.Context, seriesID int64, language string) (*str.TVCredits, *str.Response, error) {
	urlStr, err := uri.AddQuery(fmt.Sprintf("tv/%d/credits", seriesID), &uri.LanguageOptions{Language: language})
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(http.MethodGet, urlStr, nil)
	if err != nil {
		return nil, nil, err
	}

	credits := new(str.TVCredits)
	resp, err := s.client.Do(ctx, req, credits)
	if err != nil {
		return nil, resp, err
	}

	return credits, resp, nil
}

// GetImages fetches the images for a single TV series.
//
// Api docs: https://developer.themoviedb.org/reference/tv-series-images
func (s *TVService) GetImages(ctx context.Context, seriesID int64, opts *uri.ImagesOptions) (*str.Images, *str.Response, error) {
	urlStr, err := uri.AddQuery(fmt.Sprintf("tv/%d/images", seriesID), opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(http.MethodGet, urlStr, nil)
	if err != nil {
		return nil, nil, err
	}

	images := new(str.Images)
	resp, err := s.client.Do(ctx, req, images)
	if err != nil {
		return nil, resp, err
	}

	return images, resp, nil
}

// GetKeywords fetches the keywords for a single TV series.
//
// Api docs: https://developer.themoviedb.org/reference/tv-series-keywords
func (s *TVService) GetKeywords(ctx context.Context, seriesID int64) (*str.TVKeywords, *str.Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("tv/%d/keywords", seriesID), nil)
	if err != nil {
		return nil, nil, err
	}

	keywords := new(str.TVKeywords)
	resp, err := s.client.Do(ctx, req, keywords)
	if err != nil {
		return nil, resp, err
	}

	return keywords, resp, nil
}

// GetLatest fetches the most recently created TV series on TMDB.
//
// Api docs: https://developer.themoviedb.org/reference/tv-series-latest-id
func (s *TVService) GetLatest(ctx context.Context) (*str.TV, *str.Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, "tv/latest", nil)
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
