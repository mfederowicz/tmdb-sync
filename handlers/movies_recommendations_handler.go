package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
)

// MoviesRecommendationsHandler handles `movies -a recommendations -i <movie_id>`.
type MoviesRecommendationsHandler struct {
	MovieID    int64
	Language   string
	PagesLimit int
}

// Handle fetches movies recommended off a single movie.
func (h MoviesRecommendationsHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	movies, err := client.Movies.GetRecommendations(ctx, h.MovieID, h.Language, h.PagesLimit)
	if err != nil {
		return nil, err
	}
	return movies, nil
}
