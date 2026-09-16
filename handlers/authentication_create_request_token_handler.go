package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
)

// AuthenticationCreateRequestTokenHandler handles `authentication -a create-request-token`.
type AuthenticationCreateRequestTokenHandler struct{}

// Handle requests a new, unapproved request token.
func (AuthenticationCreateRequestTokenHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	token, _, err := client.Auth.CreateRequestToken(ctx)
	if err != nil {
		return nil, err
	}
	return token, nil
}
