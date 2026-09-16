package cmds

import (
	"context"
	"flag"
	"fmt"

	"github.com/mfederowicz/tmdb-sync/cfg"
	"github.com/mfederowicz/tmdb-sync/cli"
	"github.com/mfederowicz/tmdb-sync/handlers"
	"github.com/mfederowicz/tmdb-sync/internal"
	"github.com/spf13/afero"

	"github.com/mfederowicz/tmdb-sync/str"
)

// AccountCmd is the "account" 🔒 module. Every action requires a v3 session,
// established on demand via cli.HandleToken.
var AccountCmd = &Command{
	Name:   "account",
	Abbrev: "acc",
	Short:  "🔒 account details, favorites, watchlist, lists, rated",
	Exec:   execAccount,
}

func execAccount(fs afero.Fs, client *internal.Client, config *cfg.Config, options *str.Options, args []string) error {
	flagSet := flag.NewFlagSet("account", flag.ContinueOnError)
	action := flagSet.String("a", "", "action: details (required)")
	accountID := flagSet.Int64("i", 0, "account id (optional for -a details: omit to self-resolve via the session and cache it)")
	if err := flagSet.Parse(args); err != nil {
		return err
	}

	if err := cli.HandleToken(fs, config, client, options); err != nil {
		return fmt.Errorf("account: %w", err)
	}

	id := *accountID
	if id == 0 && options.Account != nil && options.Account.ID != 0 {
		id = options.Account.ID
	}

	var handler handlers.Handler
	var params []string
	switch *action {
	case "":
		return fmt.Errorf("account: -a is required (action: details)")
	case "details":
		handler = handlers.AccountDetailsHandler{AccountID: id, SessionID: options.Session.SessionID}
		if id != 0 {
			params = []string{fmt.Sprintf("id-%d", id)}
		}
	default:
		return fmt.Errorf("account: unknown action %q", *action)
	}

	result, err := handler.Handle(context.Background(), client)
	if err != nil {
		return err
	}

	if account, ok := result.(*str.Account); ok {
		if err := cfg.WriteAccount(fs, config.AccountPath, account); err != nil {
			return fmt.Errorf("account: cache account details: %w", err)
		}
		if len(params) == 0 {
			params = []string{fmt.Sprintf("id-%d", account.ID)}
		}
	}

	return writeResult(fs, config, "account", *action, result, params...)
}
