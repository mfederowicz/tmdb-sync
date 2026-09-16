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
	accountID := flagSet.Int64("i", 0, "account id, required for -a details")
	if err := flagSet.Parse(args); err != nil {
		return err
	}

	if err := cli.HandleToken(fs, config, client, options); err != nil {
		return fmt.Errorf("account: %w", err)
	}

	var handler handlers.Handler
	var params []string
	switch *action {
	case "":
		return fmt.Errorf("account: -a is required (action: details)")
	case "details":
		if *accountID == 0 {
			return fmt.Errorf("account: -i <account_id> is required for -a details")
		}
		handler = handlers.AccountDetailsHandler{AccountID: *accountID, SessionID: options.Session.SessionID}
		params = []string{fmt.Sprintf("id-%d", *accountID)}
	default:
		return fmt.Errorf("account: unknown action %q", *action)
	}

	result, err := handler.Handle(context.Background(), client)
	if err != nil {
		return err
	}

	return writeResult(fs, config, "account", *action, result, params...)
}
