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

// GetPopularMovies returns a page of the current popular-movies list.
func (s *MoviesService) GetPopularMovies(ctx context.Context, page int) (*str.Movies, *str.Response, error) {
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
