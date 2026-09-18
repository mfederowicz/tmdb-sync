package internal

import (
	"context"
	"fmt"
	"net/http"

	"github.com/mfederowicz/tmdb-sync/str"
	"github.com/mfederowicz/tmdb-sync/uri"
)

// TVEpisodesService handles communication with the
// /tv/{series_id}/season/{season_number}/episode/{episode_number} endpoints
// of the TMDB API.
type TVEpisodesService Service

// GetTVEpisode fetches details for a single TV episode by series id, season
// number, and episode number.
//
// Api docs: https://developer.themoviedb.org/reference/tv-episode-details
func (s *TVEpisodesService) GetTVEpisode(ctx context.Context, seriesID int64, seasonNumber int, episodeNumber int) (*str.TVEpisode, *str.Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("tv/%d/season/%d/episode/%d", seriesID, seasonNumber, episodeNumber), nil)
	if err != nil {
		return nil, nil, err
	}

	episode := new(str.TVEpisode)
	resp, err := s.client.Do(ctx, req, episode)
	if err != nil {
		return nil, resp, err
	}

	return episode, resp, nil
}

// GetAccountStates fetches an account's rated status for a single TV
// episode.
//
// Api docs: https://developer.themoviedb.org/reference/tv-episode-account-states
func (s *TVEpisodesService) GetAccountStates(ctx context.Context, seriesID int64, seasonNumber int, episodeNumber int, sessionID string) (*str.TVEpisodeAccountStates, *str.Response, error) {
	urlStr, err := uri.AddQuery(fmt.Sprintf("tv/%d/season/%d/episode/%d/account_states", seriesID, seasonNumber, episodeNumber), &uri.AccountOptions{SessionID: sessionID})
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(http.MethodGet, urlStr, nil)
	if err != nil {
		return nil, nil, err
	}

	states := new(str.TVEpisodeAccountStates)
	resp, err := s.client.Do(ctx, req, states)
	if err != nil {
		return nil, resp, err
	}

	return states, resp, nil
}

// GetCredits fetches the cast and crew for a single TV episode.
//
// Api docs: https://developer.themoviedb.org/reference/tv-episode-credits
func (s *TVEpisodesService) GetCredits(ctx context.Context, seriesID int64, seasonNumber int, episodeNumber int, language string) (*str.TVCredits, *str.Response, error) {
	urlStr, err := uri.AddQuery(fmt.Sprintf("tv/%d/season/%d/episode/%d/credits", seriesID, seasonNumber, episodeNumber), &uri.LanguageOptions{Language: language})
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

// GetExternalIDs fetches the external ids for a single TV episode.
//
// Api docs: https://developer.themoviedb.org/reference/tv-episode-external-ids
func (s *TVEpisodesService) GetExternalIDs(ctx context.Context, seriesID int64, seasonNumber int, episodeNumber int) (*str.TVEpisodeExternalIDs, *str.Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("tv/%d/season/%d/episode/%d/external_ids", seriesID, seasonNumber, episodeNumber), nil)
	if err != nil {
		return nil, nil, err
	}

	ids := new(str.TVEpisodeExternalIDs)
	resp, err := s.client.Do(ctx, req, ids)
	if err != nil {
		return nil, resp, err
	}

	return ids, resp, nil
}

// GetImages fetches the images for a single TV episode.
//
// Api docs: https://developer.themoviedb.org/reference/tv-episode-images
func (s *TVEpisodesService) GetImages(ctx context.Context, seriesID int64, seasonNumber int, episodeNumber int, opts *uri.ImagesOptions) (*str.TVEpisodeImages, *str.Response, error) {
	urlStr, err := uri.AddQuery(fmt.Sprintf("tv/%d/season/%d/episode/%d/images", seriesID, seasonNumber, episodeNumber), opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(http.MethodGet, urlStr, nil)
	if err != nil {
		return nil, nil, err
	}

	images := new(str.TVEpisodeImages)
	resp, err := s.client.Do(ctx, req, images)
	if err != nil {
		return nil, resp, err
	}

	return images, resp, nil
}

// GetTranslations fetches the translations for a single TV episode.
//
// Api docs: https://developer.themoviedb.org/reference/tv-episode-translations
func (s *TVEpisodesService) GetTranslations(ctx context.Context, seriesID int64, seasonNumber int, episodeNumber int) (*str.Translations, *str.Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("tv/%d/season/%d/episode/%d/translations", seriesID, seasonNumber, episodeNumber), nil)
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
// episode.
//
// Api docs: https://developer.themoviedb.org/reference/tv-episode-videos
func (s *TVEpisodesService) GetVideos(ctx context.Context, seriesID int64, seasonNumber int, episodeNumber int, language string) (*str.TVVideos, *str.Response, error) {
	urlStr, err := uri.AddQuery(fmt.Sprintf("tv/%d/season/%d/episode/%d/videos", seriesID, seasonNumber, episodeNumber), &uri.LanguageOptions{Language: language})
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

// AddRating rates a TV episode on behalf of the session's account, or a
// guest session when guestSessionID is set instead of sessionID.
//
// Api docs: https://developer.themoviedb.org/reference/tv-episode-add-rating
func (s *TVEpisodesService) AddRating(ctx context.Context, seriesID int64, seasonNumber int, episodeNumber int, sessionID, guestSessionID string, value float64) (*str.AuthStatus, *str.Response, error) {
	urlStr, err := uri.AddQuery(fmt.Sprintf("tv/%d/season/%d/episode/%d/rating", seriesID, seasonNumber, episodeNumber), &uri.RatingOptions{SessionID: sessionID, GuestSessionID: guestSessionID})
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(http.MethodPost, urlStr, &str.TVRatingRequest{Value: value})
	if err != nil {
		return nil, nil, err
	}

	status := new(str.AuthStatus)
	resp, err := s.client.Do(ctx, req, status)
	if err != nil {
		return nil, resp, err
	}

	return status, resp, nil
}

// DeleteRating removes the session's account's rating for a TV episode.
//
// Api docs: https://developer.themoviedb.org/reference/tv-episode-delete-rating
func (s *TVEpisodesService) DeleteRating(ctx context.Context, seriesID int64, seasonNumber int, episodeNumber int, sessionID string) (*str.AuthStatus, *str.Response, error) {
	urlStr, err := uri.AddQuery(fmt.Sprintf("tv/%d/season/%d/episode/%d/rating", seriesID, seasonNumber, episodeNumber), &uri.AccountOptions{SessionID: sessionID})
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(http.MethodDelete, urlStr, nil)
	if err != nil {
		return nil, nil, err
	}

	status := new(str.AuthStatus)
	resp, err := s.client.Do(ctx, req, status)
	if err != nil {
		return nil, resp, err
	}

	return status, resp, nil
}
