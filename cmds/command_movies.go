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

func execMovies(_ afero.Fs, client *internal.Client, _ *cfg.Config, _ *str.Options, args []string) error {
	fs := flag.NewFlagSet("movies", flag.ContinueOnError)
	action := fs.String("a", "popular", "action: details, popular")
	movieID := fs.Int64("i", 0, "movie id, required for -a details")
	page := fs.Int("p", 1, "page number, used by -a popular")
	if err := fs.Parse(args); err != nil {
		return err
	}

	var handler handlers.Handler
	switch *action {
	case "details":
		if *movieID == 0 {
			return fmt.Errorf("movies: -i <movie_id> is required for -a details")
		}
		handler = handlers.MoviesDetailsHandler{MovieID: *movieID}
	case "popular":
		handler = handlers.MoviesPopularHandler{Page: *page}
	default:
		return fmt.Errorf("movies: unknown action %q", *action)
	}

	result, err := handler.Handle(context.Background(), client)
	if err != nil {
		return err
	}

	return printJSON(result)
}
