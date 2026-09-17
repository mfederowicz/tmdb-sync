package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
)

// PeopleLatestHandler handles `people -a latest`.
type PeopleLatestHandler struct{}

// Handle fetches the most recently created person entry.
func (h PeopleLatestHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	person, _, err := client.People.GetLatest(ctx)
	if err != nil {
		return nil, err
	}
	return person, nil
}
