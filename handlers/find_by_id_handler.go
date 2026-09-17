package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
	"github.com/mfederowicz/tmdb-sync/uri"
)

// FindByIDHandler handles `find -a by-id -i <external_id> --source <source>`.
type FindByIDHandler struct {
	ExternalID string
	Options    uri.FindOptions
}

// Handle looks up ExternalID and returns matching TMDB results.
func (h FindByIDHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	results, _, err := client.Find.Find(ctx, h.ExternalID, &h.Options)
	if err != nil {
		return nil, err
	}
	return results, nil
}
