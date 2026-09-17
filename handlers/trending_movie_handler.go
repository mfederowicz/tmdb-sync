package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
)

// TrendingMovieHandler handles `trending -a movie -w <day|week>`.
type TrendingMovieHandler struct {
	TimeWindow string
	PagesLimit int
}

// Handle returns the day's or week's trending movies.
func (h TrendingMovieHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	results, err := client.Trending.TrendingMovies(ctx, h.TimeWindow, h.PagesLimit)
	if err != nil {
		return nil, err
	}
	return results, nil
}
