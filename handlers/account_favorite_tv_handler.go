package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
)

// AccountFavoriteTVHandler handles `account -a favorite-tv`.
type AccountFavoriteTVHandler struct {
	AccountID  int64
	SessionID  string
	PagesLimit int
}

// Handle fetches an account's favorited TV shows.
func (h AccountFavoriteTVHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	shows, err := client.Account.GetFavoriteTV(ctx, h.AccountID, h.SessionID, h.PagesLimit)
	if err != nil {
		return nil, err
	}
	return shows, nil
}
