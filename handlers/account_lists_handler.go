package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
)

// AccountListsHandler handles `account -a lists`.
type AccountListsHandler struct {
	AccountID  int64
	SessionID  string
	PagesLimit int
}

// Handle fetches an account's custom lists.
func (h AccountListsHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	lists, err := client.Account.GetLists(ctx, h.AccountID, h.SessionID, h.PagesLimit)
	if err != nil {
		return nil, err
	}
	return lists, nil
}
