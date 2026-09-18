package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
)

// GuestSessionsRatedTVEpisodesHandler handles `guest-sessions -a rated-tv-episodes`.
type GuestSessionsRatedTVEpisodesHandler struct {
	GuestSessionID string
	Language       string
	SortBy         string
	PagesLimit     int
}

// Handle fetches a guest session's rated TV episodes.
func (h GuestSessionsRatedTVEpisodesHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	episodes, err := client.GuestSessions.GetRatedTVEpisodes(ctx, h.GuestSessionID, h.Language, h.SortBy, h.PagesLimit)
	if err != nil {
		return nil, err
	}
	return episodes, nil
}
