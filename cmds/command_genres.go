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

// GenresCmd is the "genres" module.
var GenresCmd = &Command{
	Name:   "genres",
	Abbrev: "genre",
	Short:  "TMDB official genre lists (movie/tv)",
	Exec:   execGenres,
}

func execGenres(fs afero.Fs, client *internal.Client, config *cfg.Config, _ *str.Options, args []string) error {
	flagSet := flag.NewFlagSet("genres", flag.ContinueOnError)
	action := flagSet.String("a", "", "action: movie, tv (required)")
	if err := flagSet.Parse(args); err != nil {
		return err
	}

	var handler handlers.Handler
	switch *action {
	case "":
		return fmt.Errorf("genres: -a is required (action: movie, tv)")
	case "movie":
		handler = handlers.GenreMovieHandler{}
	case "tv":
		handler = handlers.GenreTVHandler{}
	default:
		return fmt.Errorf("genres: unknown action %q", *action)
	}

	result, err := handler.Handle(context.Background(), client)
	if err != nil {
		return err
	}

	return writeResult(fs, config, "genres", *action, result)
}
