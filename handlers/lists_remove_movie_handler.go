package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
)

// ListsRemoveMovieHandler handles `lists -a remove-movie -i <list_id> -media-id <movie_id>`.
type ListsRemoveMovieHandler struct {
	ListID    string
	SessionID string
	MovieID   int64
}

// Handle removes a movie from a list.
func (h ListsRemoveMovieHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	status, _, err := client.Lists.RemoveMovie(ctx, h.ListID, h.SessionID, h.MovieID)
	if err != nil {
		return nil, err
	}
	return status, nil
}
