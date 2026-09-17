package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
	"github.com/mfederowicz/tmdb-sync/uri"
)

// TVSeasonsImagesHandler handles `tv-seasons -a images -i <series_id> -s <season_number>`.
type TVSeasonsImagesHandler struct {
	SeriesID             int64
	SeasonNumber         int
	Language             string
	IncludeImageLanguage string
}

// Handle fetches the images for a TV season.
func (h TVSeasonsImagesHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	images, _, err := client.TVSeasons.GetImages(ctx, h.SeriesID, h.SeasonNumber, &uri.ImagesOptions{
		Language:             h.Language,
		IncludeImageLanguage: h.IncludeImageLanguage,
	})
	if err != nil {
		return nil, err
	}
	return images, nil
}
