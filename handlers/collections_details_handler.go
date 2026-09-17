package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
)

// CollectionsDetailsHandler handles `collections -a details -i <collection_id>`.
type CollectionsDetailsHandler struct {
	CollectionID int64
}

// Handle fetches details for a single collection.
func (h CollectionsDetailsHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	collection, _, err := client.Collections.GetCollection(ctx, h.CollectionID)
	if err != nil {
		return nil, err
	}
	return collection, nil
}
