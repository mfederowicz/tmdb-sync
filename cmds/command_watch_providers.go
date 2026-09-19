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
	action := flagSet.String("a", "", "action: available-regions, movie-providers, tv-providers (required)")
	language := flagSet.String("language", "", "ISO 639-1 language code, used by -a available-regions, movie-providers, tv-providers")
	watchRegion := flagSet.String("watch-region", "", "ISO 3166-1 region code, used by -a movie-providers, tv-providers")
	if err := flagSet.Parse(args); err != nil {
		return err
	}

	var handler handlers.Handler
	switch *action {
	case "":
		return fmt.Errorf("watch-providers: -a is required (action: available-regions, movie-providers, tv-providers)")
	case "available-regions":
		handler = handlers.WatchProvidersAvailableRegionsHandler{Language: *language}
	case "movie-providers":
		handler = handlers.WatchProvidersMovieProvidersHandler{Language: *language, WatchRegion: *watchRegion}
	case "tv-providers":
		handler = handlers.WatchProvidersTVProvidersHandler{Language: *language, WatchRegion: *watchRegion}
	default:
		return fmt.Errorf("watch-providers: unknown action %q", *action)
	}

	result, err := handler.Handle(context.Background(), client)
	if err != nil {
		return err
	}

	return writeResult(fs, config, "watch-providers", *action, result, setFlagParams(flagSet, "a")...)
}
