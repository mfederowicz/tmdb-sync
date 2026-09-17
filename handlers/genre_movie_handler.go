package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
)

// GenreMovieHandler handles `genres -a movie`.
type GenreMovieHandler struct{}

// Handle fetches TMDB's list of official movie genres.
func (GenreMovieHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	genres, _, err := client.Genre.GetMovieList(ctx)
	if err != nil {
		return nil, err
	}
	return genres, nil
}
