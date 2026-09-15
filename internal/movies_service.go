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

// getPopularMoviesPage fetches a single page of the popular-movies list.
func (s *MoviesService) getPopularMoviesPage(ctx context.Context, page int) (*str.Movies, *str.Response, error) {
	urlStr, err := uri.AddPage("movie/popular", page)
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
func (s *MoviesService) GetPopularMovies(ctx context.Context, pagesLimit int) ([]str.Movie, error) {
	return FetchAllPages(ctx, pagesLimit, func(ctx context.Context, page int) (PageResult[str.Movie], error) {
		movies, _, err := s.getPopularMoviesPage(ctx, page)
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
