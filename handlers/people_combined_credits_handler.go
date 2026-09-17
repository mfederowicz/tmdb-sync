package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
)

// PeopleCombinedCreditsHandler handles `people -a combined-credits -i <person_id>`.
type PeopleCombinedCreditsHandler struct {
	PersonID int64
}

// Handle fetches the combined (movie + TV) credits for a single person.
func (h PeopleCombinedCreditsHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	credits, _, err := client.People.GetPersonCombinedCredits(ctx, h.PersonID)
	if err != nil {
		return nil, err
	}
	return credits, nil
}
