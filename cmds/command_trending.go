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

// TrendingCmd is the "trending" module.
var TrendingCmd = &Command{
	Name:   "trending",
	Abbrev: "tr",
	Short:  "list trending movies, tv shows, and people",
	Exec:   execTrending,
}

func execTrending(fs afero.Fs, client *internal.Client, config *cfg.Config, _ *str.Options, args []string) error {
	flagSet := flag.NewFlagSet("trending", flag.ContinueOnError)
	action := flagSet.String("a", "", "action: all, movie, tv, person (required)")
	timeWindow := flagSet.String("w", "day", "time window: day or week, used by all actions")
	pagesLimit := flagSet.Int("pages-limit", config.PagesLimit, "pages limit, used by all actions (default: pages_limit from config, 0 = unlimited)")
	if err := flagSet.Parse(args); err != nil {
		return err
	}

	if *timeWindow != "day" && *timeWindow != "week" {
		return fmt.Errorf("trending: -w must be %q or %q", "day", "week")
	}

	var handler handlers.Handler
	switch *action {
	case "":
		return fmt.Errorf("trending: -a is required (action: all, movie, tv, person)")
	case "all":
		handler = handlers.TrendingAllHandler{
			TimeWindow: *timeWindow,
			PagesLimit: *pagesLimit,
		}
	case "movie":
		handler = handlers.TrendingMovieHandler{
			TimeWindow: *timeWindow,
			PagesLimit: *pagesLimit,
		}
	case "tv":
		handler = handlers.TrendingTVHandler{
			TimeWindow: *timeWindow,
			PagesLimit: *pagesLimit,
		}
	case "person":
		handler = handlers.TrendingPersonHandler{
			TimeWindow: *timeWindow,
			PagesLimit: *pagesLimit,
		}
	default:
		return fmt.Errorf("trending: unknown action %q", *action)
	}

	result, err := handler.Handle(context.Background(), client)
	if err != nil {
		return err
	}

	params := []string{fmt.Sprintf("window-%s", *timeWindow)}
	return writeResult(fs, config, "trending", *action, result, params...)
}
