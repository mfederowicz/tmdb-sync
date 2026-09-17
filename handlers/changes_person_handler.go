package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
)

// ChangesPersonHandler handles `changes -a person`.
type ChangesPersonHandler struct {
	StartDate  string
	EndDate    string
	PagesLimit int
}

// Handle fetches person ids changed in the requested date range.
func (h ChangesPersonHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	changes, err := client.Changes.GetPersonChanges(ctx, h.StartDate, h.EndDate, h.PagesLimit)
	if err != nil {
		return nil, err
	}
	return changes, nil
}
