package cmds

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/mfederowicz/tmdb-sync/cfg"
	"github.com/mfederowicz/tmdb-sync/handlers"
	"github.com/mfederowicz/tmdb-sync/internal"
	"github.com/mfederowicz/tmdb-sync/printer"
	"github.com/mfederowicz/tmdb-sync/str"

	"github.com/spf13/afero"
)

// AuthCmd is the "auth" module: v4 login/logout, plus v3 logout.
var AuthCmd = &Command{
	Name:   "auth",
	Abbrev: "au",
	Short:  "v4 auth: request token, access token, login, logout (logout also v3)",
	Exec:   execAuth,
}

func execAuth(fs afero.Fs, client *internal.Client, config *cfg.Config, options *str.Options, args []string) error {
	flagSet := flag.NewFlagSet("auth", flag.ContinueOnError)
	action := flagSet.String("a", "", "action: request-token, access-token, login, logout (required)")
	v3 := flagSet.Bool("v3", false, "use the v3 API (default; only -a logout exists in v3)")
	v4 := flagSet.Bool("v4", false, "use the v4 API (required for every action except a v3 logout)")
	requestToken := flagSet.String("request-token", "", "approved v4 request token, required by -a access-token")
	redirectTo := flagSet.String("redirect-to", "", "URL TMDB redirects to after approval, used by -a request-token")
	if err := flagSet.Parse(args); err != nil {
		return err
	}

	version, err := resolveAPIVersion(*v3, *v4)
	if err != nil {
		return fmt.Errorf("auth: %w", err)
	}
	if version != apiV4 && *action != "logout" {
		return fmt.Errorf("auth: only v4 is supported, add -v4 (v3 login happens on demand in account commands; v3 supports only -a logout)")
	}

	var handler handlers.Handler
	switch *action {
	case "":
		return fmt.Errorf("auth: -a is required (action: request-token, access-token, login, logout)")
	case "request-token":
		handler = handlers.AuthRequestTokenHandler{RedirectTo: *redirectTo}
	case "access-token":
		if *requestToken == "" {
			return fmt.Errorf("auth: -request-token is required for -a access-token (approve one from -a request-token first)")
		}
		handler = handlers.AuthAccessTokenHandler{RequestToken: *requestToken}
	case "login":
		handler = handlers.AuthLoginHandler{Fs: fs, Config: config}
	case "logout":
		if version != apiV4 {
			if options.Session == nil || options.Session.SessionID == "" {
				return fmt.Errorf("auth: no v3 session cached at %s, nothing to log out", config.SessionPath)
			}
			handler = handlers.AuthLogoutHandler{SessionID: options.Session.SessionID}
			break
		}
		if options.AccessTokenV4 == nil || options.AccessTokenV4.AccessToken == "" {
			return fmt.Errorf("auth: no v4 access token cached at %s, run -a login first", config.AccessTokenPath)
		}
		handler = handlers.AuthLogoutHandler{V4: true, AccessToken: options.AccessTokenV4.AccessToken}
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

	if *action == "logout" && version != apiV4 {
		// The cached account id belongs to the session that just ended.
		for _, path := range []string{config.SessionPath, config.AccountPath} {
			if err := fs.Remove(path); err != nil && !os.IsNotExist(err) {
				return fmt.Errorf("auth: remove cached %s: %w", path, err)
			}
		}
		printer.Println("logged out, removed", config.SessionPath, "and", config.AccountPath)
		return writeResult(fs, config, "auth", *action, result, "v3")
	}

	if *action == "logout" {
		if err := fs.Remove(config.AccessTokenPath); err != nil && !os.IsNotExist(err) {
			return fmt.Errorf("auth: remove cached access token: %w", err)
		}
		printer.Println("logged out, removed", config.AccessTokenPath)
	}

	return writeResult(fs, config, "auth", *action, result, "v4")
}
