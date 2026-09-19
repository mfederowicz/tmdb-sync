package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
	"github.com/mfederowicz/tmdb-sync/uri"
)

// AccountRatedMoviesHandler handles `account -a rated-movies` (v3) and
// `account -v4 -a rated-movies` (V4 set: uses the v4 user access token and account id).
type AccountRatedMoviesHandler struct {
	AccountID   int64
	SessionID   string
	PagesLimit  int
	V4          bool
	AccessToken string
	V4AccountID string
	V4Options   uri.AccountV4Options
}

// Handle fetches an account's rated movies.
func (h AccountRatedMoviesHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	if h.V4 {
		movies, err := client.Account.GetRatedMoviesV4(ctx, h.AccessToken, h.V4AccountID, h.PagesLimit, h.V4Options)
		if err != nil {
			return nil, err
		}
		return movies, nil
	}

	movies, err := client.Account.GetRatedMovies(ctx, h.AccountID, h.SessionID, h.PagesLimit)
	if err != nil {
		return nil, err
	}
	return movies, nil
}
