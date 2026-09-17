package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
)

// MoviesSimilarHandler handles `movies -a similar -i <movie_id>`.
type MoviesSimilarHandler struct {
	MovieID    int64
	Language   string
	PagesLimit int
}

// Handle fetches movies similar to a single movie.
func (h MoviesSimilarHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	movies, err := client.Movies.GetSimilar(ctx, h.MovieID, h.Language, h.PagesLimit)
	if err != nil {
		return nil, err
	}
	return movies, nil
}
