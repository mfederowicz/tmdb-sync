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

// MoviesCmd is the "movies" module.
var MoviesCmd = &Command{
	Name:   "movies",
	Abbrev: "m",
	Short:  "movie details, popular, ...",
	Exec:   execMovies,
}

func execMovies(fs afero.Fs, client *internal.Client, config *cfg.Config, _ *str.Options, args []string) error {
	flagSet := flag.NewFlagSet("movies", flag.ContinueOnError)
	action := flagSet.String("a", "", "action: details, popular (required)")
	movieID := flagSet.Int64("i", 0, "movie id, required for -a details")
	pagesLimit := flagSet.Int("pages-limit", config.PagesLimit, "pages limit, used by -a popular (default: pages_limit from config, 0 = unlimited)")
	if err := flagSet.Parse(args); err != nil {
		return err
	}

	var handler handlers.Handler
	var params []string
	switch *action {
	case "":
		return fmt.Errorf("movies: -a is required (action: details, popular)")
	case "details":
		if *movieID == 0 {
			return fmt.Errorf("movies: -i <movie_id> is required for -a details")
		}
		handler = handlers.MoviesDetailsHandler{MovieID: *movieID}
		params = []string{fmt.Sprintf("id-%d", *movieID)}
	case "popular":
		// Fetches every page up to pagesLimit (0 = unlimited, bounded by
		// TMDB's total_pages) - see internal.FetchAllPages. Defaults to
		// config.PagesLimit, overridable per-invocation via -pages-limit.
		handler = handlers.MoviesPopularHandler{PagesLimit: *pagesLimit}
	default:
		return fmt.Errorf("movies: unknown action %q", *action)
	}

	result, err := handler.Handle(context.Background(), client)
	if err != nil {
		return err
	}

	return writeResult(fs, config, "movies", *action, result, params...)
}
