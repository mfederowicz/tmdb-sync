package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
)

// AccountFavoriteMoviesHandler handles `account -a favorite-movies`.
type AccountFavoriteMoviesHandler struct {
	AccountID  int64
	SessionID  string
	PagesLimit int
}

// Handle fetches an account's favorited movies.
func (h AccountFavoriteMoviesHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	movies, err := client.Account.GetFavoriteMovies(ctx, h.AccountID, h.SessionID, h.PagesLimit)
	if err != nil {
		return nil, err
	}
	return movies, nil
}
