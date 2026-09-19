package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
	"github.com/mfederowicz/tmdb-sync/uri"
)

// AccountFavoriteMoviesHandler handles `account -a favorite-movies` (v3) and
// `account -v4 -a favorite-movies` (V4 set: uses the v4 user access token and account id).
type AccountFavoriteMoviesHandler struct {
	AccountID   int64
	SessionID   string
	PagesLimit  int
	V4          bool
	AccessToken string
	V4AccountID string
	V4Options   uri.AccountV4Options
}

// Handle fetches an account's favorited movies.
func (h AccountFavoriteMoviesHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	if h.V4 {
		movies, err := client.Account.GetFavoriteMoviesV4(ctx, h.AccessToken, h.V4AccountID, h.PagesLimit, h.V4Options)
		if err != nil {
			return nil, err
		}
		return movies, nil
	}

	movies, err := client.Account.GetFavoriteMovies(ctx, h.AccountID, h.SessionID, h.PagesLimit)
	if err != nil {
		return nil, err
	}
	return movies, nil
}
