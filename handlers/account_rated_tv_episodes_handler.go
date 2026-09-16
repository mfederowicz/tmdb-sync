package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
)

// AccountRatedTVEpisodesHandler handles `account -a rated-tv-episodes`.
type AccountRatedTVEpisodesHandler struct {
	AccountID  int64
	SessionID  string
	PagesLimit int
}

// Handle fetches an account's rated TV episodes.
func (h AccountRatedTVEpisodesHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	episodes, err := client.Account.GetRatedTVEpisodes(ctx, h.AccountID, h.SessionID, h.PagesLimit)
	if err != nil {
		return nil, err
	}
	return episodes, nil
}
