package internal

import (
	"context"
	"fmt"
	"net/http"

	"github.com/mfederowicz/tmdb-sync/str"
	"github.com/mfederowicz/tmdb-sync/uri"
)

// FindService handles communication with the /find endpoint of the TMDB API.
type FindService Service

// Find looks up an external id (IMDb, TVDB, ...) and returns matching TMDB
// movie/tv/person/episode/season results.
//
// Api docs: https://developer.themoviedb.org/reference/find-by-id
func (s *FindService) Find(ctx context.Context, externalID string, opts *uri.FindOptions) (*str.FindResults, *str.Response, error) {
	urlStr, err := uri.AddQuery(fmt.Sprintf("find/%s", externalID), opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(http.MethodGet, urlStr, nil)
	if err != nil {
		return nil, nil, err
	}

	results := new(str.FindResults)
	resp, err := s.client.Do(ctx, req, results)
	if err != nil {
		return nil, resp, err
	}

	return results, resp, nil
}
