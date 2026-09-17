package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
)

// TrendingPersonHandler handles `trending -a person -w <day|week>`.
type TrendingPersonHandler struct {
	TimeWindow string
	PagesLimit int
}

// Handle returns the day's or week's trending people.
func (h TrendingPersonHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	results, err := client.Trending.TrendingPeople(ctx, h.TimeWindow, h.PagesLimit)
	if err != nil {
		return nil, err
	}
	return results, nil
}
