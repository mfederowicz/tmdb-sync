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

// DiscoverCmd is the "discover" module.
var DiscoverCmd = &Command{
	Name:   "discover",
	Abbrev: "d",
	Short:  "discover movies and TV shows by filter",
	Exec:   execDiscover,
}

func execDiscover(fs afero.Fs, client *internal.Client, config *cfg.Config, _ *str.Options, args []string) error {
	flagSet := flag.NewFlagSet("discover", flag.ContinueOnError)
	action := flagSet.String("a", "", "action: movie, tv (required)")
	sortBy := flagSet.String("sort-by", "", "sort order, e.g. popularity.desc")
	language := flagSet.String("language", "", "ISO 639-1 language code")
	region := flagSet.String("region", "", "ISO 3166-1 region code (movie only)")
	includeAdult := flagSet.Bool("include-adult", false, "include adult results")
	year := flagSet.Int("year", 0, "primary release year (movie) or first air date year (tv)")
	withGenres := flagSet.String("with-genres", "", "comma-separated genre ids")
	voteAverageGte := flagSet.Float64("vote-average-gte", 0, "minimum vote average")
	voteAverageLte := flagSet.Float64("vote-average-lte", 0, "maximum vote average")
	withWatchProviders := flagSet.String("with-watch-providers", "", "comma-separated watch provider ids")
	watchRegion := flagSet.String("watch-region", "", "ISO 3166-1 region code for watch providers")
	pagesLimit := flagSet.Int("pages-limit", config.PagesLimit, "pages limit (default: pages_limit from config, 0 = unlimited)")
	if err := flagSet.Parse(args); err != nil {
		return err
	}

	var handler handlers.Handler
	switch *action {
	case "":
		return fmt.Errorf("discover: -a is required (action: movie, tv)")
	case "movie":
		handler = handlers.DiscoverMovieHandler{
			Options: uri.DiscoverMovieOptions{
				SortBy:             *sortBy,
				Language:           *language,
				Region:             *region,
				IncludeAdult:       *includeAdult,
				PrimaryReleaseYear: *year,
				WithGenres:         *withGenres,
				VoteAverageGTE:     *voteAverageGte,
				VoteAverageLTE:     *voteAverageLte,
				WithWatchProviders: *withWatchProviders,
				WatchRegion:        *watchRegion,
			},
			PagesLimit: *pagesLimit,
		}
	case "tv":
		handler = handlers.DiscoverTVHandler{
			Options: uri.DiscoverTVOptions{
				SortBy:             *sortBy,
				Language:           *language,
				IncludeAdult:       *includeAdult,
				FirstAirDateYear:   *year,
				WithGenres:         *withGenres,
				VoteAverageGTE:     *voteAverageGte,
				VoteAverageLTE:     *voteAverageLte,
				WithWatchProviders: *withWatchProviders,
				WatchRegion:        *watchRegion,
			},
			PagesLimit: *pagesLimit,
		}
	default:
		return fmt.Errorf("discover: unknown action %q", *action)
	}

	result, err := handler.Handle(context.Background(), client)
	if err != nil {
		return err
	}

	return writeResult(fs, config, "discover", *action, result, setFlagParams(flagSet, "a", "pages-limit")...)
}
