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
	action := flagSet.String("a", "", "action: collections, companies (required)")
	query := flagSet.String("query", "", "search query, required for all actions")
	language := flagSet.String("language", "", "ISO 639-1 language code, used by -a collections")
	region := flagSet.String("region", "", "ISO 3166-1 region code, used by -a collections")
	includeAdult := flagSet.Bool("include-adult", false, "include adult results, used by -a collections")
	pagesLimit := flagSet.Int("pages-limit", config.PagesLimit, "pages limit, used by -a collections (default: pages_limit from config, 0 = unlimited)")
	if err := flagSet.Parse(args); err != nil {
		return err
	}

	var handler handlers.Handler
	switch *action {
	case "":
		return fmt.Errorf("search: -a is required (action: collections, companies)")
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
