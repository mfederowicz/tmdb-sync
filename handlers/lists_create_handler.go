package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
	"github.com/mfederowicz/tmdb-sync/str"
)

// ListsCreateHandler handles `lists -a create -name <name> ...` (v3) and
// `lists -v4 -a create ...` (V4 set: user access token, country, visibility).
type ListsCreateHandler struct {
	SessionID   string
	Name        string
	Description string
	Language    string
	V4          bool
	AccessToken string
	Country     string
	Public      bool
}

// Handle creates a new list owned by the session's account.
func (h ListsCreateHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	if h.V4 {
		created, err := client.Lists.CreateListV4(ctx, h.AccessToken, &str.ListCreateRequestV4{
			Name:        h.Name,
			Description: h.Description,
			ISO6391:     h.Language,
			ISO31661:    h.Country,
			Public:      h.Public,
		})
		if err != nil {
			return nil, err
		}
		return created, nil
	}

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
