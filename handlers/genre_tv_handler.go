package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
)

// GenreTVHandler handles `genres -a tv`.
type GenreTVHandler struct{}

// Handle fetches TMDB's list of official TV genres.
func (GenreTVHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	genres, _, err := client.Genre.GetTVList(ctx)
	if err != nil {
		return nil, err
	}
	return genres, nil
}
