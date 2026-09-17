package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
)

// MoviesKeywordsHandler handles `movies -a keywords -i <movie_id>`.
type MoviesKeywordsHandler struct {
	MovieID int64
}

// Handle fetches the keywords for a movie.
func (h MoviesKeywordsHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	keywords, _, err := client.Movies.GetKeywords(ctx, h.MovieID)
	if err != nil {
		return nil, err
	}
	return keywords, nil
}
