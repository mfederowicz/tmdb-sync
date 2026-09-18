package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
)

// TVEpisodesTranslationsHandler handles `tv-episodes -a translations -i <series_id> -s <season_number> -e <episode_number>`.
type TVEpisodesTranslationsHandler struct {
	SeriesID      int64
	SeasonNumber  int
	EpisodeNumber int
}

// Handle fetches the translations for a single TV episode.
func (h TVEpisodesTranslationsHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	translations, _, err := client.TVEpisodes.GetTranslations(ctx, h.SeriesID, h.SeasonNumber, h.EpisodeNumber)
	if err != nil {
		return nil, err
	}
	return translations, nil
}
