package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
)

// ListsDeleteHandler handles `lists -a delete -i <list_id>`.
type ListsDeleteHandler struct {
	ListID    string
	SessionID string
}

// Handle deletes a list.
func (h ListsDeleteHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	status, _, err := client.Lists.DeleteList(ctx, h.ListID, h.SessionID)
	if err != nil {
		return nil, err
	}
	return status, nil
}
