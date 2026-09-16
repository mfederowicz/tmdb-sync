package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
)

// AccountRatedMoviesHandler handles `account -a rated-movies`.
type AccountRatedMoviesHandler struct {
	AccountID  int64
	SessionID  string
	PagesLimit int
}

// Handle fetches an account's rated movies.
func (h AccountRatedMoviesHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	movies, err := client.Account.GetRatedMovies(ctx, h.AccountID, h.SessionID, h.PagesLimit)
	if err != nil {
		return nil, err
	}
	return movies, nil
}
