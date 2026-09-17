package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
)

// CollectionsTranslationsHandler handles `collections -a translations -i <collection_id>`.
type CollectionsTranslationsHandler struct {
	CollectionID int64
}

// Handle fetches the translations for a single collection.
func (h CollectionsTranslationsHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	translations, _, err := client.Collections.GetCollectionTranslations(ctx, h.CollectionID)
	if err != nil {
		return nil, err
	}
	return translations, nil
}
