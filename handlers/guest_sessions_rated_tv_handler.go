package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
)

// GuestSessionsRatedTVHandler handles `guest-sessions -a rated-tv`.
type GuestSessionsRatedTVHandler struct {
	GuestSessionID string
	Language       string
	SortBy         string
	PagesLimit     int
}

// Handle fetches a guest session's rated TV shows.
func (h GuestSessionsRatedTVHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	tvShows, err := client.GuestSessions.GetRatedTV(ctx, h.GuestSessionID, h.Language, h.SortBy, h.PagesLimit)
	if err != nil {
		return nil, err
	}
	return tvShows, nil
}
