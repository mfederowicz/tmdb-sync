package internal

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"
)

func setup() (*Client, *http.ServeMux, func()) {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)

	client := NewClient(nil)
	client.BaseURL, _ = url.Parse(server.URL + "/")

	return client, mux, server.Close
}

func TestNewRequest_APIKeyQueryParam(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()
	client.UpdateHeaders(map[string]any{APIKeyParam: "mykey"})

	var gotQuery string
	mux.HandleFunc("/movie/1", func(w http.ResponseWriter, r *http.Request) {
		gotQuery = r.URL.Query().Get(APIKeyParam)
		w.Write([]byte(`{}`))
	})

	req, err := client.NewRequest(http.MethodGet, "movie/1", nil)
	if err != nil {
		t.Fatalf("NewRequest() error = %v", err)
	}
	if _, err := client.Do(context.Background(), req, nil); err != nil {
		t.Fatalf("Do() error = %v", err)
	}
	if gotQuery != "mykey" {
		t.Errorf("api_key query param = %q, want %q", gotQuery, "mykey")
	}
}

func TestNewRequest_BearerHeader(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()
	client.UpdateHeaders(map[string]any{"Authorization": "Bearer mytoken"})

	var gotAuth string
	mux.HandleFunc("/movie/1", func(w http.ResponseWriter, r *http.Request) {
		gotAuth = r.Header.Get("Authorization")
		w.Write([]byte(`{}`))
	})

	req, err := client.NewRequest(http.MethodGet, "movie/1", nil)
	if err != nil {
		t.Fatalf("NewRequest() error = %v", err)
	}
	if _, err := client.Do(context.Background(), req, nil); err != nil {
		t.Fatalf("Do() error = %v", err)
	}
	if gotAuth != "Bearer mytoken" {
		t.Errorf("Authorization header = %q, want %q", gotAuth, "Bearer mytoken")
	}
}

func TestDo_DecodesJSON(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/movie/1", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]any{"id": float64(1), "title": "Some Movie"})
	})

	req, _ := client.NewRequest(http.MethodGet, "movie/1", nil)
	var out struct {
		ID    int    `json:"id"`
		Title string `json:"title"`
	}
	if _, err := client.Do(context.Background(), req, &out); err != nil {
		t.Fatalf("Do() error = %v", err)
	}
	if out.Title != "Some Movie" {
		t.Errorf("Title = %q, want %q", out.Title, "Some Movie")
	}
}

func TestDo_ErrorResponse(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/movie/999", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]any{
			"status_code":    34,
			"status_message": "The resource you requested could not be found.",
			"success":        false,
		})
	})

	req, _ := client.NewRequest(http.MethodGet, "movie/999", nil)
	_, err := client.Do(context.Background(), req, nil)
	if err == nil {
		t.Fatal("expected error for 404 response")
	}
}

func TestCheckRetryAfter(t *testing.T) {
	client, _, teardown := setup()
	defer teardown()

	client.rateMu.Lock()
	client.RateLimitReset = time.Now().Add(time.Minute)
	client.rateMu.Unlock()

	req, _ := client.NewRequest(http.MethodGet, "movie/1", nil)
	if rateErr := client.CheckRetryAfter(req); rateErr == nil {
		t.Fatal("expected CheckRetryAfter to return an error while reset is in the future")
	}
}

func TestCheckRetryAfter_NoLimit(t *testing.T) {
	client, _, teardown := setup()
	defer teardown()

	req, _ := client.NewRequest(http.MethodGet, "movie/1", nil)
	if rateErr := client.CheckRetryAfter(req); rateErr != nil {
		t.Fatalf("expected no error, got %v", rateErr)
	}
}
