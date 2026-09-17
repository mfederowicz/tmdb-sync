package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
)

// PeoplePopularHandler handles `people -a popular`, fetching pages up to
// PagesLimit (0 = unlimited, bounded by TMDB's total_pages).
type PeoplePopularHandler struct {
	PagesLimit int
}

// Handle fetches the popular-people list.
func (h PeoplePopularHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	persons, err := client.People.GetPopularPeople(ctx, h.PagesLimit)
	if err != nil {
		return nil, err
	}
	return persons, nil
}
