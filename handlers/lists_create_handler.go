package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
	"github.com/mfederowicz/tmdb-sync/str"
)

// ListsCreateHandler handles `lists -a create -name <name> ...`.
type ListsCreateHandler struct {
	SessionID   string
	Name        string
	Description string
	Language    string
}

// Handle creates a new list owned by the session's account.
func (h ListsCreateHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	created, _, err := client.Lists.CreateList(ctx, h.SessionID, &str.ListCreateRequest{
		Name:        h.Name,
		Description: h.Description,
		Language:    h.Language,
	})
	if err != nil {
		return nil, err
	}
	return created, nil
}
