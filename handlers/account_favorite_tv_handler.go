package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
)

// AccountFavoriteTVHandler handles `account -a favorite-tv` (v3) and
// `account -v4 -a favorite-tv` (V4 set: uses the v4 user access token and account id).
type AccountFavoriteTVHandler struct {
	AccountID   int64
	SessionID   string
	PagesLimit  int
	V4          bool
	AccessToken string
	V4AccountID string
}

// Handle fetches an account's favorited TV shows.
func (h AccountFavoriteTVHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	if h.V4 {
		shows, err := client.Account.GetFavoriteTVV4(ctx, h.AccessToken, h.V4AccountID, h.PagesLimit)
		if err != nil {
			return nil, err
		}
		return shows, nil
	}

	shows, err := client.Account.GetFavoriteTV(ctx, h.AccountID, h.SessionID, h.PagesLimit)
	if err != nil {
		return nil, err
	}
	return shows, nil
}
