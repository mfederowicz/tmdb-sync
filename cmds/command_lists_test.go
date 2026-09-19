package cmds

import (
	"net/http"
	"net/url"
	"strings"
	"testing"

	"github.com/mfederowicz/tmdb-sync/cfg"
	"github.com/mfederowicz/tmdb-sync/str"
	"github.com/spf13/afero"
)

func TestExecListsV4_Unauthorized_DoesNotStartV3Login(t *testing.T) {
	client, mux, teardown := setupAccountClient()
	defer teardown()
	client.BaseURLV4, _ = url.Parse(client.BaseURL.String())
	client.UpdateHeaders(map[string]any{"Authorization": "Bearer read"})

	mux.HandleFunc("/list/abc", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`{"status_code":3,"status_message":"Authentication failed","success":false}`))
	})
	refreshCalls := stubRefreshSession(t, func() (*str.Session, error) {
		return &str.Session{SessionID: "new"}, nil
	})

	options := &str.Options{AccessTokenV4: &str.AccessTokenV4{AccessToken: "stale", AccountID: "acc123"}}
	err := execLists(afero.NewMemMapFs(), client, cfg.DefaultConfig(), options, []string{"-v4", "-a", "delete", "-i", "abc"})

	if err == nil || !strings.Contains(err.Error(), "auth -v4 -a login") {
		t.Errorf("error = %v, want a hint to run `auth -v4 -a login`", err)
	}
	if *refreshCalls != 0 {
		t.Errorf("v3 re-login ran %d times for a v4 401, want 0", *refreshCalls)
	}
}
