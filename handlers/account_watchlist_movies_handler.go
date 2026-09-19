package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
)

// AccountWatchlistMoviesHandler handles `account -a watchlist-movies` (v3) and
// `account -v4 -a watchlist-movies` (V4 set: uses the v4 user access token and account id).
type AccountWatchlistMoviesHandler struct {
	AccountID   int64
	SessionID   string
	PagesLimit  int
	V4          bool
	AccessToken string
	V4AccountID string
}

// Handle fetches an account's watchlisted movies.
func (h AccountWatchlistMoviesHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	if h.V4 {
		movies, err := client.Account.GetWatchlistMoviesV4(ctx, h.AccessToken, h.V4AccountID, h.PagesLimit)
		if err != nil {
			return nil, err
		}
		return movies, nil
	}

	movies, err := client.Account.GetWatchlistMovies(ctx, h.AccountID, h.SessionID, h.PagesLimit)
	if err != nil {
		return nil, err
	}
	return movies, nil
}
