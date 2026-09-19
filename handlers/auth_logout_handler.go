package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
)

// AuthLogoutHandler handles `auth -v4 -a logout`.
type AuthLogoutHandler struct {
	AccessToken string
}

// Handle invalidates the v4 user access token.
func (h AuthLogoutHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	status, _, err := client.Auth.DeleteAccessTokenV4(ctx, h.AccessToken)
	if err != nil {
		return nil, err
	}
	return status, nil
}
