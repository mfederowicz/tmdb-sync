// Package writer saves command output to disk as JSON.
package writer

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"regexp"
	"strings"

	"github.com/spf13/afero"
)

var unsafeFilenameChars = regexp.MustCompile(`[^a-zA-Z0-9_.-]+`)

// lossyFilenameChars matches characters that sanitizing replaces in a way that
// lets different inputs collide (e.g. "Amélie" and "Am lie", or every all-CJK
// query). A plain space is left out on purpose: it has always become "-", and
// treating it as lossy would rename the files of every multi-word query.
var lossyFilenameChars = regexp.MustCompile(`[^a-zA-Z0-9_.\- ]`)

// maxFilenameStem keeps the name (plus hash suffix and ".json") well inside the
// usual 255-byte filename limit.
const maxFilenameStem = 200

// BuildFilename builds an output filename from a module name, an action name,
// and optional "key-value" parameter parts (e.g. "id-550", "page-2"), e.g.
// BuildFilename("movies", "details", "id-550") -> "movies_details_id-550.json".
//
// When sanitizing would lose information (see lossyFilenameChars) or the name
// is too long, a short hash of the unsanitized input is appended so distinct
// inputs get distinct files. Other names are unchanged.
func BuildFilename(module, action string, params ...string) string {
	raw := strings.Join(append([]string{module, action}, params...), "_")
	stem := unsafeFilenameChars.ReplaceAllString(raw, "-")

	if lossyFilenameChars.MatchString(raw) || len(stem) > maxFilenameStem {
		if len(stem) > maxFilenameStem {
			stem = stem[:maxFilenameStem]
		}
		sum := sha256.Sum256([]byte(raw))
		stem += "_" + hex.EncodeToString(sum[:4])
	}
	return stem + ".json"
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
