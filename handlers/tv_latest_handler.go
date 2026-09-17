package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
)

// TVLatestHandler handles `tv -a latest`.
type TVLatestHandler struct{}

// Handle fetches the most recently created TV series on TMDB.
func (TVLatestHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	tv, _, err := client.TV.GetLatest(ctx)
	if err != nil {
		return nil, err
	}
	return tv, nil
}
