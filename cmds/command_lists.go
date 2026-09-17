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

// listsSessionActions are the actions that require a v3 session.
var listsSessionActions = map[string]bool{
	"create":       true,
	"add-movie":    true,
	"remove-movie": true,
	"clear":        true,
	"delete":       true,
}

// ListsCmd is the "lists" module. Read actions (details, item-status) are
// public; mutation actions (create, add-movie, remove-movie, clear, delete)
// are 🔒, established on demand via cli.HandleToken.
var ListsCmd = &Command{
	Name:   "lists",
	Abbrev: "li",
	Short:  "list details, item status, 🔒 create/add-movie/remove-movie/clear/delete",
	Exec:   execLists,
}

func execLists(fs afero.Fs, client *internal.Client, config *cfg.Config, options *str.Options, args []string) error {
	return execListsAttempt(fs, client, config, options, args, false)
}

// execListsAttempt is execLists's body, split out so a stale session
// (detected via isSessionInvalid) can trigger one transparent re-login and
// retry instead of failing outright.
func execListsAttempt(fs afero.Fs, client *internal.Client, config *cfg.Config, options *str.Options, args []string, retried bool) error {
	flagSet := flag.NewFlagSet("lists", flag.ContinueOnError)
	action := flagSet.String("a", "", "action: details, item-status, create, add-movie, remove-movie, clear, delete (required)")
	listID := flagSet.String("i", "", "list id, required for all actions except create")
	movieID := flagSet.Int64("media-id", 0, "movie id, required for -a item-status, add-movie, remove-movie")
	name := flagSet.String("name", "", "list name, required for -a create")
	description := flagSet.String("description", "", "list description, used by -a create")
	language := flagSet.String("language", "", "list language (ISO 639-1), used by -a create")
	if err := flagSet.Parse(args); err != nil {
		return err
	}

	if listsSessionActions[*action] {
		if err := cli.HandleToken(fs, config, client, options); err != nil {
			return fmt.Errorf("lists: %w", err)
		}
	}

	var handler handlers.Handler
	var params []string
	switch *action {
	case "":
		return fmt.Errorf("lists: -a is required (action: details, item-status, create, add-movie, remove-movie, clear, delete)")
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
	case "create":
		if *name == "" {
			return fmt.Errorf("lists: -name <name> is required for -a create")
		}
		handler = handlers.ListsCreateHandler{
			SessionID:   options.Session.SessionID,
			Name:        *name,
			Description: *description,
			Language:    *language,
		}
	case "add-movie":
		if *listID == "" {
			return fmt.Errorf("lists: -i <list_id> is required for -a add-movie")
		}
		if *movieID == 0 {
			return fmt.Errorf("lists: -media-id <movie_id> is required for -a add-movie")
		}
		handler = handlers.ListsAddMovieHandler{ListID: *listID, SessionID: options.Session.SessionID, MovieID: *movieID}
		params = []string{fmt.Sprintf("id-%s", *listID), fmt.Sprintf("media-%d", *movieID)}
	case "remove-movie":
		if *listID == "" {
			return fmt.Errorf("lists: -i <list_id> is required for -a remove-movie")
		}
		if *movieID == 0 {
			return fmt.Errorf("lists: -media-id <movie_id> is required for -a remove-movie")
		}
		handler = handlers.ListsRemoveMovieHandler{ListID: *listID, SessionID: options.Session.SessionID, MovieID: *movieID}
		params = []string{fmt.Sprintf("id-%s", *listID), fmt.Sprintf("media-%d", *movieID)}
	case "clear":
		if *listID == "" {
			return fmt.Errorf("lists: -i <list_id> is required for -a clear")
		}
		handler = handlers.ListsClearHandler{ListID: *listID, SessionID: options.Session.SessionID}
		params = []string{fmt.Sprintf("id-%s", *listID)}
	case "delete":
		if *listID == "" {
			return fmt.Errorf("lists: -i <list_id> is required for -a delete")
		}
		handler = handlers.ListsDeleteHandler{ListID: *listID, SessionID: options.Session.SessionID}
		params = []string{fmt.Sprintf("id-%s", *listID)}
	default:
		return fmt.Errorf("lists: unknown action %q", *action)
	}

	result, err := handler.Handle(context.Background(), client)
	if err != nil {
		if !retried && listsSessionActions[*action] && isSessionInvalid(err) {
			if newSession, refreshErr := refreshSession(fs, config, client); refreshErr == nil {
				options.Session = newSession
				return execListsAttempt(fs, client, config, options, args, true)
			}
		}
		return err
	}

	if created, ok := result.(*str.ListCreateResponse); ok && len(params) == 0 {
		params = []string{fmt.Sprintf("id-%d", created.ListID)}
	}

	return writeResult(fs, config, "lists", *action, result, params...)
}
