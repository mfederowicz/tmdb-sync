package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
)

// MoviesAlternativeTitlesHandler handles `movies -a alternative-titles -i <movie_id>`.
type MoviesAlternativeTitlesHandler struct {
	MovieID int64
	Country string
}

// Handle fetches the alternative titles for a movie.
func (h MoviesAlternativeTitlesHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	titles, _, err := client.Movies.GetAlternativeTitles(ctx, h.MovieID, h.Country)
	if err != nil {
		return nil, err
	}
	return titles, nil
}
