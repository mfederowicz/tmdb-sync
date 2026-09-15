// Package writer saves command output to disk as JSON.
package writer

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"

	"github.com/spf13/afero"
)

var unsafeFilenameChars = regexp.MustCompile(`[^a-zA-Z0-9_.-]+`)

// BuildFilename builds an output filename from a module name, an action name,
// and optional "key-value" parameter parts (e.g. "id-550", "page-2"), e.g.
// BuildFilename("movies", "details", "id-550") -> "movies_details_id-550.json".
func BuildFilename(module, action string, params ...string) string {
	parts := append([]string{module, action}, params...)
	name := strings.Join(parts, "_")
	name = unsafeFilenameChars.ReplaceAllString(name, "-")
	return name + ".json"
}

// WriteJSON marshals v as indented JSON and writes it to dir/filename
// (dir may be "" for the current directory), creating dir if needed.
func WriteJSON(fs afero.Fs, dir, filename string, v any) (string, error) {
	data, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return "", fmt.Errorf("marshal output: %w", err)
	}

	path := filename
	if dir != "" {
		if err := fs.MkdirAll(dir, 0o755); err != nil {
			return "", fmt.Errorf("create output dir %q: %w", dir, err)
		}
		path = dir + "/" + filename
	}

	if err := afero.WriteFile(fs, path, data, 0o644); err != nil {
		return "", fmt.Errorf("write output file %q: %w", path, err)
	}

	return path, nil
}
