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

// ListsCmd is the "lists" module.
var ListsCmd = &Command{
	Name:   "lists",
	Abbrev: "li",
	Short:  "list details, item status",
	Exec:   execLists,
}

func execLists(fs afero.Fs, client *internal.Client, config *cfg.Config, _ *str.Options, args []string) error {
	flagSet := flag.NewFlagSet("lists", flag.ContinueOnError)
	action := flagSet.String("a", "", "action: details, item-status (required)")
	listID := flagSet.String("i", "", "list id, required for all actions")
	movieID := flagSet.Int64("media-id", 0, "movie id, required for -a item-status")
	if err := flagSet.Parse(args); err != nil {
		return err
	}

	var handler handlers.Handler
	var params []string
	switch *action {
	case "":
		return fmt.Errorf("lists: -a is required (action: details, item-status)")
	case "details":
		if *listID == "" {
			return fmt.Errorf("lists: -i <list_id> is required for -a details")
		}
		handler = handlers.ListsDetailsHandler{ListID: *listID}
		params = []string{fmt.Sprintf("id-%s", *listID)}
	case "item-status":
		if *listID == "" {
			return fmt.Errorf("lists: -i <list_id> is required for -a item-status")
		}
		if *movieID == 0 {
			return fmt.Errorf("lists: -media-id <movie_id> is required for -a item-status")
		}
		handler = handlers.ListsItemStatusHandler{ListID: *listID, MovieID: *movieID}
		params = []string{fmt.Sprintf("id-%s", *listID), fmt.Sprintf("media-%d", *movieID)}
	default:
		return fmt.Errorf("lists: unknown action %q", *action)
	}

	result, err := handler.Handle(context.Background(), client)
	if err != nil {
		return err
	}

	return writeResult(fs, config, "lists", *action, result, params...)
}
