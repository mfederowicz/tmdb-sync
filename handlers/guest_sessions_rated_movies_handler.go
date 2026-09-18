package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
)

// GuestSessionsRatedMoviesHandler handles `guest-sessions -a rated-movies`.
type GuestSessionsRatedMoviesHandler struct {
	GuestSessionID string
	Language       string
	SortBy         string
	PagesLimit     int
}

// Handle fetches a guest session's rated movies.
func (h GuestSessionsRatedMoviesHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	movies, err := client.GuestSessions.GetRatedMovies(ctx, h.GuestSessionID, h.Language, h.SortBy, h.PagesLimit)
	if err != nil {
		return nil, err
	}
	return movies, nil
}
