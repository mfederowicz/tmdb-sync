package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
)

// ListsDeleteHandler handles `lists -a delete -i <list_id>` (v3) and
// `lists -v4 -a delete -i <list_id>` (V4 set: user access token).
type ListsDeleteHandler struct {
	ListID      string
	SessionID   string
	V4          bool
	AccessToken string
}

// Handle deletes a list.
func (h ListsDeleteHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	if h.V4 {
		status, err := client.Lists.DeleteListV4(ctx, h.AccessToken, h.ListID)
		if err != nil {
			return nil, err
		}
		return status, nil
	}

	status, _, err := client.Lists.DeleteList(ctx, h.ListID, h.SessionID)
	if err != nil {
		return nil, err
	}
	return status, nil
}
