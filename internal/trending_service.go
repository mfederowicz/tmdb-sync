package internal

import (
	"context"
	"fmt"
	"net/http"

	"github.com/mfederowicz/tmdb-sync/str"
	"github.com/mfederowicz/tmdb-sync/uri"
)

// TrendingService handles communication with the /trending endpoints of the TMDB API.
type TrendingService Service

// getTrendingAllPage fetches a single page of the trending/all list.
func (s *TrendingService) getTrendingAllPage(ctx context.Context, timeWindow string, opts *uri.TrendingOptions) (*str.Trending, *str.Response, error) {
	urlStr, err := uri.AddQuery(fmt.Sprintf("trending/all/%s", timeWindow), opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(http.MethodGet, urlStr, nil)
	if err != nil {
		return nil, nil, err
	}

	trending := new(str.Trending)
	resp, err := s.client.Do(ctx, req, trending)
	if err != nil {
		return nil, resp, err
	}

	return trending, resp, nil
}

// TrendingAll returns the day's or week's trending movies, tv shows, and
// people, walking pages until TMDB reports no more (total_pages) or
// pagesLimit is reached (0 = unlimited).
//
// Api docs: https://developer.themoviedb.org/reference/trending-all
func (s *TrendingService) TrendingAll(ctx context.Context, timeWindow string, pagesLimit int) ([]str.SearchMultiResult, error) {
	return FetchAllPages(ctx, pagesLimit, func(ctx context.Context, page int) (PageResult[str.SearchMultiResult], error) {
		trending, _, err := s.getTrendingAllPage(ctx, timeWindow, &uri.TrendingOptions{Page: page})
		if err != nil {
			return PageResult[str.SearchMultiResult]{}, err
		}
		return PageResult[str.SearchMultiResult]{
			Results:    trending.Results,
			Page:       trending.Page,
			TotalPages: trending.TotalPages,
		}, nil
	})
}

// getTrendingMoviesPage fetches a single page of the trending/movie list.
func (s *TrendingService) getTrendingMoviesPage(ctx context.Context, timeWindow string, opts *uri.TrendingOptions) (*str.Movies, *str.Response, error) {
	urlStr, err := uri.AddQuery(fmt.Sprintf("trending/movie/%s", timeWindow), opts)
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

// TrendingMovies returns the day's or week's trending movies, walking pages
// until TMDB reports no more (total_pages) or pagesLimit is reached
// (0 = unlimited).
//
// Api docs: https://developer.themoviedb.org/reference/trending-movies
func (s *TrendingService) TrendingMovies(ctx context.Context, timeWindow string, pagesLimit int) ([]str.Movie, error) {
	return FetchAllPages(ctx, pagesLimit, func(ctx context.Context, page int) (PageResult[str.Movie], error) {
		movies, _, err := s.getTrendingMoviesPage(ctx, timeWindow, &uri.TrendingOptions{Page: page})
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

// getTrendingTVPage fetches a single page of the trending/tv list.
func (s *TrendingService) getTrendingTVPage(ctx context.Context, timeWindow string, opts *uri.TrendingOptions) (*str.TVShows, *str.Response, error) {
	urlStr, err := uri.AddQuery(fmt.Sprintf("trending/tv/%s", timeWindow), opts)
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

// TrendingTV returns the day's or week's trending TV shows, walking pages
// until TMDB reports no more (total_pages) or pagesLimit is reached
// (0 = unlimited).
//
// Api docs: https://developer.themoviedb.org/reference/trending-tv
func (s *TrendingService) TrendingTV(ctx context.Context, timeWindow string, pagesLimit int) ([]str.TV, error) {
	return FetchAllPages(ctx, pagesLimit, func(ctx context.Context, page int) (PageResult[str.TV], error) {
		shows, _, err := s.getTrendingTVPage(ctx, timeWindow, &uri.TrendingOptions{Page: page})
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

// getTrendingPeoplePage fetches a single page of the trending/person list.
func (s *TrendingService) getTrendingPeoplePage(ctx context.Context, timeWindow string, opts *uri.TrendingOptions) (*str.PopularPersons, *str.Response, error) {
	urlStr, err := uri.AddQuery(fmt.Sprintf("trending/person/%s", timeWindow), opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(http.MethodGet, urlStr, nil)
	if err != nil {
		return nil, nil, err
	}

	people := new(str.PopularPersons)
	resp, err := s.client.Do(ctx, req, people)
	if err != nil {
		return nil, resp, err
	}

	return people, resp, nil
}

// TrendingPeople returns the day's or week's trending people, walking pages
// until TMDB reports no more (total_pages) or pagesLimit is reached
// (0 = unlimited).
//
// Api docs: https://developer.themoviedb.org/reference/trending-people
func (s *TrendingService) TrendingPeople(ctx context.Context, timeWindow string, pagesLimit int) ([]str.PopularPerson, error) {
	return FetchAllPages(ctx, pagesLimit, func(ctx context.Context, page int) (PageResult[str.PopularPerson], error) {
		people, _, err := s.getTrendingPeoplePage(ctx, timeWindow, &uri.TrendingOptions{Page: page})
		if err != nil {
			return PageResult[str.PopularPerson]{}, err
		}
		return PageResult[str.PopularPerson]{
			Results:    people.Results,
			Page:       people.Page,
			TotalPages: people.TotalPages,
		}, nil
	})
}
