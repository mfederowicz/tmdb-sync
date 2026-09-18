// Package cmds used for commands modules
package cmds

import (
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
// add-rating action: an explicit -guest-session-id flag wins; otherwise, if
// no account session is set up, it falls back to the guest session cached by
// `guest-sessions -a create`. Returns "" for any other action, or when an
// account session is available (movie/tv/tv-episode add-rating then rates on
// behalf of the account session as before).
func resolveRatingGuestSessionID(action, guestSessionIDFlag string, options *str.Options) string {
	if action != "add-rating" {
		return ""
	}
	if guestSessionIDFlag != "" {
		return guestSessionIDFlag
	}
	if options.Session != nil && options.Session.SessionID != "" {
		return ""
	}
	if options.GuestSession != nil {
		return options.GuestSession.GuestSessionID
	}
	return ""
}
