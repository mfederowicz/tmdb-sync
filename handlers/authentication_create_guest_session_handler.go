package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
)

// AuthenticationCreateGuestSessionHandler handles `authentication -a create-guest-session`.
type AuthenticationCreateGuestSessionHandler struct{}

// Handle requests a new guest session.
func (AuthenticationCreateGuestSessionHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	guestSession, _, err := client.Auth.CreateGuestSession(ctx)
	if err != nil {
		return nil, err
	}
	return guestSession, nil
}
