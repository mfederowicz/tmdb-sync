package internal

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"
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
