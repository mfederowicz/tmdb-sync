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

func execGuestSessions(fs afero.Fs, client *internal.Client, config *cfg.Config, _ *str.Options, args []string) error {
	flagSet := flag.NewFlagSet("guest-sessions", flag.ContinueOnError)
	action := flagSet.String("a", "", "action: rated-movies (required)")
	guestSessionID := flagSet.String("i", "", "guest session id, required for all actions")
	language := flagSet.String("language", "", "ISO 639-1 language code, used by -a rated-movies")
	sortBy := flagSet.String("sort-by", "", "sort order (created_at.asc, created_at.desc), used by -a rated-movies")
	pagesLimit := flagSet.Int("pages-limit", config.PagesLimit, "pages limit, used by -a rated-movies (default: pages_limit from config, 0 = unlimited)")
	if err := flagSet.Parse(args); err != nil {
		return err
	}

	var handler handlers.Handler
	switch *action {
	case "":
		return fmt.Errorf("guest-sessions: -a is required (action: rated-movies)")
	case "rated-movies":
		if *guestSessionID == "" {
			return fmt.Errorf("guest-sessions: -i <guest_session_id> is required for -a rated-movies")
		}
		handler = handlers.GuestSessionsRatedMoviesHandler{GuestSessionID: *guestSessionID, Language: *language, SortBy: *sortBy, PagesLimit: *pagesLimit}
	default:
		return fmt.Errorf("guest-sessions: unknown action %q", *action)
	}

	result, err := handler.Handle(context.Background(), client)
	if err != nil {
		return err
	}

	params := []string{fmt.Sprintf("id-%s", *guestSessionID)}
	return writeResult(fs, config, "guest-sessions", *action, result, params...)
}
