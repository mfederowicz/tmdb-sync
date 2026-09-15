package internal

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"
)

func TestCreateRequestToken(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/authentication/token/new", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		json.NewEncoder(w).Encode(map[string]any{
			"success":       true,
			"expires_at":    "2024-01-01 00:00:00 UTC",
			"request_token": "abc123",
		})
	})

	token, _, err := client.Auth.CreateRequestToken(context.Background())
	if err != nil {
		t.Fatalf("CreateRequestToken() error = %v", err)
	}
	if token.RequestToken != "abc123" {
		t.Errorf("RequestToken = %q, want %q", token.RequestToken, "abc123")
	}
}

func TestCreateSession(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/authentication/session/new", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("method = %s, want POST", r.Method)
		}
		var body struct {
			RequestToken string `json:"request_token"`
		}
		json.NewDecoder(r.Body).Decode(&body)
		if body.RequestToken != "abc123" {
			t.Errorf("request_token in body = %q, want %q", body.RequestToken, "abc123")
		}
		json.NewEncoder(w).Encode(map[string]any{
			"success":    true,
			"session_id": "sess123",
		})
	})

	session, _, err := client.Auth.CreateSession(context.Background(), "abc123")
	if err != nil {
		t.Fatalf("CreateSession() error = %v", err)
	}
	if !session.Valid() {
		t.Fatalf("session not valid: %+v", session)
	}
	if session.SessionID != "sess123" {
		t.Errorf("SessionID = %q, want %q", session.SessionID, "sess123")
	}
}
