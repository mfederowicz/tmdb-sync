package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
)

// MoviesAddRatingHandler handles `movies -a add-rating -i <movie_id> -value <rating>`.
type MoviesAddRatingHandler struct {
	MovieID        int64
	SessionID      string
	GuestSessionID string
	Value          float64
}

// Handle rates a movie on behalf of the session's account, or a guest
// session when GuestSessionID is set instead of SessionID.
func (h MoviesAddRatingHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	status, _, err := client.Movies.AddRating(ctx, h.MovieID, h.SessionID, h.GuestSessionID, h.Value)
	if err != nil {
		return nil, err
	}
	return status, nil
}
