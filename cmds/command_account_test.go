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

func TestExecAccountV4_Lists(t *testing.T) {
	client, mux, teardown := setupAccountClient()
	defer teardown()
	client.BaseURLV4, _ = url.Parse(client.BaseURL.String())
	client.UpdateHeaders(map[string]any{"Authorization": "Bearer read"})

	mux.HandleFunc("/account/acc123/lists", func(w http.ResponseWriter, r *http.Request) {
		if got := r.Header.Get("Authorization"); got != "Bearer usertok" {
			t.Errorf("Authorization = %q, want user access token", got)
		}
		w.Write([]byte(`{"page":1,"total_pages":1,"total_results":1,"results":[{"id":7,"name":"mine"}]}`))
	})

	fs := afero.NewMemMapFs()
	config := cfg.DefaultConfig()
	config.OutputDir = "out"
	options := &str.Options{AccessTokenV4: &str.AccessTokenV4{Success: true, AccessToken: "usertok", AccountID: "acc123"}}

	if err := execAccount(fs, client, config, options, []string{"-v4", "-a", "lists"}); err != nil {
		t.Fatalf("execAccount() error = %v", err)
	}
	if exists, _ := afero.Exists(fs, "out/account_lists_v4.json"); !exists {
		t.Error("expected out/account_lists_v4.json to be written")
	}
}

func TestExecAccountV4_Errors(t *testing.T) {
	client, _, teardown := setupAccountClient()
	defer teardown()

	loggedIn := &str.Options{AccessTokenV4: &str.AccessTokenV4{Success: true, AccessToken: "usertok", AccountID: "acc123"}}
	tests := []struct {
		name    string
		options *str.Options
		args    []string
	}{
		{"no cached token", &str.Options{}, []string{"-v4", "-a", "lists"}},
		{"missing action", loggedIn, []string{"-v4"}},
		{"unknown action", loggedIn, []string{"-v4", "-a", "details"}},
		{"v3-only action", loggedIn, []string{"-v4", "-a", "rated-tv-episodes"}},
		{"v4-only action without -v4", loggedIn, []string{"-a", "recommended-movies"}},
		{"both versions", loggedIn, []string{"-v3", "-v4", "-a", "lists"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := execAccount(afero.NewMemMapFs(), client, cfg.DefaultConfig(), tt.options, tt.args)
			if err == nil {
				t.Error("execAccount() error = nil, want error")
			}
		})
	}
}
