package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
)

// ListsAddMovieHandler handles `lists -a add-movie -i <list_id> -media-id <movie_id>`.
type ListsAddMovieHandler struct {
	ListID    string
	SessionID string
	MovieID   int64
}

// Handle adds a movie to a list.
func (h ListsAddMovieHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	status, _, err := client.Lists.AddMovie(ctx, h.ListID, h.SessionID, h.MovieID)
	if err != nil {
		return nil, err
	}
	return status, nil
}
