package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
)

// MoviesVideosHandler handles `movies -a videos -i <movie_id>`.
type MoviesVideosHandler struct {
	MovieID  int64
	Language string
}

// Handle fetches the videos (trailers, teasers, ...) for a movie.
func (h MoviesVideosHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	videos, _, err := client.Movies.GetVideos(ctx, h.MovieID, h.Language)
	if err != nil {
		return nil, err
	}
	return videos, nil
}
