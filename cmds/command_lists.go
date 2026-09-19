package cmds

import (
	"context"
	"flag"
	"fmt"
	"slices"
	"strings"

	"github.com/mfederowicz/tmdb-sync/cfg"
	"github.com/mfederowicz/tmdb-sync/cli"
	"github.com/mfederowicz/tmdb-sync/handlers"
	"github.com/mfederowicz/tmdb-sync/internal"
	"github.com/mfederowicz/tmdb-sync/str"
	"github.com/mfederowicz/tmdb-sync/uri"

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

// listsV4Actions are the `lists` actions implemented for -v4.
var listsV4Actions = []string{"details", "create"}

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
	v3 := flagSet.Bool("v3", false, "use the v3 API (default)")
	v4 := flagSet.Bool("v4", false, "use the v4 API (optionally with `auth -v4 -a login` for private lists); actions: details, create")
	sortBy := flagSet.String("sort-by", "", "v4 only: sort order of the items, for -a details (e.g. original_order.asc, vote_average.desc)")
	country := flagSet.String("country", "", "v4 only: list country (ISO 3166-1, e.g. US), required for -a create")
	public := flagSet.Bool("public", false, "v4 only: make the list public, used by -a create")
	pagesLimit := flagSet.Int("pages-limit", config.PagesLimit, "v4 only: item pages limit for -a details (default: pages_limit from config, 0 = unlimited)")
	if err := flagSet.Parse(args); err != nil {
		return err
	}

	version, err := resolveAPIVersion(*v3, *v4)
	if err != nil {
		return fmt.Errorf("lists: %w", err)
	}
	v4Mode := version == apiV4
	if v4Mode && *action != "" && !slices.Contains(listsV4Actions, *action) {
		return fmt.Errorf("lists: action %q is not available with -v4 (v4 actions: %s)", *action, strings.Join(listsV4Actions, ", "))
	}
	if !v4Mode && (*sortBy != "" || *country != "" || *public) {
		return fmt.Errorf("lists: -sort-by, -country and -public need -v4")
	}
	if v4Mode && *action != "" && (options.AccessTokenV4 == nil || options.AccessTokenV4.AccessToken == "") && *action != "details" {
		return fmt.Errorf("lists: no v4 access token cached at %s, run `auth -v4 -a login` first", config.AccessTokenPath)
	}

	if !v4Mode && listsSessionActions[*action] {
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
		if v4Mode {
			accessToken := ""
			if options.AccessTokenV4 != nil {
				accessToken = options.AccessTokenV4.AccessToken
			}
			handler = handlers.ListsDetailsHandler{ListID: *listID, V4: true, AccessToken: accessToken, PagesLimit: *pagesLimit, V4Options: uri.ListV4Options{Language: *language, SortBy: *sortBy}}
			params = []string{fmt.Sprintf("id-%s", *listID), "v4"}
			break
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
		if v4Mode {
			if *language == "" || *country == "" {
				return fmt.Errorf("lists: -language and -country are required for -v4 -a create")
			}
			handler = handlers.ListsCreateHandler{V4: true, AccessToken: options.AccessTokenV4.AccessToken, Name: *name, Description: *description, Language: *language, Country: *country, Public: *public}
			params = []string{"v4"}
			break
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
	if created, ok := result.(*str.ListCreateResponseV4); ok {
		params = []string{fmt.Sprintf("id-%d", created.ID), "v4"}
	}

	return writeResult(fs, config, "lists", *action, result, params...)
}
