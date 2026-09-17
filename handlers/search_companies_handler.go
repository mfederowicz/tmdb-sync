package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
	"github.com/mfederowicz/tmdb-sync/uri"
)

// SearchCompaniesHandler handles `search -a companies -q <query>`.
type SearchCompaniesHandler struct {
	Query      string
	PagesLimit int
}

// Handle searches for companies matching the query.
func (h SearchCompaniesHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	opts := uri.SearchCompanyOptions{Query: h.Query}
	results, err := client.Search.SearchCompanies(ctx, opts, h.PagesLimit)
	if err != nil {
		return nil, err
	}
	return results, nil
}
