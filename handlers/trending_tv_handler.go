package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
)

// TrendingTVHandler handles `trending -a tv -w <day|week>`.
type TrendingTVHandler struct {
	TimeWindow string
	PagesLimit int
}

// Handle returns the day's or week's trending TV shows.
func (h TrendingTVHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	results, err := client.Trending.TrendingTV(ctx, h.TimeWindow, h.PagesLimit)
	if err != nil {
		return nil, err
	}
	return results, nil
}
