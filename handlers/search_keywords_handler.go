package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
	"github.com/mfederowicz/tmdb-sync/uri"
)

// SearchKeywordsHandler handles `search -a keywords -q <query>`.
type SearchKeywordsHandler struct {
	Query      string
	PagesLimit int
}

// Handle searches for keywords matching the query.
func (h SearchKeywordsHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	opts := uri.SearchKeywordOptions{Query: h.Query}
	results, err := client.Search.SearchKeywords(ctx, opts, h.PagesLimit)
	if err != nil {
		return nil, err
	}
	return results, nil
}
