package cmds

import (
	"context"
	"flag"
	"fmt"

	"github.com/mfederowicz/tmdb-sync/cfg"
	"github.com/mfederowicz/tmdb-sync/handlers"
	"github.com/mfederowicz/tmdb-sync/internal"
	"github.com/mfederowicz/tmdb-sync/str"
	"github.com/mfederowicz/tmdb-sync/uri"

	"github.com/spf13/afero"
)

// FindCmd is the "find" module.
var FindCmd = &Command{
	Name:   "find",
	Abbrev: "f",
	Short:  "find TMDB items by an external id (IMDb, TVDB, ...)",
	Exec:   execFind,
}

func execFind(fs afero.Fs, client *internal.Client, config *cfg.Config, _ *str.Options, args []string) error {
	flagSet := flag.NewFlagSet("find", flag.ContinueOnError)
	action := flagSet.String("a", "", "action: by-id (required)")
	externalID := flagSet.String("i", "", "external id, required for all actions")
	source := flagSet.String("source", "", "external source, e.g. imdb_id, tvdb_id (required)")
	language := flagSet.String("language", "", "ISO 639-1 language code")
	if err := flagSet.Parse(args); err != nil {
		return err
	}

	var handler handlers.Handler
	switch *action {
	case "":
		return fmt.Errorf("find: -a is required (action: by-id)")
	case "by-id":
		if *externalID == "" {
			return fmt.Errorf("find: -i <external_id> is required for -a by-id")
		}
		if *source == "" {
			return fmt.Errorf("find: --source <external_source> is required for -a by-id")
		}
		handler = handlers.FindByIDHandler{
			ExternalID: *externalID,
			Options: uri.FindOptions{
				ExternalSource: *source,
				Language:       *language,
			},
		}
	default:
		return fmt.Errorf("find: unknown action %q", *action)
	}

	result, err := handler.Handle(context.Background(), client)
	if err != nil {
		return err
	}

	params := []string{fmt.Sprintf("id-%s", *externalID)}
	return writeResult(fs, config, "find", *action, result, params...)
}
