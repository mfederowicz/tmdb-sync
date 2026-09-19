package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
	"github.com/mfederowicz/tmdb-sync/uri"
)

// AccountRecommendedTVHandler handles `account -v4 -a recommended-tv`
// (v4 only, no v3 equivalent).
type AccountRecommendedTVHandler struct {
	AccessToken string
	AccountID   string
	PagesLimit  int
	Options     uri.AccountV4Options
}

// Handle fetches TV show recommendations for a v4 account.
func (h AccountRecommendedTVHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	shows, err := client.Account.GetRecommendedTVV4(ctx, h.AccessToken, h.AccountID, h.PagesLimit, h.Options)
	if err != nil {
		return nil, err
	}
	return shows, nil
}
