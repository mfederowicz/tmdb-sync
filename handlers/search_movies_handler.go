package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
	"github.com/mfederowicz/tmdb-sync/uri"
)

// SearchMoviesHandler handles `search -a movies -q <query>`.
type SearchMoviesHandler struct {
	Query              string
	Language           string
	IncludeAdult       bool
	Region             string
	Year               int
	PrimaryReleaseYear int
	PagesLimit         int
}

// Handle searches for movies matching the query.
func (h SearchMoviesHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	opts := uri.SearchMovieOptions{
		Query:              h.Query,
		Language:           h.Language,
		IncludeAdult:       h.IncludeAdult,
		Region:             h.Region,
		Year:               h.Year,
		PrimaryReleaseYear: h.PrimaryReleaseYear,
	}
	results, err := client.Search.SearchMovies(ctx, opts, h.PagesLimit)
	if err != nil {
		return nil, err
	}
	return results, nil
}
