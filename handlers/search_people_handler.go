package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
	"github.com/mfederowicz/tmdb-sync/uri"
)

// SearchPeopleHandler handles `search -a people -q <query>`.
type SearchPeopleHandler struct {
	Query        string
	Language     string
	IncludeAdult bool
	PagesLimit   int
}

// Handle searches for people matching the query.
func (h SearchPeopleHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	opts := uri.SearchPersonOptions{
		Query:        h.Query,
		Language:     h.Language,
		IncludeAdult: h.IncludeAdult,
	}
	results, err := client.Search.SearchPeople(ctx, opts, h.PagesLimit)
	if err != nil {
		return nil, err
	}
	return results, nil
}
