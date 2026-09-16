package internal

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/mfederowicz/tmdb-sync/str"
)

func TestGetAccountDetails(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/account/123", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		if got := r.URL.Query().Get("session_id"); got != "abc123" {
			t.Errorf("session_id query param = %q, want %q", got, "abc123")
		}
		json.NewEncoder(w).Encode(map[string]any{
			"id":       123,
			"username": "testuser",
		})
	})

	account, _, err := client.Account.GetDetails(context.Background(), 123, "abc123")
	if err != nil {
		t.Fatalf("GetDetails() error = %v", err)
	}
	if account.ID != 123 || account.Username != "testuser" {
		t.Errorf("GetDetails() = %+v, want ID=123 Username=testuser", account)
	}
}

func TestGetAccountDetails_SelfResolve(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/account", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		if got := r.URL.Query().Get("session_id"); got != "abc123" {
			t.Errorf("session_id query param = %q, want %q", got, "abc123")
		}
		json.NewEncoder(w).Encode(map[string]any{
			"id":       456,
			"username": "selfuser",
		})
	})

	account, _, err := client.Account.GetDetails(context.Background(), 0, "abc123")
	if err != nil {
		t.Fatalf("GetDetails() error = %v", err)
	}
	if account.ID != 456 || account.Username != "selfuser" {
		t.Errorf("GetDetails() = %+v, want ID=456 Username=selfuser", account)
	}
}

func TestAddToWatchlist(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/account/123/watchlist", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("method = %s, want POST", r.Method)
		}
		if got := r.URL.Query().Get("session_id"); got != "abc123" {
			t.Errorf("session_id query param = %q, want %q", got, "abc123")
		}
		var body str.AccountWatchlistRequest
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			t.Fatalf("decode body: %v", err)
		}
		if body.MediaType != "movie" || body.MediaID != 550 || !body.Watchlist {
			t.Errorf("body = %+v, want MediaType=movie MediaID=550 Watchlist=true", body)
		}
		json.NewEncoder(w).Encode(map[string]any{
			"success":        true,
			"status_code":    1,
			"status_message": "Success.",
		})
	})

	status, _, err := client.Account.AddToWatchlist(context.Background(), 123, "abc123", &str.AccountWatchlistRequest{
		MediaType: "movie",
		MediaID:   550,
		Watchlist: true,
	})
	if err != nil {
		t.Fatalf("AddToWatchlist() error = %v", err)
	}
	if !status.Success {
		t.Errorf("AddToWatchlist() Success = %v, want true", status.Success)
	}
}

func TestAddRemoveFavorite(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/account/123/favorite", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("method = %s, want POST", r.Method)
		}
		if got := r.URL.Query().Get("session_id"); got != "abc123" {
			t.Errorf("session_id query param = %q, want %q", got, "abc123")
		}
		var body str.AccountFavoriteRequest
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			t.Fatalf("decode body: %v", err)
		}
		if body.MediaType != "movie" || body.MediaID != 550 || !body.Favorite {
			t.Errorf("body = %+v, want MediaType=movie MediaID=550 Favorite=true", body)
		}
		json.NewEncoder(w).Encode(map[string]any{
			"success":        true,
			"status_code":    1,
			"status_message": "Success.",
		})
	})

	status, _, err := client.Account.AddRemoveFavorite(context.Background(), 123, "abc123", &str.AccountFavoriteRequest{
		MediaType: "movie",
		MediaID:   550,
		Favorite:  true,
	})
	if err != nil {
		t.Fatalf("AddRemoveFavorite() error = %v", err)
	}
	if !status.Success {
		t.Errorf("AddRemoveFavorite() Success = %v, want true", status.Success)
	}
}

func TestGetFavoriteMovies(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/account/123/favorite/movies", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		if got := r.URL.Query().Get("session_id"); got != "abc123" {
			t.Errorf("session_id query param = %q, want %q", got, "abc123")
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

	movies, err := client.Account.GetFavoriteMovies(context.Background(), 123, "abc123", 0)
	if err != nil {
		t.Fatalf("GetFavoriteMovies() error = %v", err)
	}
	if len(movies) != 1 || movies[0].ID != 550 {
		t.Errorf("GetFavoriteMovies() = %+v, want one movie with ID=550", movies)
	}
}

func TestGetFavoriteTV(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/account/123/favorite/tv", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		if got := r.URL.Query().Get("session_id"); got != "abc123" {
			t.Errorf("session_id query param = %q, want %q", got, "abc123")
		}
		json.NewEncoder(w).Encode(map[string]any{
			"page":          1,
			"total_pages":   1,
			"total_results": 1,
			"results": []map[string]any{
				{"id": 1396, "name": "Breaking Bad"},
			},
		})
	})

	shows, err := client.Account.GetFavoriteTV(context.Background(), 123, "abc123", 0)
	if err != nil {
		t.Fatalf("GetFavoriteTV() error = %v", err)
	}
	if len(shows) != 1 || shows[0].ID != 1396 {
		t.Errorf("GetFavoriteTV() = %+v, want one show with ID=1396", shows)
	}
}
