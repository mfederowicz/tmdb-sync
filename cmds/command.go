// Package cmds used for commands modules
package cmds

import (
	"github.com/mfederowicz/tmdb-sync/cfg"
	"github.com/mfederowicz/tmdb-sync/internal"
	"github.com/mfederowicz/tmdb-sync/str"

	"github.com/spf13/afero"
)

// Command represents one top-level module (e.g. "movies", "configuration").
type Command struct {
	Name   string
	Abbrev string
	Short  string
	Exec   func(fs afero.Fs, client *internal.Client, config *cfg.Config, options *str.Options, args []string) error
}
