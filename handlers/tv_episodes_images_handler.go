package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
	"github.com/mfederowicz/tmdb-sync/uri"
)

// TVEpisodesImagesHandler handles `tv-episodes -a images -i <series_id> -s <season_number> -e <episode_number>`.
type TVEpisodesImagesHandler struct {
	SeriesID             int64
	SeasonNumber         int
	EpisodeNumber        int
	Language             string
	IncludeImageLanguage string
}

// Handle fetches the images for a single TV episode.
func (h TVEpisodesImagesHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	images, _, err := client.TVEpisodes.GetImages(ctx, h.SeriesID, h.SeasonNumber, h.EpisodeNumber, &uri.ImagesOptions{
		Language:             h.Language,
		IncludeImageLanguage: h.IncludeImageLanguage,
	})
	if err != nil {
		return nil, err
	}
	return images, nil
}
