package cmds

import (
	"github.com/mfederowicz/tmdb-sync/cfg"
	"github.com/mfederowicz/tmdb-sync/printer"
	"github.com/mfederowicz/tmdb-sync/writer"

	"github.com/spf13/afero"
)

// writeResult saves a handler's result to a JSON file named after the
// module/action/params (see writer.BuildFilename) under config.OutputDir,
// and prints a short confirmation instead of dumping the raw JSON to stdout.
func writeResult(fs afero.Fs, config *cfg.Config, module, action string, v any, params ...string) error {
	filename := writer.BuildFilename(module, action, params...)

	path, err := writer.WriteJSON(fs, config.OutputDir, filename, v)
	if err != nil {
		return err
	}

	printer.Println("wrote", path)
	return nil
}
