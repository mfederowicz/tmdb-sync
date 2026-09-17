package internal

import (
	"context"
	"net/http"

	"github.com/mfederowicz/tmdb-sync/str"
)

// GenreService handles the TMDB /genre endpoints.
type GenreService Service

// GetMovieList returns the list of official genres for movies.
//
// Api docs: https://developer.themoviedb.org/reference/genre-movie-list
func (s *GenreService) GetMovieList(ctx context.Context) (*str.GenreList, *str.Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, "genre/movie/list", nil)
	if err != nil {
		return nil, nil, err
	}

	genres := new(str.GenreList)
	resp, err := s.client.Do(ctx, req, genres)
	if err != nil {
		return nil, resp, err
	}

	return genres, resp, nil
}

// GetTVList returns the list of official genres for TV shows.
//
// Api docs: https://developer.themoviedb.org/reference/genre-tv-list
func (s *GenreService) GetTVList(ctx context.Context) (*str.GenreList, *str.Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, "genre/tv/list", nil)
	if err != nil {
		return nil, nil, err
	}

	genres := new(str.GenreList)
	resp, err := s.client.Do(ctx, req, genres)
	if err != nil {
		return nil, resp, err
	}

	return genres, resp, nil
}
