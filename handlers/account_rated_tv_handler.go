package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
)

// AccountRatedTVHandler handles `account -a rated-tv`.
type AccountRatedTVHandler struct {
	AccountID  int64
	SessionID  string
	PagesLimit int
}

// Handle fetches an account's rated TV shows.
func (h AccountRatedTVHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	shows, err := client.Account.GetRatedTV(ctx, h.AccountID, h.SessionID, h.PagesLimit)
	if err != nil {
		return nil, err
	}
	return shows, nil
}
