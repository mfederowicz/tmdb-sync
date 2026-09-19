package cmds

import (
	"net/http"
	"strings"
	"testing"

	"github.com/mfederowicz/tmdb-sync/cfg"
	"github.com/mfederowicz/tmdb-sync/str"
	"github.com/spf13/afero"
)

func TestValidateRatingValue(t *testing.T) {
	tests := []struct {
		name    string
		action  string
		value   float64
		wantErr bool
	}{
		{"other action ignores value", "details", 99, false},
		{"missing", "add-rating", 0, true},
		{"minimum", "add-rating", 0.5, false},
		{"maximum", "add-rating", 10, false},
		{"whole", "add-rating", 7, false},
		{"half step", "add-rating", 7.5, false},
		{"below minimum", "add-rating", 0.3, true},
		{"negative", "add-rating", -3, true},
		{"above maximum", "add-rating", 10.5, true},
		{"way above", "add-rating", 15, true},
		{"not a half step", "add-rating", 7.3, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateRatingValue("movies", tt.action, tt.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("validateRatingValue(%q, %v) error = %v, wantErr %v", tt.action, tt.value, err, tt.wantErr)
			}
		})
	}
}

func TestExecMovies_AddRatingRejectsBadValueBeforeLogin(t *testing.T) {
	client, mux, teardown := setupAccountClient()
	defer teardown()
	hits := 0
	mux.HandleFunc("/", func(http.ResponseWriter, *http.Request) { hits++ })

	// No session and no guest session: an accepted value would start the login flow.
	err := execMovies(afero.NewMemMapFs(), client, cfg.DefaultConfig(), &str.Options{}, []string{"-a", "add-rating", "-i", "550", "-value", "15"})
	if err == nil || !strings.Contains(err.Error(), "-value") {
		t.Errorf("error = %v, want a -value range error", err)
	}
	if hits != 0 {
		t.Errorf("server hit %d times, want 0", hits)
	}
}

func TestExecTVSeasons_RequiresSeasonFlag(t *testing.T) {
	client, mux, teardown := setupAccountClient()
	defer teardown()
	hits := 0
	mux.HandleFunc("/tv/1399/season/0", func(w http.ResponseWriter, _ *http.Request) {
		hits++
		w.Write([]byte(`{"id":1,"season_number":0}`))
	})
	config := cfg.DefaultConfig()
	config.OutputDir = "out"

	err := execTVSeasons(afero.NewMemMapFs(), client, config, &str.Options{}, []string{"-a", "details", "-i", "1399"})
	if err == nil || !strings.Contains(err.Error(), "-s") {
		t.Errorf("without -s: error = %v, want a -s required error", err)
	}
	if hits != 0 {
		t.Errorf("server hit %d times without -s, want 0", hits)
	}

	// An explicit -s 0 asks for the specials season and must still work.
	if err := execTVSeasons(afero.NewMemMapFs(), client, config, &str.Options{}, []string{"-a", "details", "-i", "1399", "-s", "0"}); err != nil {
		t.Fatalf("with -s 0: error = %v", err)
	}
	if hits != 1 {
		t.Errorf("server hit %d times with -s 0, want 1", hits)
	}
}

func TestExecTVEpisodes_RequiresSeasonAndEpisodeFlags(t *testing.T) {
	client, _, teardown := setupAccountClient()
	defer teardown()

	tests := []struct {
		name string
		args []string
		want string
	}{
		{"no -s", []string{"-a", "details", "-i", "1399", "-e", "1"}, "-s"},
		{"no -e", []string{"-a", "details", "-i", "1399", "-s", "1"}, "-e"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := execTVEpisodes(afero.NewMemMapFs(), client, cfg.DefaultConfig(), &str.Options{}, tt.args)
			if err == nil || !strings.Contains(err.Error(), tt.want) {
				t.Errorf("error = %v, want it to mention %s", err, tt.want)
			}
		})
	}
}

func TestExecTVSeasons_UnknownActionKeepsItsOwnError(t *testing.T) {
	client, _, teardown := setupAccountClient()
	defer teardown()

	err := execTVSeasons(afero.NewMemMapFs(), client, cfg.DefaultConfig(), &str.Options{}, []string{"-a", "bogus", "-i", "1"})
	if err == nil || !strings.Contains(err.Error(), "unknown action") {
		t.Errorf("error = %v, want an unknown action error (not a -s error)", err)
	}
}
