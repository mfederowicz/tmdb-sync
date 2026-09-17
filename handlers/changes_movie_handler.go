package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
)

// ChangesMovieHandler handles `changes -a movie`.
type ChangesMovieHandler struct {
	StartDate  string
	EndDate    string
	PagesLimit int
}

// Handle fetches movie ids changed in the requested date range.
func (h ChangesMovieHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	changes, err := client.Changes.GetMovieChanges(ctx, h.StartDate, h.EndDate, h.PagesLimit)
	if err != nil {
		return nil, err
	}
	return changes, nil
}
