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

// GetCredits fetches the cast and crew for a single TV season.
//
// Api docs: https://developer.themoviedb.org/reference/tv-season-credits
func (s *TVSeasonsService) GetCredits(ctx context.Context, seriesID int64, seasonNumber int, language string) (*str.TVCredits, *str.Response, error) {
	urlStr, err := uri.AddQuery(fmt.Sprintf("tv/%d/season/%d/credits", seriesID, seasonNumber), &uri.LanguageOptions{Language: language})
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

// GetExternalIDs fetches the external ids for a single TV season.
//
// Api docs: https://developer.themoviedb.org/reference/tv-season-external-ids
func (s *TVSeasonsService) GetExternalIDs(ctx context.Context, seriesID int64, seasonNumber int) (*str.TVSeasonExternalIDs, *str.Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("tv/%d/season/%d/external_ids", seriesID, seasonNumber), nil)
	if err != nil {
		return nil, nil, err
	}

	ids := new(str.TVSeasonExternalIDs)
	resp, err := s.client.Do(ctx, req, ids)
	if err != nil {
		return nil, resp, err
	}

	return ids, resp, nil
}

// GetImages fetches the images for a single TV season.
//
// Api docs: https://developer.themoviedb.org/reference/tv-season-images
func (s *TVSeasonsService) GetImages(ctx context.Context, seriesID int64, seasonNumber int, opts *uri.ImagesOptions) (*str.Images, *str.Response, error) {
	urlStr, err := uri.AddQuery(fmt.Sprintf("tv/%d/season/%d/images", seriesID, seasonNumber), opts)
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

// GetTranslations fetches the translations for a single TV season.
//
// Api docs: https://developer.themoviedb.org/reference/tv-season-translations
func (s *TVSeasonsService) GetTranslations(ctx context.Context, seriesID int64, seasonNumber int) (*str.Translations, *str.Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("tv/%d/season/%d/translations", seriesID, seasonNumber), nil)
	if err != nil {
		return nil, nil, err
	}

	translations := new(str.Translations)
	resp, err := s.client.Do(ctx, req, translations)
	if err != nil {
		return nil, resp, err
	}

	return translations, resp, nil
}

// GetVideos fetches the videos (trailers, teasers, ...) for a single TV
// season.
//
// Api docs: https://developer.themoviedb.org/reference/tv-season-videos
func (s *TVSeasonsService) GetVideos(ctx context.Context, seriesID int64, seasonNumber int, language string) (*str.TVVideos, *str.Response, error) {
	urlStr, err := uri.AddQuery(fmt.Sprintf("tv/%d/season/%d/videos", seriesID, seasonNumber), &uri.LanguageOptions{Language: language})
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(http.MethodGet, urlStr, nil)
	if err != nil {
		return nil, nil, err
	}

	videos := new(str.TVVideos)
	resp, err := s.client.Do(ctx, req, videos)
	if err != nil {
		return nil, resp, err
	}

	return videos, resp, nil
}
