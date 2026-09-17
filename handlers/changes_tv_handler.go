package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
)

// ChangesTVHandler handles `changes -a tv`.
type ChangesTVHandler struct {
	StartDate  string
	EndDate    string
	PagesLimit int
}

// Handle fetches TV show ids changed in the requested date range.
func (h ChangesTVHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	changes, err := client.Changes.GetTVChanges(ctx, h.StartDate, h.EndDate, h.PagesLimit)
	if err != nil {
		return nil, err
	}
	return changes, nil
}
