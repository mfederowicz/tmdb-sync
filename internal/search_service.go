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
