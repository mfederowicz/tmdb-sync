package cmds

import (
	"context"
	"flag"
	"fmt"

	"github.com/mfederowicz/tmdb-sync/cfg"
	"github.com/mfederowicz/tmdb-sync/handlers"
	"github.com/mfederowicz/tmdb-sync/internal"
	"github.com/mfederowicz/tmdb-sync/printer"
	"github.com/mfederowicz/tmdb-sync/str"

	"github.com/spf13/afero"
)

// AuthCmd is the "auth" module (v4 only).
var AuthCmd = &Command{
	Name:   "auth",
	Abbrev: "au",
	Short:  "v4 auth: request token, access token, login",
	Exec:   execAuth,
}

func execAuth(fs afero.Fs, client *internal.Client, config *cfg.Config, _ *str.Options, args []string) error {
	flagSet := flag.NewFlagSet("auth", flag.ContinueOnError)
	action := flagSet.String("a", "", "action: request-token, access-token, login (required)")
	v3 := flagSet.Bool("v3", false, "use the v3 API (not available for auth)")
	v4 := flagSet.Bool("v4", false, "use the v4 API (required for auth)")
	requestToken := flagSet.String("request-token", "", "approved v4 request token, required by -a access-token")
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
		return fmt.Errorf("auth: -a is required (action: request-token, access-token, login)")
	case "request-token":
		handler = handlers.AuthRequestTokenHandler{RedirectTo: *redirectTo}
	case "access-token":
		if *requestToken == "" {
			return fmt.Errorf("auth: -request-token is required for -a access-token (approve one from -a request-token first)")
		}
		handler = handlers.AuthAccessTokenHandler{RequestToken: *requestToken}
	case "login":
		handler = handlers.AuthLoginHandler{Fs: fs, Config: config}
	default:
		return fmt.Errorf("auth: unknown action %q", *action)
	}

	result, err := handler.Handle(context.Background(), client)
	if err != nil {
		return err
	}

	// The user access token is a secret: persist it to access_token_path only,
	// never into the output dir like other results.
	if accessToken, ok := result.(*str.AccessTokenV4); ok {
		if !accessToken.Valid() {
			return fmt.Errorf("auth: tmdb refused to create an access token for the request token")
		}
		if err := cfg.WriteAccessToken(fs, config.AccessTokenPath, accessToken); err != nil {
			return fmt.Errorf("auth: persist access token: %w", err)
		}
		printer.Println("logged in, access token saved to", config.AccessTokenPath, "(account_id:", accessToken.AccountID+")")
		return nil
	}

	return writeResult(fs, config, "auth", *action, result, "v4")
}
