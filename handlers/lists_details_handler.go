package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
)

// ListsDetailsHandler handles `lists -a details -i <list_id>`.
type ListsDetailsHandler struct {
	ListID string
}

// Handle fetches details for a single list.
func (h ListsDetailsHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	list, _, err := client.Lists.GetList(ctx, h.ListID)
	if err != nil {
		return nil, err
	}
	return list, nil
}
