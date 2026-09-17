package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
)

// MoviesTranslationsHandler handles `movies -a translations -i <movie_id>`.
type MoviesTranslationsHandler struct {
	MovieID int64
}

// Handle fetches the translated fields for a movie.
func (h MoviesTranslationsHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	translations, _, err := client.Movies.GetTranslations(ctx, h.MovieID)
	if err != nil {
		return nil, err
	}
	return translations, nil
}
