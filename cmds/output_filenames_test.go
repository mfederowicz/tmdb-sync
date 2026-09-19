package cmds

import (
	"flag"
	"net/http"
	"reflect"
	"testing"

	"github.com/mfederowicz/tmdb-sync/cfg"
	"github.com/mfederowicz/tmdb-sync/internal"
	"github.com/mfederowicz/tmdb-sync/str"
	"github.com/spf13/afero"
)

func TestSetFlagParams(t *testing.T) {
	newSet := func() *flag.FlagSet {
		fs := flag.NewFlagSet("t", flag.ContinueOnError)
		fs.String("a", "", "")
		fs.String("query", "", "")
		fs.Int("year", 0, "")
		fs.String("language", "", "")
		fs.Bool("include-adult", false, "")
		fs.Int("pages-limit", 10, "")
		return fs
	}

	tests := []struct {
		name string
		args []string
		want []string
	}{
		{"nothing set", nil, nil},
		{"skipped flags only", []string{"-a", "movies", "-query", "x", "-pages-limit", "3"}, nil},
		{"filters in flag-name order", []string{"-year", "2020", "-language", "pl-PL", "-include-adult"}, []string{"include-adult-true", "language-pl-PL", "year-2020"}},
		{"explicit default still counts", []string{"-year", "0"}, []string{"year-0"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := newSet()
			if err := fs.Parse(tt.args); err != nil {
				t.Fatal(err)
			}
			if got := setFlagParams(fs, "a", "query", "pages-limit"); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("setFlagParams() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOutputFilenames_DistinguishRuns(t *testing.T) {
	client, mux, teardown := setupAccountClient()
	defer teardown()
	empty := `{"page":1,"total_pages":1,"total_results":0,"results":[]}`
	for _, path := range []string{"/search/movie", "/discover/movie", "/changes"} {
		mux.HandleFunc(path, func(w http.ResponseWriter, _ *http.Request) { w.Write([]byte(empty)) })
	}
	mux.HandleFunc("/movie/changes", func(w http.ResponseWriter, _ *http.Request) { w.Write([]byte(empty)) })

	config := cfg.DefaultConfig()
	config.OutputDir = "out"
	fs := afero.NewMemMapFs()
	opts := &str.Options{}

	run := func(exec func(afero.Fs, *internal.Client, *cfg.Config, *str.Options, []string) error, args ...string) {
		t.Helper()
		if err := exec(fs, client, config, opts, args); err != nil {
			t.Fatalf("%v: %v", args, err)
		}
	}
	run(execSearch, "-a", "movies", "-query", "matrix")
	run(execSearch, "-a", "movies", "-query", "matrix", "-year", "1999")
	run(execSearch, "-a", "movies", "-query", "Amélie")
	run(execSearch, "-a", "movies", "-query", "千と千尋")
	run(execDiscover, "-a", "movie")
	run(execDiscover, "-a", "movie", "-with-genres", "28")
	run(execDiscover, "-a", "movie", "-with-genres", "35")
	run(execChanges, "-a", "movie", "-start-date", "2026-09-01")

	files, err := afero.ReadDir(fs, "out")
	if err != nil {
		t.Fatal(err)
	}
	got := map[string]bool{}
	for _, f := range files {
		got[f.Name()] = true
	}

	// 8 runs, 8 distinct files: nothing overwrote anything.
	if len(got) != 8 {
		t.Errorf("wrote %d distinct files, want 8: %v", len(got), got)
	}
	// Runs with no extra filters keep their long-standing names.
	for _, name := range []string{"search_movies_query-matrix.json", "discover_movie.json", "search_movies_query-matrix_year-1999.json"} {
		if !got[name] {
			t.Errorf("expected file %q, got %v", name, got)
		}
	}
}
