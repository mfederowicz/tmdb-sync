package cmds

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"

	"github.com/mfederowicz/tmdb-sync/cfg"
	"github.com/mfederowicz/tmdb-sync/handlers"
	"github.com/mfederowicz/tmdb-sync/internal"
	"github.com/mfederowicz/tmdb-sync/printer"
	"github.com/mfederowicz/tmdb-sync/str"

	"github.com/spf13/afero"
)

// ConfigurationCmd is the "configuration" module.
var ConfigurationCmd = &Command{
	Name:   "configuration",
	Abbrev: "config",
	Short:  "TMDB API configuration (image base urls/sizes)",
	Exec:   execConfiguration,
}

func execConfiguration(_ afero.Fs, client *internal.Client, _ *cfg.Config, _ *str.Options, args []string) error {
	fs := flag.NewFlagSet("configuration", flag.ContinueOnError)
	action := fs.String("a", "details", "action: details")
	if err := fs.Parse(args); err != nil {
		return err
	}

	var handler handlers.Handler
	switch *action {
	case "details":
		handler = handlers.ConfigurationDetailsHandler{}
	default:
		return fmt.Errorf("configuration: unknown action %q", *action)
	}

	result, err := handler.Handle(context.Background(), client)
	if err != nil {
		return err
	}

	return printJSON(result)
}

func printJSON(v any) error {
	data, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return err
	}
	printer.Println(string(data))
	return nil
}
