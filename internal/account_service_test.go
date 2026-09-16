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
