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

func TestAuthV4CreateAccessToken(t *testing.T) {
	client, mux, teardown := setupV4()
	defer teardown()
	client.UpdateHeaders(map[string]any{"Authorization": "Bearer read", APIKeyParam: "mykey"})

	mux.HandleFunc("/auth/access_token", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("method = %s, want POST", r.Method)
		}
		if r.URL.Query().Has(APIKeyParam) {
			t.Errorf("v4 request must not carry api_key, got query %q", r.URL.RawQuery)
		}
		var body struct {
			RequestToken string `json:"request_token"`
		}
		json.NewDecoder(r.Body).Decode(&body)
		if body.RequestToken != "approved" {
			t.Errorf("request_token in body = %q, want %q", body.RequestToken, "approved")
		}
		json.NewEncoder(w).Encode(map[string]any{
			"success":      true,
			"status_code":  1,
			"access_token": "user-tok",
			"account_id":   "acc-obj-1",
		})
	})

	got, _, err := client.AuthV4.CreateAccessToken(context.Background(), "approved")
	if err != nil {
		t.Fatalf("CreateAccessToken() error = %v", err)
	}
	if !got.Valid() || got.AccessToken != "user-tok" || got.AccountID != "acc-obj-1" {
		t.Errorf("got %+v, want valid token user-tok / acc-obj-1", got)
	}
}

func TestAuthV4CreateAccessToken_Unapproved(t *testing.T) {
	client, mux, teardown := setupV4()
	defer teardown()
	client.UpdateHeaders(map[string]any{"Authorization": "Bearer read"})

	mux.HandleFunc("/auth/access_token", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]any{"success": false, "status_code": 33, "status_message": "not approved"})
	})

	if _, _, err := client.AuthV4.CreateAccessToken(context.Background(), "nope"); err == nil {
		t.Fatal("expected error for an unapproved request token")
	}
}

func TestAuthV4DeleteAccessToken(t *testing.T) {
	client, mux, teardown := setupV4()
	defer teardown()
	client.UpdateHeaders(map[string]any{"Authorization": "Bearer read", APIKeyParam: "mykey"})

	mux.HandleFunc("/auth/access_token", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Errorf("method = %s, want DELETE", r.Method)
		}
		if got := r.Header.Get("Authorization"); got != "Bearer read" {
			t.Errorf("Authorization = %q, want the read token", got)
		}
		if r.URL.Query().Has(APIKeyParam) {
			t.Errorf("v4 request must not carry api_key, got query %q", r.URL.RawQuery)
		}
		var body struct {
			AccessToken string `json:"access_token"`
		}
		json.NewDecoder(r.Body).Decode(&body)
		if body.AccessToken != "user-tok" {
			t.Errorf("access_token in body = %q, want %q", body.AccessToken, "user-tok")
		}
		json.NewEncoder(w).Encode(map[string]any{"success": true, "status_code": 1, "status_message": "Success."})
	})

	status, _, err := client.AuthV4.DeleteAccessToken(context.Background(), "user-tok")
	if err != nil {
		t.Fatalf("DeleteAccessToken() error = %v", err)
	}
	if !status.Success {
		t.Errorf("status = %+v, want success", status)
	}
}

func TestAuthV4DeleteAccessToken_Invalid(t *testing.T) {
	client, mux, teardown := setupV4()
	defer teardown()
	client.UpdateHeaders(map[string]any{"Authorization": "Bearer read"})

	mux.HandleFunc("/auth/access_token", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]any{"success": false, "status_code": 3, "status_message": "Authentication failed"})
	})

	if _, _, err := client.AuthV4.DeleteAccessToken(context.Background(), "bad"); err == nil {
		t.Fatal("expected error for an invalid access token")
	}
}
