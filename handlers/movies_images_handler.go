package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
	"github.com/mfederowicz/tmdb-sync/uri"
)

// MoviesImagesHandler handles `movies -a images -i <movie_id>`.
type MoviesImagesHandler struct {
	MovieID              int64
	Language             string
	IncludeImageLanguage string
}

// Handle fetches the images for a movie.
func (h MoviesImagesHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	images, _, err := client.Movies.GetImages(ctx, h.MovieID, &uri.ImagesOptions{
		Language:             h.Language,
		IncludeImageLanguage: h.IncludeImageLanguage,
	})
	if err != nil {
		return nil, err
	}
	return images, nil
}
