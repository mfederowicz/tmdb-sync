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

// AuthCmd is the "auth" module (v4 only).
var AuthCmd = &Command{
	Name:   "auth",
	Abbrev: "au",
	Short:  "v4 auth: request token",
	Exec:   execAuth,
}

func execAuth(fs afero.Fs, client *internal.Client, config *cfg.Config, _ *str.Options, args []string) error {
	flagSet := flag.NewFlagSet("auth", flag.ContinueOnError)
	action := flagSet.String("a", "", "action: request-token (required)")
	v3 := flagSet.Bool("v3", false, "use the v3 API (not available for auth)")
	v4 := flagSet.Bool("v4", false, "use the v4 API (required for auth)")
	redirectTo := flagSet.String("redirect-to", "", "URL TMDB redirects to after approval, used by -a request-token")
	if err := flagSet.Parse(args); err != nil {
		return err
	}

	version, err := resolveAPIVersion(*v3, *v4)
	if err != nil {
		return fmt.Errorf("auth: %w", err)
	}
	if version != apiV4 {
		return fmt.Errorf("auth: only v4 is supported, add -v4 (v3 login happens on demand in account commands)")
	}

	var handler handlers.Handler
	switch *action {
	case "":
		return fmt.Errorf("auth: -a is required (action: request-token)")
	case "request-token":
		handler = handlers.AuthRequestTokenHandler{RedirectTo: *redirectTo}
	default:
		return fmt.Errorf("auth: unknown action %q", *action)
	}

	result, err := handler.Handle(context.Background(), client)
	if err != nil {
		return err
	}

	return writeResult(fs, config, "auth", *action, result, "v4")
}
