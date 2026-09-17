package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
)

// MoviesListsHandler handles `movies -a lists -i <movie_id>`.
type MoviesListsHandler struct {
	MovieID    int64
	Language   string
	PagesLimit int
}

// Handle fetches the lists a movie belongs to.
func (h MoviesListsHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	lists, err := client.Movies.GetLists(ctx, h.MovieID, h.Language, h.PagesLimit)
	if err != nil {
		return nil, err
	}
	return lists, nil
}
