package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
	"github.com/mfederowicz/tmdb-sync/uri"
)

// SearchCollectionsHandler handles `search -a collections -q <query>`.
type SearchCollectionsHandler struct {
	Query        string
	Language     string
	IncludeAdult bool
	Region       string
	PagesLimit   int
}

// Handle searches for collections matching the query.
func (h SearchCollectionsHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	opts := uri.SearchCollectionOptions{
		Query:        h.Query,
		Language:     h.Language,
		IncludeAdult: h.IncludeAdult,
		Region:       h.Region,
	}
	results, err := client.Search.SearchCollections(ctx, opts, h.PagesLimit)
	if err != nil {
		return nil, err
	}
	return results, nil
}
