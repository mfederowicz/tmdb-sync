package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
)

// AuthLogoutHandler handles `auth -a logout` (v3: invalidates the cached
// session id) and `auth -v4 -a logout` (V4 set: invalidates the user access token).
type AuthLogoutHandler struct {
	SessionID   string
	V4          bool
	AccessToken string
}

// Handle invalidates the v3 session or the v4 user access token.
func (h AuthLogoutHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	if !h.V4 {
		status, _, err := client.Auth.DeleteSession(ctx, h.SessionID)
		if err != nil {
			return nil, err
		}
		return status, nil
	}

	status, _, err := client.Auth.DeleteAccessTokenV4(ctx, h.AccessToken)
	if err != nil {
		return nil, err
	}
	return status, nil
}
