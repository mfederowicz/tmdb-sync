package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
)

// PeopleImagesHandler handles `people -a images -i <person_id>`.
type PeopleImagesHandler struct {
	PersonID int64
}

// Handle fetches the profile images for a single person.
func (h PeopleImagesHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	images, _, err := client.People.GetPersonImages(ctx, h.PersonID)
	if err != nil {
		return nil, err
	}
	return images, nil
}
