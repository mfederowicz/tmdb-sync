package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
)

// PeopleExternalIDsHandler handles `people -a external-ids -i <person_id>`.
type PeopleExternalIDsHandler struct {
	PersonID int64
}

// Handle fetches the external ids for a single person.
func (h PeopleExternalIDsHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	ids, _, err := client.People.GetPersonExternalIDs(ctx, h.PersonID)
	if err != nil {
		return nil, err
	}
	return ids, nil
}
