package internal

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"
)

func TestGetRatedMoviesGuestSession(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/guest_session/abc123/rated/movies", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		json.NewEncoder(w).Encode(map[string]any{
			"page":          1,
			"total_pages":   1,
			"total_results": 1,
			"results": []map[string]any{
				{"id": 550, "title": "Fight Club", "rating": 8.5},
			},
		})
	})

	movies, err := client.GuestSessions.GetRatedMovies(context.Background(), "abc123", "", "", 0)
	if err != nil {
		t.Fatalf("GetRatedMovies() error = %v", err)
	}
	if len(movies) != 1 || movies[0].Title != "Fight Club" {
		t.Errorf("GetRatedMovies() = %+v, want one result titled Fight Club", movies)
	}
}
