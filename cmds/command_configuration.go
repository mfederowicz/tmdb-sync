package cmds

import (
	"context"
	"flag"
	"fmt"

	"github.com/mfederowicz/tmdb-sync/cfg"
	"github.com/mfederowicz/tmdb-sync/handlers"
	"github.com/mfederowicz/tmdb-sync/internal"
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

func execConfiguration(fs afero.Fs, client *internal.Client, config *cfg.Config, _ *str.Options, args []string) error {
	flagSet := flag.NewFlagSet("configuration", flag.ContinueOnError)
	action := flagSet.String("a", "details", "action: details")
	if err := flagSet.Parse(args); err != nil {
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

	return writeResult(fs, config, "configuration", *action, result)
}
