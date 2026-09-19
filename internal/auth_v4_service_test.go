package internal

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"testing"
)

func setupV4() (*Client, *http.ServeMux, func()) {
	client, mux, teardown := setup()
	client.BaseURLV4, _ = url.Parse(client.BaseURL.String())
	return client, mux, teardown
}

func TestAuthV4CreateRequestToken(t *testing.T) {
	client, mux, teardown := setupV4()
	defer teardown()
	client.UpdateHeaders(map[string]any{"Authorization": "Bearer read", APIKeyParam: "mykey"})

	mux.HandleFunc("/auth/request_token", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("method = %s, want POST", r.Method)
		}
		if got := r.Header.Get("Authorization"); got != "Bearer read" {
			t.Errorf("Authorization = %q, want bearer token", got)
		}
		if r.URL.Query().Has(APIKeyParam) {
			t.Errorf("v4 request must not carry api_key, got query %q", r.URL.RawQuery)
		}
		var body struct {
			RedirectTo string `json:"redirect_to"`
		}
		json.NewDecoder(r.Body).Decode(&body)
		if body.RedirectTo != "http://localhost/cb" {
			t.Errorf("redirect_to = %q, want %q", body.RedirectTo, "http://localhost/cb")
		}
		json.NewEncoder(w).Encode(map[string]any{
			"success":        true,
			"status_code":    1,
			"request_token":  "v4tok",
			"status_message": "Success.",
		})
	})

	token, _, err := client.AuthV4.CreateRequestToken(context.Background(), "http://localhost/cb")
	if err != nil {
		t.Fatalf("CreateRequestToken() error = %v", err)
	}
	if token.RequestToken != "v4tok" || !token.Success {
		t.Errorf("token = %+v, want success with v4tok", token)
	}
}

func TestAuthV4CreateRequestToken_NoBodyWithoutRedirect(t *testing.T) {
	client, mux, teardown := setupV4()
	defer teardown()
	client.UpdateHeaders(map[string]any{"Authorization": "Bearer read"})

	mux.HandleFunc("/auth/request_token", func(w http.ResponseWriter, r *http.Request) {
		if r.ContentLength > 0 {
			t.Errorf("expected empty body, got %d bytes", r.ContentLength)
		}
		w.Write([]byte(`{"success":true,"request_token":"x"}`))
	})

	if _, _, err := client.AuthV4.CreateRequestToken(context.Background(), ""); err != nil {
		t.Fatalf("CreateRequestToken() error = %v", err)
	}
}

func TestAuthV4RequiresReadAccessToken(t *testing.T) {
	client, _, teardown := setupV4()
	defer teardown()
	client.UpdateHeaders(map[string]any{APIKeyParam: "mykey"})

	if _, _, err := client.AuthV4.CreateRequestToken(context.Background(), ""); err == nil {
		t.Fatal("expected error when only api_key is configured")
	}
}
