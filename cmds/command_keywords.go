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

// KeywordsCmd is the "keywords" module.
var KeywordsCmd = &Command{
	Name:   "keywords",
	Abbrev: "kw",
	Short:  "keyword details",
	Exec:   execKeywords,
}

func execKeywords(fs afero.Fs, client *internal.Client, config *cfg.Config, _ *str.Options, args []string) error {
	flagSet := flag.NewFlagSet("keywords", flag.ContinueOnError)
	action := flagSet.String("a", "", "action: details (required)")
	keywordID := flagSet.String("i", "", "keyword id, required for all actions")
	if err := flagSet.Parse(args); err != nil {
		return err
	}

	var handler handlers.Handler
	switch *action {
	case "":
		return fmt.Errorf("keywords: -a is required (action: details)")
	case "details":
		if *keywordID == "" {
			return fmt.Errorf("keywords: -i <keyword_id> is required for -a details")
		}
		handler = handlers.KeywordsDetailsHandler{KeywordID: *keywordID}
	default:
		return fmt.Errorf("keywords: unknown action %q", *action)
	}

	result, err := handler.Handle(context.Background(), client)
	if err != nil {
		return err
	}

	params := []string{fmt.Sprintf("id-%s", *keywordID)}
	return writeResult(fs, config, "keywords", *action, result, params...)
}
