package cmds

import (
	"context"
	"flag"
	"fmt"

	"github.com/mfederowicz/tmdb-sync/cfg"
	"github.com/mfederowicz/tmdb-sync/cli"
	"github.com/mfederowicz/tmdb-sync/handlers"
	"github.com/mfederowicz/tmdb-sync/internal"
	"github.com/mfederowicz/tmdb-sync/str"

	"github.com/spf13/afero"
)

// moviesSessionActions are the actions that require a v3 session.
var moviesSessionActions = map[string]bool{
	"account-states": true,
}

// MoviesCmd is the "movies" module.
var MoviesCmd = &Command{
	Name:   "movies",
	Abbrev: "m",
	Short:  "movie details, popular, 🔒 account-states, ...",
	Exec:   execMovies,
}

func execMovies(fs afero.Fs, client *internal.Client, config *cfg.Config, options *str.Options, args []string) error {
	return execMoviesAttempt(fs, client, config, options, args, false)
}

// execMoviesAttempt is execMovies's body, split out so a stale session
// (detected via isSessionInvalid) can trigger one transparent re-login and
// retry instead of failing outright.
func execMoviesAttempt(fs afero.Fs, client *internal.Client, config *cfg.Config, options *str.Options, args []string, retried bool) error {
	flagSet := flag.NewFlagSet("movies", flag.ContinueOnError)
	action := flagSet.String("a", "", "action: details, popular, account-states, alternative-titles, credits, external-ids, images, keywords, latest, lists, now-playing, recommendations (required)")
	movieID := flagSet.Int64("i", 0, "movie id, required for -a details, account-states, alternative-titles, credits, external-ids, images, keywords")
	pagesLimit := flagSet.Int("pages-limit", config.PagesLimit, "pages limit, used by -a popular, lists, now-playing, recommendations (default: pages_limit from config, 0 = unlimited)")
	country := flagSet.String("country", "", "ISO 3166-1 country code, used by -a alternative-titles")
	language := flagSet.String("language", "", "ISO 639-1 language code, used by -a credits, images")
	includeImageLanguage := flagSet.String("include-image-language", "", "comma-separated language codes, used by -a images")
	if err := flagSet.Parse(args); err != nil {
		return err
	}

	if moviesSessionActions[*action] {
		if err := cli.HandleToken(fs, config, client, options); err != nil {
			return fmt.Errorf("movies: %w", err)
		}
	}

	var handler handlers.Handler
	var params []string
	switch *action {
	case "":
		return fmt.Errorf("movies: -a is required (action: details, popular, account-states, alternative-titles, credits, external-ids, images, keywords, latest, lists, now-playing, recommendations)")
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
	case "account-states":
		if *movieID == 0 {
			return fmt.Errorf("movies: -i <movie_id> is required for -a account-states")
		}
		handler = handlers.MoviesAccountStatesHandler{MovieID: *movieID, SessionID: options.Session.SessionID}
		params = []string{fmt.Sprintf("id-%d", *movieID)}
	case "alternative-titles":
		if *movieID == 0 {
			return fmt.Errorf("movies: -i <movie_id> is required for -a alternative-titles")
		}
		handler = handlers.MoviesAlternativeTitlesHandler{MovieID: *movieID, Country: *country}
		params = []string{fmt.Sprintf("id-%d", *movieID)}
	case "credits":
		if *movieID == 0 {
			return fmt.Errorf("movies: -i <movie_id> is required for -a credits")
		}
		handler = handlers.MoviesCreditsHandler{MovieID: *movieID, Language: *language}
		params = []string{fmt.Sprintf("id-%d", *movieID)}
	case "external-ids":
		if *movieID == 0 {
			return fmt.Errorf("movies: -i <movie_id> is required for -a external-ids")
		}
		handler = handlers.MoviesExternalIDsHandler{MovieID: *movieID}
		params = []string{fmt.Sprintf("id-%d", *movieID)}
	case "images":
		if *movieID == 0 {
			return fmt.Errorf("movies: -i <movie_id> is required for -a images")
		}
		handler = handlers.MoviesImagesHandler{MovieID: *movieID, Language: *language, IncludeImageLanguage: *includeImageLanguage}
		params = []string{fmt.Sprintf("id-%d", *movieID)}
	case "keywords":
		if *movieID == 0 {
			return fmt.Errorf("movies: -i <movie_id> is required for -a keywords")
		}
		handler = handlers.MoviesKeywordsHandler{MovieID: *movieID}
		params = []string{fmt.Sprintf("id-%d", *movieID)}
	case "latest":
		handler = handlers.MoviesLatestHandler{}
	case "lists":
		if *movieID == 0 {
			return fmt.Errorf("movies: -i <movie_id> is required for -a lists")
		}
		handler = handlers.MoviesListsHandler{MovieID: *movieID, Language: *language, PagesLimit: *pagesLimit}
		params = []string{fmt.Sprintf("id-%d", *movieID)}
	case "now-playing":
		handler = handlers.MoviesNowPlayingHandler{PagesLimit: *pagesLimit}
	case "recommendations":
		if *movieID == 0 {
			return fmt.Errorf("movies: -i <movie_id> is required for -a recommendations")
		}
		handler = handlers.MoviesRecommendationsHandler{MovieID: *movieID, Language: *language, PagesLimit: *pagesLimit}
		params = []string{fmt.Sprintf("id-%d", *movieID)}
	default:
		return fmt.Errorf("movies: unknown action %q", *action)
	}

	result, err := handler.Handle(context.Background(), client)
	if err != nil {
		if !retried && moviesSessionActions[*action] && isSessionInvalid(err) {
			if newSession, refreshErr := refreshSession(fs, config, client); refreshErr == nil {
				options.Session = newSession
				return execMoviesAttempt(fs, client, config, options, args, true)
			}
		}
		return err
	}

	return writeResult(fs, config, "movies", *action, result, params...)
}
