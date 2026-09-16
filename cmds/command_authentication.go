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

// AuthenticationCmd is the "authentication" module.
var AuthenticationCmd = &Command{
	Name:   "authentication",
	Abbrev: "auth",
	Short:  "validate key, request tokens, sessions, guest sessions",
	Exec:   execAuthentication,
}

func execAuthentication(fs afero.Fs, client *internal.Client, config *cfg.Config, options *str.Options, args []string) error {
	flagSet := flag.NewFlagSet("authentication", flag.ContinueOnError)
	action := flagSet.String("a", "validate-key", "action: validate-key, create-request-token, create-session, create-guest-session, delete-session")
	sessionID := flagSet.String("s", "", "session id, used by -a delete-session (default: the persisted session)")
	if err := flagSet.Parse(args); err != nil {
		return err
	}

	// create-session and delete-session need fs/config/options beyond what
	// handlers.Handler's (ctx, client) signature carries, so they call
	// straight into the cli/internal auth flow instead of going through a
	// handlers.Handler like the other three actions below.
	switch *action {
	case "create-session":
		session, err := cli.CreateSessionInteractively(fs, config, client)
		if err != nil {
			return err
		}
		return writeResult(fs, config, "authentication", *action, session)
	case "delete-session":
		id := *sessionID
		if id == "" && options != nil && options.Session != nil {
			id = options.Session.SessionID
		}
		if id == "" {
			return fmt.Errorf("authentication: -s <session_id> is required for -a delete-session (no persisted session found)")
		}
		status, _, err := client.Auth.DeleteSession(context.Background(), id)
		if err != nil {
			return err
		}
		return writeResult(fs, config, "authentication", *action, status)
	}

	var handler handlers.Handler
	switch *action {
	case "validate-key":
		handler = handlers.AuthenticationValidateKeyHandler{}
	case "create-request-token":
		handler = handlers.AuthenticationCreateRequestTokenHandler{}
	case "create-guest-session":
		handler = handlers.AuthenticationCreateGuestSessionHandler{}
	default:
		return fmt.Errorf("authentication: unknown action %q", *action)
	}

	result, err := handler.Handle(context.Background(), client)
	if err != nil {
		return err
	}

	return writeResult(fs, config, "authentication", *action, result)
}
