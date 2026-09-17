package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
)

// TVPopularHandler handles `tv -a popular`, fetching pages up to PagesLimit
// (0 = unlimited, bounded by TMDB's total_pages).
type TVPopularHandler struct {
	PagesLimit int
}

// Handle fetches the popular-tv-series list.
func (h TVPopularHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	shows, err := client.TV.GetPopularTV(ctx, h.PagesLimit)
	if err != nil {
		return nil, err
	}
	return shows, nil
}
