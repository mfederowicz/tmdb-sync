package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
)

// AccountWatchlistMoviesHandler handles `account -a watchlist-movies`.
type AccountWatchlistMoviesHandler struct {
	AccountID  int64
	SessionID  string
	PagesLimit int
}

// Handle fetches an account's watchlisted movies.
func (h AccountWatchlistMoviesHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	movies, err := client.Account.GetWatchlistMovies(ctx, h.AccountID, h.SessionID, h.PagesLimit)
	if err != nil {
		return nil, err
	}
	return movies, nil
}
