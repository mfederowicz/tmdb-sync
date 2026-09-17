package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
)

// ListsItemStatusHandler handles `lists -a item-status -i <list_id> -media-id <movie_id>`.
type ListsItemStatusHandler struct {
	ListID  string
	MovieID int64
}

// Handle reports whether a movie is present on a list.
func (h ListsItemStatusHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	status, _, err := client.Lists.GetItemStatus(ctx, h.ListID, h.MovieID)
	if err != nil {
		return nil, err
	}
	return status, nil
}
