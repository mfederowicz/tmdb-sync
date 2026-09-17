package internal

import (
	"context"
	"fmt"
	"net/http"

	"github.com/mfederowicz/tmdb-sync/str"
)

// KeywordsService handles communication with the /keyword endpoints of the TMDB API.
type KeywordsService Service

// GetKeyword fetches details for a single keyword by TMDB keyword id.
//
// Api docs: https://developer.themoviedb.org/reference/keyword-details
func (s *KeywordsService) GetKeyword(ctx context.Context, keywordID string) (*str.Keyword, *str.Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("keyword/%s", keywordID), nil)
	if err != nil {
		return nil, nil, err
	}

	keyword := new(str.Keyword)
	resp, err := s.client.Do(ctx, req, keyword)
	if err != nil {
		return nil, resp, err
	}

	return keyword, resp, nil
}
