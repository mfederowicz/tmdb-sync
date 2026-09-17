package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
)

// TrendingAllHandler handles `trending -a all -w <day|week>`.
type TrendingAllHandler struct {
	TimeWindow string
	PagesLimit int
}

// Handle returns the day's or week's trending movies, tv shows, and people.
func (h TrendingAllHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	results, err := client.Trending.TrendingAll(ctx, h.TimeWindow, h.PagesLimit)
	if err != nil {
		return nil, err
	}
	return results, nil
}
