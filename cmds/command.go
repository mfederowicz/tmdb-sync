// Package cmds used for commands modules
package cmds

import (
	"errors"

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
