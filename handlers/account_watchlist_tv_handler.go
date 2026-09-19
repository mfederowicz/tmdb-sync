package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
	"github.com/mfederowicz/tmdb-sync/uri"
)

// AccountWatchlistTVHandler handles `account -a watchlist-tv` (v3) and
// `account -v4 -a watchlist-tv` (V4 set: uses the v4 user access token and account id).
type AccountWatchlistTVHandler struct {
	AccountID   int64
	SessionID   string
	PagesLimit  int
	V4          bool
	AccessToken string
	V4AccountID string
	V4Options   uri.AccountV4Options
}

// Handle fetches an account's watchlisted TV shows.
func (h AccountWatchlistTVHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	if h.V4 {
		shows, err := client.Account.GetWatchlistTVV4(ctx, h.AccessToken, h.V4AccountID, h.PagesLimit, h.V4Options)
		if err != nil {
			return nil, err
		}
		return shows, nil
	}

	shows, err := client.Account.GetWatchlistTV(ctx, h.AccountID, h.SessionID, h.PagesLimit)
	if err != nil {
		return nil, err
	}
	return shows, nil
}
