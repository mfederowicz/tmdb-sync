package cmds

import (
	"net/http"
	"testing"

	"github.com/mfederowicz/tmdb-sync/cfg"
	"github.com/mfederowicz/tmdb-sync/str"
	"github.com/spf13/afero"
)

func TestExecAuth_V3Logout(t *testing.T) {
	client, mux, teardown := setupAccountClient()
	defer teardown()

	mux.HandleFunc("/authentication/session", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Errorf("method = %s, want DELETE", r.Method)
		}
		w.Write([]byte(`{"success":true}`))
	})

	fs := afero.NewMemMapFs()
	config := cfg.DefaultConfig()
	config.OutputDir = "out"
	config.SessionPath = "cfg/session.json"
	config.AccountPath = "cfg/account.json"
	afero.WriteFile(fs, config.SessionPath, []byte("{}"), 0o600)
	afero.WriteFile(fs, config.AccountPath, []byte("{}"), 0o600)
	options := &str.Options{Session: &str.Session{Success: true, SessionID: "sess123"}}

	if err := execAuth(fs, client, config, options, []string{"-a", "logout"}); err != nil {
		t.Fatalf("execAuth() error = %v", err)
	}
	for _, path := range []string{config.SessionPath, config.AccountPath} {
		if exists, _ := afero.Exists(fs, path); exists {
			t.Errorf("%s still exists after logout", path)
		}
	}
	if exists, _ := afero.Exists(fs, "out/auth_logout_v3.json"); !exists {
		t.Error("expected out/auth_logout_v3.json to be written")
	}
}

func TestExecAuth_V3Errors(t *testing.T) {
	client, _, teardown := setupAccountClient()
	defer teardown()

	tests := []struct {
		name    string
		options *str.Options
		args    []string
	}{
		{"logout without cached session", &str.Options{}, []string{"-a", "logout"}},
		{"login without -v4", &str.Options{}, []string{"-a", "login"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := execAuth(afero.NewMemMapFs(), client, cfg.DefaultConfig(), tt.options, tt.args); err == nil {
				t.Error("execAuth() error = nil, want error")
			}
		})
	}
}
