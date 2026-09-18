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

// GuestSessionsCmd is the "guest-sessions" module.
var GuestSessionsCmd = &Command{
	Name:   "guest-sessions",
	Abbrev: "gs",
	Short:  "guest session rated movies/tv/tv-episodes",
	Exec:   execGuestSessions,
}

func execGuestSessions(fs afero.Fs, client *internal.Client, config *cfg.Config, options *str.Options, args []string) error {
	flagSet := flag.NewFlagSet("guest-sessions", flag.ContinueOnError)
	action := flagSet.String("a", "", "action: create, rated-movies, rated-tv, rated-tv-episodes (required)")
	guestSessionID := flagSet.String("i", "", "guest session id (optional for rated-movies, rated-tv, rated-tv-episodes: omit to use the one cached by -a create)")
	language := flagSet.String("language", "", "ISO 639-1 language code, used by -a rated-movies, rated-tv, rated-tv-episodes")
	sortBy := flagSet.String("sort-by", "", "sort order (created_at.asc, created_at.desc), used by -a rated-movies, rated-tv, rated-tv-episodes")
	pagesLimit := flagSet.Int("pages-limit", config.PagesLimit, "pages limit, used by -a rated-movies, rated-tv, rated-tv-episodes (default: pages_limit from config, 0 = unlimited)")
	if err := flagSet.Parse(args); err != nil {
		return err
	}

	id := *guestSessionID
	if id == "" && options.GuestSession != nil && options.GuestSession.GuestSessionID != "" {
		id = options.GuestSession.GuestSessionID
	}

	var handler handlers.Handler
	switch *action {
	case "":
		return fmt.Errorf("guest-sessions: -a is required (action: create, rated-movies, rated-tv, rated-tv-episodes)")
	case "create":
		handler = handlers.GuestSessionsCreateHandler{}
	case "rated-movies":
		if id == "" {
			return fmt.Errorf("guest-sessions: -i <guest_session_id> is required for -a rated-movies (or run -a create once to cache it)")
		}
		handler = handlers.GuestSessionsRatedMoviesHandler{GuestSessionID: id, Language: *language, SortBy: *sortBy, PagesLimit: *pagesLimit}
	case "rated-tv":
		if id == "" {
			return fmt.Errorf("guest-sessions: -i <guest_session_id> is required for -a rated-tv (or run -a create once to cache it)")
		}
		handler = handlers.GuestSessionsRatedTVHandler{GuestSessionID: id, Language: *language, SortBy: *sortBy, PagesLimit: *pagesLimit}
	case "rated-tv-episodes":
		if id == "" {
			return fmt.Errorf("guest-sessions: -i <guest_session_id> is required for -a rated-tv-episodes (or run -a create once to cache it)")
		}
		handler = handlers.GuestSessionsRatedTVEpisodesHandler{GuestSessionID: id, Language: *language, SortBy: *sortBy, PagesLimit: *pagesLimit}
	default:
		return fmt.Errorf("guest-sessions: unknown action %q", *action)
	}

	result, err := handler.Handle(context.Background(), client)
	if err != nil {
		return err
	}

	var params []string
	if guestSession, ok := result.(*str.GuestSession); ok {
		if err := cfg.WriteGuestSession(fs, config.GuestSessionPath, guestSession); err != nil {
			return fmt.Errorf("guest-sessions: cache guest session: %w", err)
		}
	} else if id != "" {
		params = []string{fmt.Sprintf("id-%s", id)}
	}

	return writeResult(fs, config, "guest-sessions", *action, result, params...)
}
