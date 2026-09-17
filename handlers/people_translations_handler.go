package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
)

// PeopleTranslationsHandler handles `people -a translations -i <person_id>`.
type PeopleTranslationsHandler struct {
	PersonID int64
}

// Handle fetches the translations for a single person.
func (h PeopleTranslationsHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	translations, _, err := client.People.GetPersonTranslations(ctx, h.PersonID)
	if err != nil {
		return nil, err
	}
	return translations, nil
}
