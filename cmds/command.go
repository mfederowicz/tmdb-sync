// Package cmds used for commands modules
package cmds

import (
	"errors"
	"flag"
	"fmt"
	"math"
	"slices"
	"strings"

	"github.com/mfederowicz/tmdb-sync/cfg"
	"github.com/mfederowicz/tmdb-sync/internal"
	"github.com/mfederowicz/tmdb-sync/str"

	"github.com/spf13/afero"
)

// Command represents one top-level module (e.g. "movies", "configuration").
type Command struct {
	Name   string
	Abbrev string
	Short  string
	Exec   func(fs afero.Fs, client *internal.Client, config *cfg.Config, options *str.Options, args []string) error
}

// resolveRatingGuestSessionID resolves the guest session id to use for an
// add-rating action: an explicit -guest-session-id flag wins; then -guest
// selects the guest session cached by `guest-sessions -a create` (an error if
// none is cached); otherwise, if no account session is set up, it falls back
// to the cached guest session. Returns "" for any other action, or when an
// account session is available (movie/tv/tv-episode add-rating then rates on
// behalf of the account session as before).
func resolveRatingGuestSessionID(action, guestSessionIDFlag string, guest bool, options *str.Options) (string, error) {
	if action != "add-rating" {
		return "", nil
	}
	if guestSessionIDFlag != "" {
		return guestSessionIDFlag, nil
	}
	if guest {
		if options.GuestSession == nil || options.GuestSession.GuestSessionID == "" {
			return "", errors.New("no cached guest session, run guest-sessions -a create (or pass -guest-session-id)")
		}
		return options.GuestSession.GuestSessionID, nil
	}
	if options.Session != nil && options.Session.SessionID != "" {
		return "", nil
	}
	if options.GuestSession != nil {
		return options.GuestSession.GuestSessionID, nil
	}
	return "", nil
}

// validateRatingValue checks -value for add-rating up front (before any login):
// TMDB accepts only 0.5 to 10.0 in 0.5 steps. Other actions are ignored.
func validateRatingValue(module, action string, value float64) error {
	if action != "add-rating" {
		return nil
	}
	if value == 0 {
		return fmt.Errorf("%s: -value <rating> is required for -a add-rating", module)
	}
	if value < 0.5 || value > 10 || math.Mod(value*2, 1) != 0 {
		return fmt.Errorf("%s: -value must be between 0.5 and 10.0 in 0.5 steps, got %v", module, value)
	}
	return nil
}

// requireSeasonEpisodeFlags errors when a known action runs without an explicit
// -s (and -e, if withEpisode). Season and episode 0 are valid (specials), so a
// zero default can't tell "omitted" from "asked for specials"; the flag set can.
// Unknown actions are left for the caller's own unknown-action error.
func requireSeasonEpisodeFlags(flagSet *flag.FlagSet, module, action, actionsHelp string, withEpisode bool) error {
	if !slices.Contains(strings.Split(actionsHelp, ", "), action) {
		return nil
	}
	set := map[string]bool{}
	flagSet.Visit(func(f *flag.Flag) { set[f.Name] = true })
	if !set["s"] {
		return fmt.Errorf("%s: -s <season_number> is required for -a %s (0 = specials)", module, action)
	}
	if withEpisode && !set["e"] {
		return fmt.Errorf("%s: -e <episode_number> is required for -a %s", module, action)
	}
	return nil
}
