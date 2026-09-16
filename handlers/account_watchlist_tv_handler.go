package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
)

// AccountWatchlistTVHandler handles `account -a watchlist-tv`.
type AccountWatchlistTVHandler struct {
	AccountID  int64
	SessionID  string
	PagesLimit int
}

// Handle fetches an account's watchlisted TV shows.
func (h AccountWatchlistTVHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	shows, err := client.Account.GetWatchlistTV(ctx, h.AccountID, h.SessionID, h.PagesLimit)
	if err != nil {
		return nil, err
	}
	return shows, nil
}
