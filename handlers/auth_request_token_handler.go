package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
)

// AuthRequestTokenHandler handles `auth -v4 -a request-token`.
type AuthRequestTokenHandler struct {
	RedirectTo string
}

// Handle requests a new v4 request token.
func (h AuthRequestTokenHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	token, _, err := client.Auth.CreateRequestTokenV4(ctx, h.RedirectTo)
	if err != nil {
		return nil, err
	}
	return token, nil
}
