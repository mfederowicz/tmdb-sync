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
