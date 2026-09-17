package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
)

// TVAiringTodayHandler handles `tv -a airing-today`, fetching pages up to
// PagesLimit (0 = unlimited, bounded by TMDB's total_pages).
type TVAiringTodayHandler struct {
	PagesLimit int
}

// Handle fetches the airing-today-tv-series list.
func (h TVAiringTodayHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	shows, err := client.TV.GetAiringTodayTV(ctx, h.PagesLimit)
	if err != nil {
		return nil, err
	}
	return shows, nil
}
