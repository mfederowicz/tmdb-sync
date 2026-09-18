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

// WatchProvidersCmd is the "watch-providers" module.
var WatchProvidersCmd = &Command{
	Name:   "watch-providers",
	Abbrev: "wp",
	Short:  "available regions, movie/tv watch provider lists",
	Exec:   execWatchProviders,
}

func execWatchProviders(fs afero.Fs, client *internal.Client, config *cfg.Config, _ *str.Options, args []string) error {
	flagSet := flag.NewFlagSet("watch-providers", flag.ContinueOnError)
	action := flagSet.String("a", "", "action: available-regions (required)")
	language := flagSet.String("language", "", "ISO 639-1 language code, used by -a available-regions")
	if err := flagSet.Parse(args); err != nil {
		return err
	}

	var handler handlers.Handler
	switch *action {
	case "":
		return fmt.Errorf("watch-providers: -a is required (action: available-regions)")
	case "available-regions":
		handler = handlers.WatchProvidersAvailableRegionsHandler{Language: *language}
	default:
		return fmt.Errorf("watch-providers: unknown action %q", *action)
	}

	result, err := handler.Handle(context.Background(), client)
	if err != nil {
		return err
	}

	return writeResult(fs, config, "watch-providers", *action, result)
}
