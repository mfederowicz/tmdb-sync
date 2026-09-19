package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
	"github.com/mfederowicz/tmdb-sync/str"
)

// ListsUpdateHandler handles `lists -v4 -a update -i <list_id> ...` (v4 only).
type ListsUpdateHandler struct {
	ListID      string
	AccessToken string
	Body        str.ListUpdateRequestV4
}

// Handle updates the given fields of a list.
func (h ListsUpdateHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	status, err := client.Lists.UpdateListV4(ctx, h.AccessToken, h.ListID, &h.Body)
	if err != nil {
		return nil, err
	}
	return status, nil
}
