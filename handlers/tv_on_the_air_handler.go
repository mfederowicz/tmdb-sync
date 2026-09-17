package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
)

// TVOnTheAirHandler handles `tv -a on-the-air`, fetching pages up to
// PagesLimit (0 = unlimited, bounded by TMDB's total_pages).
type TVOnTheAirHandler struct {
	PagesLimit int
}

// Handle fetches the on-the-air-tv-series list.
func (h TVOnTheAirHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	shows, err := client.TV.GetOnTheAirTV(ctx, h.PagesLimit)
	if err != nil {
		return nil, err
	}
	return shows, nil
}
