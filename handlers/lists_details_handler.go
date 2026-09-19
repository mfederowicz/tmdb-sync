package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
	"github.com/mfederowicz/tmdb-sync/uri"
)

// ListsDetailsHandler handles `lists -a details -i <list_id>` (v3) and
// `lists -v4 -a details -i <list_id>` (V4 set: optional user access token,
// paging and query options).
type ListsDetailsHandler struct {
	ListID      string
	V4          bool
	AccessToken string
	PagesLimit  int
	V4Options   uri.ListV4Options
}

// Handle fetches details for a single list.
func (h ListsDetailsHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	if h.V4 {
		list, err := client.Lists.GetListV4(ctx, h.AccessToken, h.ListID, h.PagesLimit, h.V4Options)
		if err != nil {
			return nil, err
		}
		return list, nil
	}

	list, _, err := client.Lists.GetList(ctx, h.ListID)
	if err != nil {
		return nil, err
	}
	return list, nil
}
