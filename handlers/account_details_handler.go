package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
)

// AccountDetailsHandler handles `account -a details -i <account_id>`.
type AccountDetailsHandler struct {
	AccountID int64
	SessionID string
}

// Handle fetches an account's details.
func (h AccountDetailsHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	account, _, err := client.Account.GetDetails(ctx, h.AccountID, h.SessionID)
	if err != nil {
		return nil, err
	}
	return account, nil
}
