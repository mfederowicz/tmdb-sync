package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
)

// MoviesAccountStatesHandler handles `movies -a account-states -i <movie_id>`.
type MoviesAccountStatesHandler struct {
	MovieID   int64
	SessionID string
}

// Handle fetches an account's favorite/rated/watchlist status for a movie.
func (h MoviesAccountStatesHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	states, _, err := client.Movies.GetAccountStates(ctx, h.MovieID, h.SessionID)
	if err != nil {
		return nil, err
	}
	return states, nil
}
