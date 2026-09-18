package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
)

// GuestSessionsCreateHandler handles `guest-sessions -a create`.
type GuestSessionsCreateHandler struct{}

// Handle requests a new guest session.
func (GuestSessionsCreateHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	guestSession, _, err := client.Auth.CreateGuestSession(ctx)
	if err != nil {
		return nil, err
	}
	return guestSession, nil
}
