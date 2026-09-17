package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
	"github.com/mfederowicz/tmdb-sync/uri"
)

// SearchTVHandler handles `search -a tv -q <query>`.
type SearchTVHandler struct {
	Query            string
	Language         string
	IncludeAdult     bool
	FirstAirDateYear int
	Year             int
	PagesLimit       int
}

// Handle searches for TV shows matching the query.
func (h SearchTVHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	opts := uri.SearchTVOptions{
		Query:            h.Query,
		Language:         h.Language,
		IncludeAdult:     h.IncludeAdult,
		FirstAirDateYear: h.FirstAirDateYear,
		Year:             h.Year,
	}
	results, err := client.Search.SearchTV(ctx, opts, h.PagesLimit)
	if err != nil {
		return nil, err
	}
	return results, nil
}
