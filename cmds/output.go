package cmds

import (
	"flag"
	"fmt"

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

// setFlagParams returns "name-value" filename parts for the flags the user
// explicitly set, in flag-name order, so runs that differ only by a filter
// (e.g. -year or -with-genres) write different files. Flags left at their
// default add nothing, which keeps existing filenames unchanged. Pass the
// flags that don't identify a result (or are already part of the name) to skip.
func setFlagParams(flagSet *flag.FlagSet, skip ...string) []string {
	skipped := map[string]bool{}
	for _, name := range skip {
		skipped[name] = true
	}

	var params []string
	flagSet.Visit(func(f *flag.Flag) {
		if !skipped[f.Name] {
			params = append(params, fmt.Sprintf("%s-%s", f.Name, f.Value))
		}
	})
	return params
}
