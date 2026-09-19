package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
)

// ListsClearHandler handles `lists -a clear -i <list_id>` (v3) and
// `lists -v4 -a clear -i <list_id>` (V4 set: user access token).
type ListsClearHandler struct {
	ListID      string
	SessionID   string
	V4          bool
	AccessToken string
}

// Handle removes all items from a list.
func (h ListsClearHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	if h.V4 {
		status, err := client.Lists.ClearListV4(ctx, h.AccessToken, h.ListID)
		if err != nil {
			return nil, err
		}
		return status, nil
	}

	status, _, err := client.Lists.ClearList(ctx, h.ListID, h.SessionID)
	if err != nil {
		return nil, err
	}
	return status, nil
}
