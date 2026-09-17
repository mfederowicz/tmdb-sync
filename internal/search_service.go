package internal

import (
	"context"
	"net/http"

	"github.com/mfederowicz/tmdb-sync/str"
	"github.com/mfederowicz/tmdb-sync/uri"
)

// SearchService handles communication with the /search endpoints of the TMDB API.
type SearchService Service

// getSearchCollectionsPage fetches a single page of the search/collection list.
func (s *SearchService) getSearchCollectionsPage(ctx context.Context, opts *uri.SearchCollectionOptions) (*str.SearchCollections, *str.Response, error) {
	urlStr, err := uri.AddQuery("search/collection", opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(http.MethodGet, urlStr, nil)
	if err != nil {
		return nil, nil, err
	}

	collections := new(str.SearchCollections)
	resp, err := s.client.Do(ctx, req, collections)
	if err != nil {
		return nil, resp, err
	}

	return collections, resp, nil
}

// SearchCollections returns collections matching opts, walking pages until
// TMDB reports no more (total_pages) or pagesLimit is reached (0 = unlimited).
//
// Api docs: https://developer.themoviedb.org/reference/search-collection
func (s *SearchService) SearchCollections(ctx context.Context, opts uri.SearchCollectionOptions, pagesLimit int) ([]str.SearchCollectionResult, error) {
	return FetchAllPages(ctx, pagesLimit, func(ctx context.Context, page int) (PageResult[str.SearchCollectionResult], error) {
		o := opts
		o.Page = page
		collections, _, err := s.getSearchCollectionsPage(ctx, &o)
		if err != nil {
			return PageResult[str.SearchCollectionResult]{}, err
		}
		return PageResult[str.SearchCollectionResult]{
			Results:    collections.Results,
			Page:       collections.Page,
			TotalPages: collections.TotalPages,
		}, nil
	})
}

// getSearchCompaniesPage fetches a single page of the search/company list.
func (s *SearchService) getSearchCompaniesPage(ctx context.Context, opts *uri.SearchCompanyOptions) (*str.SearchCompanies, *str.Response, error) {
	urlStr, err := uri.AddQuery("search/company", opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(http.MethodGet, urlStr, nil)
	if err != nil {
		return nil, nil, err
	}

	companies := new(str.SearchCompanies)
	resp, err := s.client.Do(ctx, req, companies)
	if err != nil {
		return nil, resp, err
	}

	return companies, resp, nil
}

// SearchCompanies returns companies matching opts, walking pages until
// TMDB reports no more (total_pages) or pagesLimit is reached (0 = unlimited).
//
// Api docs: https://developer.themoviedb.org/reference/search-company
func (s *SearchService) SearchCompanies(ctx context.Context, opts uri.SearchCompanyOptions, pagesLimit int) ([]str.SearchCompanyResult, error) {
	return FetchAllPages(ctx, pagesLimit, func(ctx context.Context, page int) (PageResult[str.SearchCompanyResult], error) {
		o := opts
		o.Page = page
		companies, _, err := s.getSearchCompaniesPage(ctx, &o)
		if err != nil {
			return PageResult[str.SearchCompanyResult]{}, err
		}
		return PageResult[str.SearchCompanyResult]{
			Results:    companies.Results,
			Page:       companies.Page,
			TotalPages: companies.TotalPages,
		}, nil
	})
}

// getSearchKeywordsPage fetches a single page of the search/keyword list.
func (s *SearchService) getSearchKeywordsPage(ctx context.Context, opts *uri.SearchKeywordOptions) (*str.SearchKeywords, *str.Response, error) {
	urlStr, err := uri.AddQuery("search/keyword", opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(http.MethodGet, urlStr, nil)
	if err != nil {
		return nil, nil, err
	}

	keywords := new(str.SearchKeywords)
	resp, err := s.client.Do(ctx, req, keywords)
	if err != nil {
		return nil, resp, err
	}

	return keywords, resp, nil
}

// SearchKeywords returns keywords matching opts, walking pages until
// TMDB reports no more (total_pages) or pagesLimit is reached (0 = unlimited).
//
// Api docs: https://developer.themoviedb.org/reference/search-keyword
func (s *SearchService) SearchKeywords(ctx context.Context, opts uri.SearchKeywordOptions, pagesLimit int) ([]str.SearchKeywordResult, error) {
	return FetchAllPages(ctx, pagesLimit, func(ctx context.Context, page int) (PageResult[str.SearchKeywordResult], error) {
		o := opts
		o.Page = page
		keywords, _, err := s.getSearchKeywordsPage(ctx, &o)
		if err != nil {
			return PageResult[str.SearchKeywordResult]{}, err
		}
		return PageResult[str.SearchKeywordResult]{
			Results:    keywords.Results,
			Page:       keywords.Page,
			TotalPages: keywords.TotalPages,
		}, nil
	})
}

// getSearchMoviesPage fetches a single page of the search/movie list.
func (s *SearchService) getSearchMoviesPage(ctx context.Context, opts *uri.SearchMovieOptions) (*str.Movies, *str.Response, error) {
	urlStr, err := uri.AddQuery("search/movie", opts)
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

// SearchMovies returns movies matching opts, walking pages until TMDB
// reports no more (total_pages) or pagesLimit is reached (0 = unlimited).
//
// Api docs: https://developer.themoviedb.org/reference/search-movie
func (s *SearchService) SearchMovies(ctx context.Context, opts uri.SearchMovieOptions, pagesLimit int) ([]str.Movie, error) {
	return FetchAllPages(ctx, pagesLimit, func(ctx context.Context, page int) (PageResult[str.Movie], error) {
		o := opts
		o.Page = page
		movies, _, err := s.getSearchMoviesPage(ctx, &o)
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
