package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
)

// ListsClearHandler handles `lists -a clear -i <list_id>`.
type ListsClearHandler struct {
	ListID    string
	SessionID string
}

// Handle removes all items from a list.
func (h ListsClearHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	status, _, err := client.Lists.ClearList(ctx, h.ListID, h.SessionID)
	if err != nil {
		return nil, err
	}
	return status, nil
}
