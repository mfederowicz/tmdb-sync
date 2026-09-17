package internal

import (
	"context"
	"net/http"

	"github.com/mfederowicz/tmdb-sync/str"
	"github.com/mfederowicz/tmdb-sync/uri"
)

// DiscoverService handles communication with the /discover endpoints of the TMDB API.
type DiscoverService Service

// getDiscoverMoviesPage fetches a single page of the discover-movie list.
func (s *DiscoverService) getDiscoverMoviesPage(ctx context.Context, opts *uri.DiscoverMovieOptions) (*str.Movies, *str.Response, error) {
	urlStr, err := uri.AddQuery("discover/movie", opts)
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

// GetDiscoverMovies returns movies matching opts, walking pages until TMDB
// reports no more (total_pages) or pagesLimit is reached (0 = unlimited).
//
// Api docs: https://developer.themoviedb.org/reference/discover-movie
func (s *DiscoverService) GetDiscoverMovies(ctx context.Context, opts uri.DiscoverMovieOptions, pagesLimit int) ([]str.Movie, error) {
	return FetchAllPages(ctx, pagesLimit, func(ctx context.Context, page int) (PageResult[str.Movie], error) {
		o := opts
		o.Page = page
		movies, _, err := s.getDiscoverMoviesPage(ctx, &o)
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

// getDiscoverTVPage fetches a single page of the discover-tv list.
func (s *DiscoverService) getDiscoverTVPage(ctx context.Context, opts *uri.DiscoverTVOptions) (*str.TVShows, *str.Response, error) {
	urlStr, err := uri.AddQuery("discover/tv", opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(http.MethodGet, urlStr, nil)
	if err != nil {
		return nil, nil, err
	}

	shows := new(str.TVShows)
	resp, err := s.client.Do(ctx, req, shows)
	if err != nil {
		return nil, resp, err
	}

	return shows, resp, nil
}

// GetDiscoverTV returns TV shows matching opts, walking pages until TMDB
// reports no more (total_pages) or pagesLimit is reached (0 = unlimited).
//
// Api docs: https://developer.themoviedb.org/reference/discover-tv
func (s *DiscoverService) GetDiscoverTV(ctx context.Context, opts uri.DiscoverTVOptions, pagesLimit int) ([]str.TV, error) {
	return FetchAllPages(ctx, pagesLimit, func(ctx context.Context, page int) (PageResult[str.TV], error) {
		o := opts
		o.Page = page
		shows, _, err := s.getDiscoverTVPage(ctx, &o)
		if err != nil {
			return PageResult[str.TV]{}, err
		}
		return PageResult[str.TV]{
			Results:    shows.Results,
			Page:       shows.Page,
			TotalPages: shows.TotalPages,
		}, nil
	})
}
