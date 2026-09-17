package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
)

// TVTopRatedHandler handles `tv -a top-rated`, fetching pages up to
// PagesLimit (0 = unlimited, bounded by TMDB's total_pages).
type TVTopRatedHandler struct {
	PagesLimit int
}

// Handle fetches the top-rated-tv-series list.
func (h TVTopRatedHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	shows, err := client.TV.GetTopRatedTV(ctx, h.PagesLimit)
	if err != nil {
		return nil, err
	}
	return shows, nil
}
