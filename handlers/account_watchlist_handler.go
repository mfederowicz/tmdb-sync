package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
	"github.com/mfederowicz/tmdb-sync/str"
)

// AccountWatchlistHandler handles `account -a add-watchlist`.
type AccountWatchlistHandler struct {
	AccountID int64
	SessionID string
	MediaType string
	MediaID   int64
	Watchlist bool
}

// Handle adds or removes a movie/TV show from an account's watchlist.
func (h AccountWatchlistHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	status, _, err := client.Account.AddToWatchlist(ctx, h.AccountID, h.SessionID, &str.AccountWatchlistRequest{
		MediaType: h.MediaType,
		MediaID:   h.MediaID,
		Watchlist: h.Watchlist,
	})
	if err != nil {
		return nil, err
	}
	return status, nil
}
