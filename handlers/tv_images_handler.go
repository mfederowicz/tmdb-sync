package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
	"github.com/mfederowicz/tmdb-sync/uri"
)

// TVImagesHandler handles `tv -a images -i <series_id>`.
type TVImagesHandler struct {
	SeriesID             int64
	Language             string
	IncludeImageLanguage string
}

// Handle fetches the images for a TV series.
func (h TVImagesHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	images, _, err := client.TV.GetImages(ctx, h.SeriesID, &uri.ImagesOptions{
		Language:             h.Language,
		IncludeImageLanguage: h.IncludeImageLanguage,
	})
	if err != nil {
		return nil, err
	}
	return images, nil
}
