package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
	"github.com/mfederowicz/tmdb-sync/str"
)

// AccountFavoriteHandler handles `account -a add-favorite`.
type AccountFavoriteHandler struct {
	AccountID int64
	SessionID string
	MediaType string
	MediaID   int64
	Favorite  bool
}

// Handle marks or unmarks a movie/TV show as one of an account's favorites.
func (h AccountFavoriteHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	status, _, err := client.Account.AddRemoveFavorite(ctx, h.AccountID, h.SessionID, &str.AccountFavoriteRequest{
		MediaType: h.MediaType,
		MediaID:   h.MediaID,
		Favorite:  h.Favorite,
	})
	if err != nil {
		return nil, err
	}
	return status, nil
}
