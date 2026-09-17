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

// tvSeasonsActionsHelp lists every tv-seasons action, shared between the -a
// flag's usage string and the "-a is required" error so both stay in sync.
const tvSeasonsActionsHelp = "details"

// TVSeasonsCmd is the "tv-seasons" module.
var TVSeasonsCmd = &Command{
	Name:   "tv-seasons",
	Abbrev: "tv-seasons",
	Short:  "TV season details and more",
	Exec:   execTVSeasons,
}

func execTVSeasons(fs afero.Fs, client *internal.Client, config *cfg.Config, _ *str.Options, args []string) error {
	flagSet := flag.NewFlagSet("tv-seasons", flag.ContinueOnError)
	action := flagSet.String("a", "", "action: "+tvSeasonsActionsHelp+" (required)")
	seriesID := flagSet.Int64("i", 0, "tv series id, required for -a "+tvSeasonsActionsHelp)
	seasonNumber := flagSet.Int("s", 0, "season number, required for -a "+tvSeasonsActionsHelp)
	if err := flagSet.Parse(args); err != nil {
		return err
	}

	var handler handlers.Handler
	var params []string
	switch *action {
	case "":
		return fmt.Errorf("tv-seasons: -a is required (action: %s)", tvSeasonsActionsHelp)
	case "details":
		if *seriesID == 0 {
			return fmt.Errorf("tv-seasons: -i <series_id> is required for -a details")
		}
		handler = handlers.TVSeasonsDetailsHandler{SeriesID: *seriesID, SeasonNumber: *seasonNumber}
		params = []string{fmt.Sprintf("id-%d", *seriesID), fmt.Sprintf("season-%d", *seasonNumber)}
	default:
		return fmt.Errorf("tv-seasons: unknown action %q (action: %s)", *action, tvSeasonsActionsHelp)
	}

	result, err := handler.Handle(context.Background(), client)
	if err != nil {
		return fmt.Errorf("tv-seasons: %w", err)
	}

	return writeResult(fs, config, "tv-seasons", *action, result, params...)
}
