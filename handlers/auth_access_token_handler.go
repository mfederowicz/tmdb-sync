package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
)

// AuthAccessTokenHandler handles `auth -v4 -a access-token`.
type AuthAccessTokenHandler struct {
	RequestToken string
}

// Handle exchanges an approved v4 request token for a user access token.
func (h AuthAccessTokenHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	accessToken, _, err := client.Auth.CreateAccessTokenV4(ctx, h.RequestToken)
	if err != nil {
		return nil, err
	}
	return accessToken, nil
}
