package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
	"github.com/mfederowicz/tmdb-sync/uri"
)

// SearchMultiHandler handles `search -a multi -q <query>`.
type SearchMultiHandler struct {
	Query        string
	Language     string
	IncludeAdult bool
	PagesLimit   int
}

// Handle searches movies, tv shows, and people matching the query.
func (h SearchMultiHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	opts := uri.SearchMultiOptions{
		Query:        h.Query,
		Language:     h.Language,
		IncludeAdult: h.IncludeAdult,
	}
	results, err := client.Search.SearchMulti(ctx, opts, h.PagesLimit)
	if err != nil {
		return nil, err
	}
	return results, nil
}
