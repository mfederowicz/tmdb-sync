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

// SearchCmd is the "search" module.
var SearchCmd = &Command{
	Name:   "search",
	Abbrev: "sr",
	Short:  "search TMDB for collections, companies, keywords, movies, people, and TV shows",
	Exec:   execSearch,
}

func execSearch(fs afero.Fs, client *internal.Client, config *cfg.Config, _ *str.Options, args []string) error {
	flagSet := flag.NewFlagSet("search", flag.ContinueOnError)
	action := flagSet.String("a", "", "action: collections, companies, keywords, movies, multi (required)")
	query := flagSet.String("query", "", "search query, required for all actions")
	language := flagSet.String("language", "", "ISO 639-1 language code, used by -a collections, movies")
	region := flagSet.String("region", "", "ISO 3166-1 region code, used by -a collections, movies")
	includeAdult := flagSet.Bool("include-adult", false, "include adult results, used by -a collections, movies")
	year := flagSet.Int("year", 0, "release year, used by -a movies")
	primaryReleaseYear := flagSet.Int("primary-release-year", 0, "primary release year, used by -a movies")
	pagesLimit := flagSet.Int("pages-limit", config.PagesLimit, "pages limit, used by -a collections, companies, keywords, movies, multi (default: pages_limit from config, 0 = unlimited)")
	if err := flagSet.Parse(args); err != nil {
		return err
	}

	var handler handlers.Handler
	switch *action {
	case "":
		return fmt.Errorf("search: -a is required (action: collections, companies, keywords, movies, multi)")
	case "collections":
		if *query == "" {
			return fmt.Errorf("search: -query is required for -a collections")
		}
		handler = handlers.SearchCollectionsHandler{
			Query:        *query,
			Language:     *language,
			Region:       *region,
			IncludeAdult: *includeAdult,
			PagesLimit:   *pagesLimit,
		}
	case "companies":
		if *query == "" {
			return fmt.Errorf("search: -query is required for -a companies")
		}
		handler = handlers.SearchCompaniesHandler{
			Query:      *query,
			PagesLimit: *pagesLimit,
		}
	case "keywords":
		if *query == "" {
			return fmt.Errorf("search: -query is required for -a keywords")
		}
		handler = handlers.SearchKeywordsHandler{
			Query:      *query,
			PagesLimit: *pagesLimit,
		}
	case "movies":
		if *query == "" {
			return fmt.Errorf("search: -query is required for -a movies")
		}
		handler = handlers.SearchMoviesHandler{
			Query:              *query,
			Language:           *language,
			Region:             *region,
			IncludeAdult:       *includeAdult,
			Year:               *year,
			PrimaryReleaseYear: *primaryReleaseYear,
			PagesLimit:         *pagesLimit,
		}
	case "multi":
		if *query == "" {
			return fmt.Errorf("search: -query is required for -a multi")
		}
		handler = handlers.SearchMultiHandler{
			Query:        *query,
			Language:     *language,
			IncludeAdult: *includeAdult,
			PagesLimit:   *pagesLimit,
		}
	default:
		return fmt.Errorf("search: unknown action %q", *action)
	}

	result, err := handler.Handle(context.Background(), client)
	if err != nil {
		return err
	}

	params := []string{fmt.Sprintf("query-%s", *query)}
	return writeResult(fs, config, "search", *action, result, params...)
}
