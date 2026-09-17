package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
)

// PeopleDetailsHandler handles `people -a details -i <person_id>`.
type PeopleDetailsHandler struct {
	PersonID int64
}

// Handle fetches details for a single person.
func (h PeopleDetailsHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	person, _, err := client.People.GetPerson(ctx, h.PersonID)
	if err != nil {
		return nil, err
	}
	return person, nil
}
