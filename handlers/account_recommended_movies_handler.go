package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
	"github.com/mfederowicz/tmdb-sync/uri"
)

// AccountRecommendedMoviesHandler handles `account -v4 -a recommended-movies`
// (v4 only, no v3 equivalent).
type AccountRecommendedMoviesHandler struct {
	AccessToken string
	AccountID   string
	PagesLimit  int
	Options     uri.AccountV4Options
}

// Handle fetches movie recommendations for a v4 account.
func (h AccountRecommendedMoviesHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	movies, err := client.Account.GetRecommendedMoviesV4(ctx, h.AccessToken, h.AccountID, h.PagesLimit, h.Options)
	if err != nil {
		return nil, err
	}
	return movies, nil
}
