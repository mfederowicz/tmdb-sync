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

// ReviewsCmd is the "reviews" module.
var ReviewsCmd = &Command{
	Name:   "reviews",
	Abbrev: "rv",
	Short:  "review details",
	Exec:   execReviews,
}

func execReviews(fs afero.Fs, client *internal.Client, config *cfg.Config, _ *str.Options, args []string) error {
	flagSet := flag.NewFlagSet("reviews", flag.ContinueOnError)
	action := flagSet.String("a", "", "action: details (required)")
	reviewID := flagSet.String("i", "", "review id, required for all actions")
	if err := flagSet.Parse(args); err != nil {
		return err
	}

	var handler handlers.Handler
	switch *action {
	case "":
		return fmt.Errorf("reviews: -a is required (action: details)")
	case "details":
		if *reviewID == "" {
			return fmt.Errorf("reviews: -i <review_id> is required for -a details")
		}
		handler = handlers.ReviewsDetailsHandler{ReviewID: *reviewID}
	default:
		return fmt.Errorf("reviews: unknown action %q", *action)
	}

	result, err := handler.Handle(context.Background(), client)
	if err != nil {
		return err
	}

	params := []string{fmt.Sprintf("id-%s", *reviewID)}
	return writeResult(fs, config, "reviews", *action, result, params...)
}
