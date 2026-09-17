package internal

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/mfederowicz/tmdb-sync/uri"
)

func TestGetDiscoverMovies(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/discover/movie", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		if got := r.URL.Query().Get("with_genres"); got != "28" {
			t.Errorf("with_genres = %q, want %q", got, "28")
		}
		json.NewEncoder(w).Encode(map[string]any{
			"page":          1,
			"total_pages":   1,
			"total_results": 1,
			"results": []map[string]any{
				{"id": 550, "title": "Fight Club"},
			},
		})
	})

	movies, err := client.Discover.GetDiscoverMovies(context.Background(), uri.DiscoverMovieOptions{WithGenres: "28"}, 0)
	if err != nil {
		t.Fatalf("GetDiscoverMovies() error = %v", err)
	}
	if len(movies) != 1 || movies[0].Title != "Fight Club" {
		t.Errorf("GetDiscoverMovies() = %+v, want one result titled Fight Club", movies)
	}
}

func TestGetDiscoverTV(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/discover/tv", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		json.NewEncoder(w).Encode(map[string]any{
			"page":          1,
			"total_pages":   1,
			"total_results": 1,
			"results": []map[string]any{
				{"id": 1399, "name": "Game of Thrones"},
			},
		})
	})

	shows, err := client.Discover.GetDiscoverTV(context.Background(), uri.DiscoverTVOptions{}, 0)
	if err != nil {
		t.Fatalf("GetDiscoverTV() error = %v", err)
	}
	if len(shows) != 1 || shows[0].Name != "Game of Thrones" {
		t.Errorf("GetDiscoverTV() = %+v, want one result named Game of Thrones", shows)
	}
}
