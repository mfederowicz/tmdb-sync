package cmds

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/mfederowicz/tmdb-sync/cfg"
	"github.com/mfederowicz/tmdb-sync/internal"
	"github.com/mfederowicz/tmdb-sync/str"
	"github.com/spf13/afero"
)

func setupAccountClient() (*internal.Client, *http.ServeMux, func()) {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)

	client := internal.NewClient(nil)
	client.BaseURL, _ = url.Parse(server.URL + "/")

	return client, mux, server.Close
}

func TestIsSessionInvalid(t *testing.T) {
	tests := []struct {
		name string
		err  error
		want bool
	}{
		{"401 error response", &str.ErrorResponse{StatusCode: http.StatusUnauthorized}, true},
		{"404 error response", &str.ErrorResponse{StatusCode: http.StatusNotFound}, false},
		{"plain error", errors.New("boom"), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSessionInvalid(tt.err); got != tt.want {
				t.Errorf("isSessionInvalid(%v) = %v, want %v", tt.err, got, tt.want)
			}
		})
	}
}

func TestExecAccount_RefreshesStaleSession(t *testing.T) {
	client, mux, teardown := setupAccountClient()
	defer teardown()

	hits := 0
	mux.HandleFunc("/account/123/rated/tv", func(w http.ResponseWriter, r *http.Request) {
		hits++
		sessionID := r.URL.Query().Get("session_id")
		if hits == 1 {
			if sessionID != "old-session" {
				t.Errorf("first hit session_id = %q, want %q", sessionID, "old-session")
			}
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(`{"status_code":3,"status_message":"Authentication failed","success":false}`))
			return
		}
		if sessionID != "new-session" {
			t.Errorf("second hit session_id = %q, want %q", sessionID, "new-session")
		}
		w.Write([]byte(`{"page":1,"total_pages":1,"total_results":0,"results":[]}`))
	})

	old := refreshSession
	refreshSession = func(afero.Fs, *cfg.Config, *internal.Client) (*str.Session, error) {
		return &str.Session{Success: true, SessionID: "new-session"}, nil
	}
	defer func() { refreshSession = old }()

	fs := afero.NewMemMapFs()
	config := cfg.DefaultConfig()
	options := &str.Options{
		Session: &str.Session{Success: true, SessionID: "old-session"},
		Account: &str.Account{ID: 123},
	}

	if err := execAccount(fs, client, config, options, []string{"-a", "rated-tv"}); err != nil {
		t.Fatalf("execAccount() error = %v", err)
	}

	if hits != 2 {
		t.Errorf("mux hit %d times, want 2", hits)
	}
	if options.Session.SessionID != "new-session" {
		t.Errorf("options.Session.SessionID = %q, want %q", options.Session.SessionID, "new-session")
	}
}
