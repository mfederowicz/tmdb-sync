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

func TestGetRatedTVGuestSession(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/guest_session/abc123/rated/tv", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		json.NewEncoder(w).Encode(map[string]any{
			"page":          1,
			"total_pages":   1,
			"total_results": 1,
			"results": []map[string]any{
				{"id": 1399, "name": "Game of Thrones", "rating": 9.0},
			},
		})
	})

	tvShows, err := client.GuestSessions.GetRatedTV(context.Background(), "abc123", "", "", 0)
	if err != nil {
		t.Fatalf("GetRatedTV() error = %v", err)
	}
	if len(tvShows) != 1 || tvShows[0].Name != "Game of Thrones" {
		t.Errorf("GetRatedTV() = %+v, want one result named Game of Thrones", tvShows)
	}
}

func TestGetRatedTVEpisodesGuestSession(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/guest_session/abc123/rated/tv/episodes", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		json.NewEncoder(w).Encode(map[string]any{
			"page":          1,
			"total_pages":   1,
			"total_results": 1,
			"results": []map[string]any{
				{"id": 63056, "name": "Winter Is Coming", "season_number": 1, "episode_number": 1, "rating": 8.0},
			},
		})
	})

	episodes, err := client.GuestSessions.GetRatedTVEpisodes(context.Background(), "abc123", "", "", 0)
	if err != nil {
		t.Fatalf("GetRatedTVEpisodes() error = %v", err)
	}
	if len(episodes) != 1 || episodes[0].Name != "Winter Is Coming" {
		t.Errorf("GetRatedTVEpisodes() = %+v, want one result named Winter Is Coming", episodes)
	}
}
