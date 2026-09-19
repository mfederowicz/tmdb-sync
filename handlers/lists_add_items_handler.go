package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
	"github.com/mfederowicz/tmdb-sync/str"
)

// ListsAddItemsHandler handles `lists -v4 -a add-items -i <list_id> -item <type:id> ...` (v4 only).
type ListsAddItemsHandler struct {
	ListID      string
	AccessToken string
	Items       []str.ListMediaV4
}

// Handle adds the items to a list.
func (h ListsAddItemsHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	result, err := client.Lists.AddItemsV4(ctx, h.AccessToken, h.ListID, h.Items)
	if err != nil {
		return nil, err
	}
	return result, nil
}
