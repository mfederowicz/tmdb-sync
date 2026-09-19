package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
	"github.com/mfederowicz/tmdb-sync/str"
)

// ListsRemoveItemsHandler handles `lists -v4 -a remove-items -i <list_id> -item <type:id> ...` (v4 only).
type ListsRemoveItemsHandler struct {
	ListID      string
	AccessToken string
	Items       []str.ListMediaV4
}

// Handle removes the items from a list.
func (h ListsRemoveItemsHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	result, err := client.Lists.RemoveItemsV4(ctx, h.AccessToken, h.ListID, h.Items)
	if err != nil {
		return nil, err
	}
	return result, nil
}
