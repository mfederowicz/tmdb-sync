package internal

import (
	"context"
	"fmt"
	"net/http"

	"github.com/mfederowicz/tmdb-sync/str"
	"github.com/mfederowicz/tmdb-sync/uri"
)

// GuestSessionsService handles communication with the /guest_session
// endpoints of the TMDB API.
type GuestSessionsService Service

func (s *GuestSessionsService) getRatedMoviesPage(ctx context.Context, guestSessionID, language, sortBy string, page int) (*str.RatedMovies, *str.Response, error) {
	urlStr, err := uri.AddQuery(fmt.Sprintf("guest_session/%s/rated/movies", guestSessionID), &uri.GuestSessionListOptions{Language: language, SortBy: sortBy, Page: page})
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(http.MethodGet, urlStr, nil)
	if err != nil {
		return nil, nil, err
	}

	movies := new(str.RatedMovies)
	resp, err := s.client.Do(ctx, req, movies)
	if err != nil {
		return nil, resp, err
	}

	return movies, resp, nil
}

// GetRatedMovies returns a guest session's rated movies, walking pages until
// TMDB reports no more (total_pages) or pagesLimit is reached (0 = unlimited).
//
// Api docs: https://developer.themoviedb.org/reference/guest-session-rated-movies
func (s *GuestSessionsService) GetRatedMovies(ctx context.Context, guestSessionID, language, sortBy string, pagesLimit int) ([]str.RatedMovie, error) {
	return FetchAllPages(ctx, pagesLimit, func(ctx context.Context, page int) (PageResult[str.RatedMovie], error) {
		movies, _, err := s.getRatedMoviesPage(ctx, guestSessionID, language, sortBy, page)
		if err != nil {
			return PageResult[str.RatedMovie]{}, err
		}
		return PageResult[str.RatedMovie]{
			Results:    movies.Results,
			Page:       movies.Page,
			TotalPages: movies.TotalPages,
		}, nil
	})
}

func (s *GuestSessionsService) getRatedTVPage(ctx context.Context, guestSessionID, language, sortBy string, page int) (*str.RatedTVShows, *str.Response, error) {
	urlStr, err := uri.AddQuery(fmt.Sprintf("guest_session/%s/rated/tv", guestSessionID), &uri.GuestSessionListOptions{Language: language, SortBy: sortBy, Page: page})
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(http.MethodGet, urlStr, nil)
	if err != nil {
		return nil, nil, err
	}

	tvShows := new(str.RatedTVShows)
	resp, err := s.client.Do(ctx, req, tvShows)
	if err != nil {
		return nil, resp, err
	}

	return tvShows, resp, nil
}

// GetRatedTV returns a guest session's rated TV shows, walking pages until
// TMDB reports no more (total_pages) or pagesLimit is reached (0 = unlimited).
//
// Api docs: https://developer.themoviedb.org/reference/guest-session-rated-tv
func (s *GuestSessionsService) GetRatedTV(ctx context.Context, guestSessionID, language, sortBy string, pagesLimit int) ([]str.RatedTV, error) {
	return FetchAllPages(ctx, pagesLimit, func(ctx context.Context, page int) (PageResult[str.RatedTV], error) {
		tvShows, _, err := s.getRatedTVPage(ctx, guestSessionID, language, sortBy, page)
		if err != nil {
			return PageResult[str.RatedTV]{}, err
		}
		return PageResult[str.RatedTV]{
			Results:    tvShows.Results,
			Page:       tvShows.Page,
			TotalPages: tvShows.TotalPages,
		}, nil
	})
}

func (s *GuestSessionsService) getRatedTVEpisodesPage(ctx context.Context, guestSessionID, language, sortBy string, page int) (*str.RatedTVEpisodes, *str.Response, error) {
	urlStr, err := uri.AddQuery(fmt.Sprintf("guest_session/%s/rated/tv/episodes", guestSessionID), &uri.GuestSessionListOptions{Language: language, SortBy: sortBy, Page: page})
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(http.MethodGet, urlStr, nil)
	if err != nil {
		return nil, nil, err
	}

	episodes := new(str.RatedTVEpisodes)
	resp, err := s.client.Do(ctx, req, episodes)
	if err != nil {
		return nil, resp, err
	}

	return episodes, resp, nil
}

// GetRatedTVEpisodes returns a guest session's rated TV episodes, walking
// pages until TMDB reports no more (total_pages) or pagesLimit is reached
// (0 = unlimited).
//
// Api docs: https://developer.themoviedb.org/reference/guest-session-rated-tv-episodes
func (s *GuestSessionsService) GetRatedTVEpisodes(ctx context.Context, guestSessionID, language, sortBy string, pagesLimit int) ([]str.RatedTVEpisode, error) {
	return FetchAllPages(ctx, pagesLimit, func(ctx context.Context, page int) (PageResult[str.RatedTVEpisode], error) {
		episodes, _, err := s.getRatedTVEpisodesPage(ctx, guestSessionID, language, sortBy, page)
		if err != nil {
			return PageResult[str.RatedTVEpisode]{}, err
		}
		return PageResult[str.RatedTVEpisode]{
			Results:    episodes.Results,
			Page:       episodes.Page,
			TotalPages: episodes.TotalPages,
		}, nil
	})
}
