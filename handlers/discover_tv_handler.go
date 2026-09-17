package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
	"github.com/mfederowicz/tmdb-sync/uri"
)

// DiscoverTVHandler handles `discover -a tv [filters]`, fetching pages up to
// PagesLimit (0 = unlimited, bounded by TMDB's total_pages).
type DiscoverTVHandler struct {
	Options    uri.DiscoverTVOptions
	PagesLimit int
}

// Handle fetches the discover-tv list matching Options.
func (h DiscoverTVHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	shows, err := client.Discover.GetDiscoverTV(ctx, h.Options, h.PagesLimit)
	if err != nil {
		return nil, err
	}
	return shows, nil
}
