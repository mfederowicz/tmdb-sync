package internal

import (
	"context"
	"fmt"
	"net/http"

	"github.com/mfederowicz/tmdb-sync/str"
	"github.com/mfederowicz/tmdb-sync/uri"
)

// MoviesService handles communication with the /movie endpoints of the TMDB API.
type MoviesService Service

// GetMovie fetches details for a single movie by TMDB id.
//
// Api docs: https://developer.themoviedb.org/reference/movie-details
func (s *MoviesService) GetMovie(ctx context.Context, movieID int64) (*str.Movie, *str.Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("movie/%d", movieID), nil)
	if err != nil {
		return nil, nil, err
	}

	movie := new(str.Movie)
	resp, err := s.client.Do(ctx, req, movie)
	if err != nil {
		return nil, resp, err
	}

	return movie, resp, nil
}

// GetAccountStates fetches an account's favorite/rated/watchlist status for
// a single movie.
//
// Api docs: https://developer.themoviedb.org/reference/movie-account-states
func (s *MoviesService) GetAccountStates(ctx context.Context, movieID int64, sessionID string) (*str.MovieAccountStates, *str.Response, error) {
	urlStr, err := uri.AddQuery(fmt.Sprintf("movie/%d/account_states", movieID), &uri.AccountOptions{SessionID: sessionID})
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(http.MethodGet, urlStr, nil)
	if err != nil {
		return nil, nil, err
	}

	states := new(str.MovieAccountStates)
	resp, err := s.client.Do(ctx, req, states)
	if err != nil {
		return nil, resp, err
	}

	return states, resp, nil
}

// GetAlternativeTitles fetches the alternative titles for a single movie,
// optionally filtered to a single country (ISO 3166-1).
//
// Api docs: https://developer.themoviedb.org/reference/movie-alternative-titles
func (s *MoviesService) GetAlternativeTitles(ctx context.Context, movieID int64, country string) (*str.MovieAlternativeTitles, *str.Response, error) {
	urlStr, err := uri.AddQuery(fmt.Sprintf("movie/%d/alternative_titles", movieID), &uri.MovieAlternativeTitlesOptions{Country: country})
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(http.MethodGet, urlStr, nil)
	if err != nil {
		return nil, nil, err
	}

	titles := new(str.MovieAlternativeTitles)
	resp, err := s.client.Do(ctx, req, titles)
	if err != nil {
		return nil, resp, err
	}

	return titles, resp, nil
}

// GetCredits fetches the cast and crew for a single movie.
//
// Api docs: https://developer.themoviedb.org/reference/movie-credits
func (s *MoviesService) GetCredits(ctx context.Context, movieID int64, language string) (*str.MovieCredits, *str.Response, error) {
	urlStr, err := uri.AddQuery(fmt.Sprintf("movie/%d/credits", movieID), &uri.LanguageOptions{Language: language})
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(http.MethodGet, urlStr, nil)
	if err != nil {
		return nil, nil, err
	}

	credits := new(str.MovieCredits)
	resp, err := s.client.Do(ctx, req, credits)
	if err != nil {
		return nil, resp, err
	}

	return credits, resp, nil
}

// GetExternalIDs fetches a single movie's external ids (IMDb, Wikidata, ...).
//
// Api docs: https://developer.themoviedb.org/reference/movie-external-ids
func (s *MoviesService) GetExternalIDs(ctx context.Context, movieID int64) (*str.MovieExternalIDs, *str.Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("movie/%d/external_ids", movieID), nil)
	if err != nil {
		return nil, nil, err
	}

	ids := new(str.MovieExternalIDs)
	resp, err := s.client.Do(ctx, req, ids)
	if err != nil {
		return nil, resp, err
	}

	return ids, resp, nil
}

// getPopularMoviesPage fetches a single page of the popular-movies list.
func (s *MoviesService) getPopularMoviesPage(ctx context.Context, opts *uri.ListOptions) (*str.Movies, *str.Response, error) {
	urlStr, err := uri.AddQuery("movie/popular", opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(http.MethodGet, urlStr, nil)
	if err != nil {
		return nil, nil, err
	}

	movies := new(str.Movies)
	resp, err := s.client.Do(ctx, req, movies)
	if err != nil {
		return nil, resp, err
	}

	return movies, resp, nil
}

// GetPopularMovies returns the current popular-movies list, walking pages
// until TMDB reports no more (total_pages) or pagesLimit is reached
// (0 = unlimited). TMDB's page size is fixed at 20 by the API; pagesLimit is
// the only real lever over how much gets fetched.
//
// Api docs: https://developer.themoviedb.org/reference/movie-popular-list
func (s *MoviesService) GetPopularMovies(ctx context.Context, pagesLimit int) ([]str.Movie, error) {
	return FetchAllPages(ctx, pagesLimit, func(ctx context.Context, page int) (PageResult[str.Movie], error) {
		movies, _, err := s.getPopularMoviesPage(ctx, &uri.ListOptions{Page: page})
		if err != nil {
			return PageResult[str.Movie]{}, err
		}
		return PageResult[str.Movie]{
			Results:    movies.Results,
			Page:       movies.Page,
			TotalPages: movies.TotalPages,
		}, nil
	})
}
