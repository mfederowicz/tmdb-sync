package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
	"github.com/mfederowicz/tmdb-sync/uri"
)

// AccountRatedTVHandler handles `account -a rated-tv` (v3) and
// `account -v4 -a rated-tv` (V4 set: uses the v4 user access token and account id).
type AccountRatedTVHandler struct {
	AccountID   int64
	SessionID   string
	PagesLimit  int
	V4          bool
	AccessToken string
	V4AccountID string
	V4Options   uri.AccountV4Options
}

// Handle fetches an account's rated TV shows.
func (h AccountRatedTVHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	if h.V4 {
		shows, err := client.Account.GetRatedTVV4(ctx, h.AccessToken, h.V4AccountID, h.PagesLimit, h.V4Options)
		if err != nil {
			return nil, err
		}
		return shows, nil
	}

	shows, err := client.Account.GetRatedTV(ctx, h.AccountID, h.SessionID, h.PagesLimit)
	if err != nil {
		return nil, err
	}
	return shows, nil
}
