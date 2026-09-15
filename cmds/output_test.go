package cmds

import (
	"testing"

	"github.com/mfederowicz/tmdb-sync/cfg"

	"github.com/spf13/afero"
)

func TestWriteResult_BuildsExpectedPath(t *testing.T) {
	fs := afero.NewMemMapFs()
	config := &cfg.Config{OutputDir: "out"}

	if err := writeResult(fs, config, "movies", "details", map[string]int{"id": 550}, "id-550"); err != nil {
		t.Fatalf("writeResult() error = %v", err)
	}

	exists, err := afero.Exists(fs, "out/movies_details_id-550.json")
	if err != nil {
		t.Fatalf("Exists() error = %v", err)
	}
	if !exists {
		t.Fatal("expected out/movies_details_id-550.json to be written")
	}
}
