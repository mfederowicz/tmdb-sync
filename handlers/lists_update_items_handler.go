package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
	"github.com/mfederowicz/tmdb-sync/str"
)

// ListsUpdateItemsHandler handles `lists -v4 -a update-items -i <list_id> -item <type:id:comment> ...` (v4 only).
type ListsUpdateItemsHandler struct {
	ListID      string
	AccessToken string
	Items       []str.ListMediaV4
}

// Handle updates the comments of items on a list.
func (h ListsUpdateItemsHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	result, err := client.Lists.UpdateItemsV4(ctx, h.AccessToken, h.ListID, h.Items)
	if err != nil {
		return nil, err
	}
	return result, nil
}
