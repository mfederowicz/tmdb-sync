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

// ChangesCmd is the "changes" module.
var ChangesCmd = &Command{
	Name:   "changes",
	Abbrev: "ch",
	Short:  "movie/tv/person change lists",
	Exec:   execChanges,
}

func execChanges(fs afero.Fs, client *internal.Client, config *cfg.Config, _ *str.Options, args []string) error {
	flagSet := flag.NewFlagSet("changes", flag.ContinueOnError)
	action := flagSet.String("a", "", "action: movie, tv, person (required)")
	startDate := flagSet.String("start-date", "", "start date, YYYY-MM-DD (optional, TMDB defaults to 24 hours ago)")
	endDate := flagSet.String("end-date", "", "end date, YYYY-MM-DD (optional, TMDB defaults to now; range is capped at 14 days)")
	pagesLimit := flagSet.Int("pages-limit", config.PagesLimit, "pages limit (default: pages_limit from config, 0 = unlimited)")
	if err := flagSet.Parse(args); err != nil {
		return err
	}

	var handler handlers.Handler
	switch *action {
	case "":
		return fmt.Errorf("changes: -a is required (action: movie, tv, person)")
	case "movie":
		handler = handlers.ChangesMovieHandler{StartDate: *startDate, EndDate: *endDate, PagesLimit: *pagesLimit}
	case "tv":
		handler = handlers.ChangesTVHandler{StartDate: *startDate, EndDate: *endDate, PagesLimit: *pagesLimit}
	case "person":
		handler = handlers.ChangesPersonHandler{StartDate: *startDate, EndDate: *endDate, PagesLimit: *pagesLimit}
	default:
		return fmt.Errorf("changes: unknown action %q", *action)
	}

	result, err := handler.Handle(context.Background(), client)
	if err != nil {
		return err
	}

	return writeResult(fs, config, "changes", *action, result, setFlagParams(flagSet, "a", "pages-limit")...)
}
