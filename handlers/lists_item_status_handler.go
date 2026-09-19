package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
)

// ListsItemStatusHandler handles `lists -a item-status -i <list_id> -media-id <movie_id>`
// (v3) and `lists -v4 -a item-status -i <list_id> -media-type <type> -media-id <id>`
// (V4 set: media type and optional user access token).
type ListsItemStatusHandler struct {
	ListID      string
	MovieID     int64
	V4          bool
	AccessToken string
	MediaType   string
}

// Handle reports whether a movie is present on a list.
func (h ListsItemStatusHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	if h.V4 {
		status, err := client.Lists.GetItemStatusV4(ctx, h.AccessToken, h.ListID, h.MediaType, h.MovieID)
		if err != nil {
			return nil, err
		}
		return status, nil
	}

	status, _, err := client.Lists.GetItemStatus(ctx, h.ListID, h.MovieID)
	if err != nil {
		return nil, err
	}
	return status, nil
}
