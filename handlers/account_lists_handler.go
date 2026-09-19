package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
)

// AccountListsHandler handles `account -a lists` (v3) and
// `account -v4 -a lists` (V4 set: uses the v4 user access token and account id).
type AccountListsHandler struct {
	AccountID   int64
	SessionID   string
	PagesLimit  int
	V4          bool
	AccessToken string
	V4AccountID string
}

// Handle fetches an account's custom lists.
func (h AccountListsHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	if h.V4 {
		lists, err := client.Account.GetListsV4(ctx, h.AccessToken, h.V4AccountID, h.PagesLimit)
		if err != nil {
			return nil, err
		}
		return lists, nil
	}

	lists, err := client.Account.GetLists(ctx, h.AccountID, h.SessionID, h.PagesLimit)
	if err != nil {
		return nil, err
	}
	return lists, nil
}
