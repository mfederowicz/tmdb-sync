package internal

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"time"

	"github.com/mfederowicz/tmdb-sync/str"
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

	mux.HandleFunc("/movie/1", func(w http.ResponseWriter, _ *http.Request) {
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

	mux.HandleFunc("/movie/999", func(w http.ResponseWriter, _ *http.Request) {
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

func TestDo_NonJSONErrorKeepsStatusCode(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/movie/1", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("<html>denied</html>"))
	})

	req, _ := client.NewRequest(http.MethodGet, "movie/1", nil)
	_, err := client.Do(context.Background(), req, nil)

	var errResp *str.ErrorResponse
	if !errors.As(err, &errResp) {
		t.Fatalf("expected *str.ErrorResponse, got %T: %v", err, err)
	}
	if errResp.StatusCode != http.StatusUnauthorized {
		t.Errorf("StatusCode = %d, want %d", errResp.StatusCode, http.StatusUnauthorized)
	}
}

func TestDo_RateLimitErrorHidesAPIKey(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()
	client.UpdateHeaders(map[string]any{APIKeyParam: "supersecret"})

	mux.HandleFunc("/movie/1", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusTooManyRequests)
	})

	req, _ := client.NewRequest(http.MethodGet, "movie/1", nil)
	_, err := client.Do(context.Background(), req, nil)
	if err == nil {
		t.Fatal("expected error for 429 response")
	}
	if strings.Contains(err.Error(), "supersecret") {
		t.Errorf("error leaks the api key: %v", err)
	}
}

func TestDo_RetryAfterArmsRateLimitGuard(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	calls := 0
	mux.HandleFunc("/movie/1", func(w http.ResponseWriter, _ *http.Request) {
		calls++
		w.Header().Set("Retry-After", "30")
		w.WriteHeader(http.StatusTooManyRequests)
	})

	req, _ := client.NewRequest(http.MethodGet, "movie/1", nil)
	if _, err := client.Do(context.Background(), req, nil); err == nil {
		t.Fatal("expected error for 429 response")
	}

	req, _ = client.NewRequest(http.MethodGet, "movie/1", nil)
	_, err := client.Do(context.Background(), req, nil)
	var rateErr *AbuseRateLimitError
	if !errors.As(err, &rateErr) {
		t.Fatalf("expected *AbuseRateLimitError, got %T: %v", err, err)
	}
	if calls != 1 {
		t.Errorf("server hit %d times, want 1 (second call should short-circuit)", calls)
	}
}

func TestDo_IgnoresRetryAfterWithoutHeader(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/movie/1", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusTooManyRequests)
	})

	req, _ := client.NewRequest(http.MethodGet, "movie/1", nil)
	client.Do(context.Background(), req, nil)

	client.rateMu.Lock()
	defer client.rateMu.Unlock()
	if !client.RateLimitReset.IsZero() {
		t.Errorf("RateLimitReset = %v, want zero without a Retry-After header", client.RateLimitReset)
	}
}

func TestNewRequestV4_UserTokenWithoutReadToken(t *testing.T) {
	client, _, teardown := setup()
	defer teardown()
	client.BaseURLV4, _ = url.Parse("https://example.test/4/")

	if _, err := client.NewRequestV4(http.MethodGet, "list/1", nil); !errors.Is(err, errV4NeedsReadToken) {
		t.Fatalf("without any token: err = %v, want errV4NeedsReadToken", err)
	}

	req, err := client.NewRequestV4(http.MethodGet, "list/1", nil, withUserToken("usertoken"))
	if err != nil {
		t.Fatalf("with user token: unexpected error %v", err)
	}
	if got := req.Header.Get("Authorization"); got != "Bearer usertoken" {
		t.Errorf("Authorization = %q, want %q", got, "Bearer usertoken")
	}
}
