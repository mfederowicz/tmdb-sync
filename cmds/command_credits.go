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

// CreditsCmd is the "credits" module.
var CreditsCmd = &Command{
	Name:   "credits",
	Abbrev: "cr",
	Short:  "credit details",
	Exec:   execCredits,
}

func execCredits(fs afero.Fs, client *internal.Client, config *cfg.Config, _ *str.Options, args []string) error {
	flagSet := flag.NewFlagSet("credits", flag.ContinueOnError)
	action := flagSet.String("a", "", "action: details (required)")
	creditID := flagSet.String("i", "", "credit id, required for all actions")
	if err := flagSet.Parse(args); err != nil {
		return err
	}

	var handler handlers.Handler
	switch *action {
	case "":
		return fmt.Errorf("credits: -a is required (action: details)")
	case "details":
		if *creditID == "" {
			return fmt.Errorf("credits: -i <credit_id> is required for -a details")
		}
		handler = handlers.CreditsDetailsHandler{CreditID: *creditID}
	default:
		return fmt.Errorf("credits: unknown action %q", *action)
	}

	result, err := handler.Handle(context.Background(), client)
	if err != nil {
		return err
	}

	params := []string{fmt.Sprintf("id-%s", *creditID)}
	return writeResult(fs, config, "credits", *action, result, params...)
}
