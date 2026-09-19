package writer

import (
	"strings"
	"testing"

	"github.com/spf13/afero"
)

func TestBuildFilename(t *testing.T) {
	tests := []struct {
		module, action string
		params         []string
		want           string
	}{
		{"movies", "details", []string{"id-550"}, "movies_details_id-550.json"},
		{"movies", "popular", []string{"page-1"}, "movies_popular_page-1.json"},
		{"movies", "popular", []string{"all"}, "movies_popular_all.json"},
		{"configuration", "details", nil, "configuration_details.json"},
	}

	for _, tt := range tests {
		got := BuildFilename(tt.module, tt.action, tt.params...)
		if got != tt.want {
			t.Errorf("BuildFilename(%q, %q, %v) = %q, want %q", tt.module, tt.action, tt.params, got, tt.want)
		}
	}
}

func TestBuildFilename_KeepsExistingNames(t *testing.T) {
	tests := []struct {
		params []string
		want   string
	}{
		{[]string{"query-matrix"}, "search_movies_query-matrix.json"},
		// A space has always become "-"; it must not start adding a hash.
		{[]string{"query-star wars"}, "search_movies_query-star-wars.json"},
		{[]string{"query-star wars", "year-1977"}, "search_movies_query-star-wars_year-1977.json"},
	}
	for _, tt := range tests {
		if got := BuildFilename("search", "movies", tt.params...); got != tt.want {
			t.Errorf("BuildFilename(%v) = %q, want %q", tt.params, got, tt.want)
		}
	}
}

func TestBuildFilename_DistinctForLossyInputs(t *testing.T) {
	queries := []string{"Amélie", "Am lie", "Amelie", "千と千尋の神隠し", "もののけ姫", "Война и мир", "!!!", "???", "a/b", "a-b"}

	seen := map[string]string{}
	for _, q := range queries {
		name := BuildFilename("search", "movies", "query-"+q)
		if other, dup := seen[name]; dup {
			t.Errorf("queries %q and %q both map to %q", other, q, name)
		}
		seen[name] = q
	}

	// Deterministic, so re-running a search overwrites its own file.
	if a, b := BuildFilename("search", "movies", "query-Amélie"), BuildFilename("search", "movies", "query-Amélie"); a != b {
		t.Errorf("filename not stable across calls: %q vs %q", a, b)
	}
	// The readable part is kept where there is one.
	if got := BuildFilename("search", "movies", "query-Amélie"); !strings.HasPrefix(got, "search_movies_query-Am-lie_") {
		t.Errorf("BuildFilename() = %q, want the sanitized name as a prefix", got)
	}
}

func TestBuildFilename_CapsLongNames(t *testing.T) {
	long := strings.Repeat("a", 400)
	a := BuildFilename("search", "movies", "query-"+long+"1")
	b := BuildFilename("search", "movies", "query-"+long+"2")

	if len(a) > 255 {
		t.Errorf("filename is %d bytes, want at most 255", len(a))
	}
	if a == b {
		t.Errorf("long names differing only past the cap collide: %q", a)
	}
}

func TestWriteJSON_NoDir(t *testing.T) {
	fs := afero.NewMemMapFs()

	path, err := WriteJSON(fs, "", "out.json", map[string]int{"a": 1})
	if err != nil {
		t.Fatalf("WriteJSON() error = %v", err)
	}
	if path != "out.json" {
		t.Errorf("path = %q, want %q", path, "out.json")
	}

	data, err := afero.ReadFile(fs, "out.json")
	if err != nil {
		t.Fatalf("ReadFile() error = %v", err)
	}
	if len(data) == 0 {
		t.Fatal("expected non-empty file content")
	}
}

func TestWriteJSON_WithDir(t *testing.T) {
	fs := afero.NewMemMapFs()

	path, err := WriteJSON(fs, "out/dir", "out.json", map[string]int{"a": 1})
	if err != nil {
		t.Fatalf("WriteJSON() error = %v", err)
	}
	if path != "out/dir/out.json" {
		t.Errorf("path = %q, want %q", path, "out/dir/out.json")
	}

	if exists, _ := afero.Exists(fs, "out/dir/out.json"); !exists {
		t.Fatal("expected file to exist under out/dir")
	}
}
