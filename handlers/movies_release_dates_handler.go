package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
)

// MoviesReleaseDatesHandler handles `movies -a release-dates -i <movie_id>`.
type MoviesReleaseDatesHandler struct {
	MovieID int64
}

// Handle fetches the per-country release dates and certifications for a movie.
func (h MoviesReleaseDatesHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	dates, _, err := client.Movies.GetReleaseDates(ctx, h.MovieID)
	if err != nil {
		return nil, err
	}
	return dates, nil
}
