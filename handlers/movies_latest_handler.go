package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
)

// MoviesLatestHandler handles `movies -a latest`.
type MoviesLatestHandler struct{}

// Handle fetches the most recently created movie on TMDB.
func (MoviesLatestHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	movie, _, err := client.Movies.GetLatest(ctx)
	if err != nil {
		return nil, err
	}
	return movie, nil
}
