package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
)

// TVListsHandler handles `tv -a lists -i <series_id>`.
type TVListsHandler struct {
	SeriesID   int64
	Language   string
	PagesLimit int
}

// Handle fetches the lists a TV series belongs to.
func (h TVListsHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	lists, err := client.TV.GetLists(ctx, h.SeriesID, h.Language, h.PagesLimit)
	if err != nil {
		return nil, err
	}
	return lists, nil
}
