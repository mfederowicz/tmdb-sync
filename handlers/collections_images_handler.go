package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
	"github.com/mfederowicz/tmdb-sync/uri"
)

// CollectionsImagesHandler handles `collections -a images -i <collection_id>`.
type CollectionsImagesHandler struct {
	CollectionID         int64
	Language             string
	IncludeImageLanguage string
}

// Handle fetches the images for a single collection.
func (h CollectionsImagesHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	images, _, err := client.Collections.GetCollectionImages(ctx, h.CollectionID, &uri.ImagesOptions{
		Language:             h.Language,
		IncludeImageLanguage: h.IncludeImageLanguage,
	})
	if err != nil {
		return nil, err
	}
	return images, nil
}
