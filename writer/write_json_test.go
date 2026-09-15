package writer

import (
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
