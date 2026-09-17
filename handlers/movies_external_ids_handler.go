package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
)

// MoviesExternalIDsHandler handles `movies -a external-ids -i <movie_id>`.
type MoviesExternalIDsHandler struct {
	MovieID int64
}

// Handle fetches a movie's external ids.
func (h MoviesExternalIDsHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	ids, _, err := client.Movies.GetExternalIDs(ctx, h.MovieID)
	if err != nil {
		return nil, err
	}
	return ids, nil
}
