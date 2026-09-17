package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
)

// PeopleMovieCreditsHandler handles `people -a movie-credits -i <person_id>`.
type PeopleMovieCreditsHandler struct {
	PersonID int64
}

// Handle fetches the movie credits for a single person.
func (h PeopleMovieCreditsHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	credits, _, err := client.People.GetPersonMovieCredits(ctx, h.PersonID)
	if err != nil {
		return nil, err
	}
	return credits, nil
}
