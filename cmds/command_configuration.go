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
	action := flagSet.String("a", "", "action: details, countries, jobs, languages, primary-translations, timezones (required)")
	if err := flagSet.Parse(args); err != nil {
		return err
	}

	var handler handlers.Handler
	switch *action {
	case "":
		return fmt.Errorf("configuration: -a is required (action: details, countries, jobs, languages, primary-translations, timezones)")
	case "details":
		handler = handlers.ConfigurationDetailsHandler{}
	case "countries":
		handler = handlers.ConfigurationCountriesHandler{}
	case "jobs":
		handler = handlers.ConfigurationJobsHandler{}
	case "languages":
		handler = handlers.ConfigurationLanguagesHandler{}
	case "primary-translations":
		handler = handlers.ConfigurationPrimaryTranslationsHandler{}
	case "timezones":
		handler = handlers.ConfigurationTimezonesHandler{}
	default:
		return fmt.Errorf("configuration: unknown action %q", *action)
	}

	result, err := handler.Handle(context.Background(), client)
	if err != nil {
		return err
	}

	return writeResult(fs, config, "configuration", *action, result)
}
