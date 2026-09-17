package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
)

// PeopleTVCreditsHandler handles `people -a tv-credits -i <person_id>`.
type PeopleTVCreditsHandler struct {
	PersonID int64
}

// Handle fetches the TV credits for a single person.
func (h PeopleTVCreditsHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	credits, _, err := client.People.GetPersonTVCredits(ctx, h.PersonID)
	if err != nil {
		return nil, err
	}
	return credits, nil
}
