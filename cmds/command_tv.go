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

// tvSessionActions are the actions that require a v3 session.
var tvSessionActions = map[string]bool{
	"account-states": true,
	"add-rating":     true,
	"delete-rating":  true,
}

// tvActionsHelp lists every tv action, shared between the -a flag's usage
// string and the "-a is required" error so both stay in sync.
const tvActionsHelp = "details, account-states, aggregate-credits, alternative-titles, content-ratings, credits, episode-groups, external-ids"

// tvIDActionsHelp lists the tv actions that require -i, shared between the
// -i flag's usage string and the module doc.
const tvIDActionsHelp = "details, account-states, aggregate-credits, alternative-titles, content-ratings, credits, episode-groups, external-ids"

// TVCmd is the "tv" module.
var TVCmd = &Command{
	Name:   "tv",
	Abbrev: "tv",
	Short:  "TV series details and more, plus 🔒 account-states/add-rating/delete-rating",
	Exec:   execTV,
}

func execTV(fs afero.Fs, client *internal.Client, config *cfg.Config, options *str.Options, args []string) error {
	return execTVAttempt(fs, client, config, options, args, false)
}

// execTVAttempt is execTV's body, split out so a stale session (detected via
// isSessionInvalid) can trigger one transparent re-login and retry instead of
// failing outright.
func execTVAttempt(fs afero.Fs, client *internal.Client, config *cfg.Config, options *str.Options, args []string, retried bool) error {
	flagSet := flag.NewFlagSet("tv", flag.ContinueOnError)
	action := flagSet.String("a", "", "action: "+tvActionsHelp+" (required)")
	seriesID := flagSet.Int64("i", 0, "tv series id, required for -a "+tvIDActionsHelp)
	language := flagSet.String("language", "", "ISO 639-1 language code, used by -a aggregate-credits")
	if err := flagSet.Parse(args); err != nil {
		return err
	}

	if tvSessionActions[*action] {
		if err := cli.HandleToken(fs, config, client, options); err != nil {
			return fmt.Errorf("tv: %w", err)
		}
	}

	var handler handlers.Handler
	var params []string
	switch *action {
	case "":
		return fmt.Errorf("tv: -a is required (action: %s)", tvActionsHelp)
	case "details":
		if *seriesID == 0 {
			return fmt.Errorf("tv: -i <series_id> is required for -a details")
		}
		handler = handlers.TVDetailsHandler{SeriesID: *seriesID}
		params = []string{fmt.Sprintf("id-%d", *seriesID)}
	case "account-states":
		if *seriesID == 0 {
			return fmt.Errorf("tv: -i <series_id> is required for -a account-states")
		}
		handler = handlers.TVAccountStatesHandler{SeriesID: *seriesID, SessionID: options.Session.SessionID}
		params = []string{fmt.Sprintf("id-%d", *seriesID)}
	case "aggregate-credits":
		if *seriesID == 0 {
			return fmt.Errorf("tv: -i <series_id> is required for -a aggregate-credits")
		}
		handler = handlers.TVAggregateCreditsHandler{SeriesID: *seriesID, Language: *language}
		params = []string{fmt.Sprintf("id-%d", *seriesID)}
	case "alternative-titles":
		if *seriesID == 0 {
			return fmt.Errorf("tv: -i <series_id> is required for -a alternative-titles")
		}
		handler = handlers.TVAlternativeTitlesHandler{SeriesID: *seriesID}
		params = []string{fmt.Sprintf("id-%d", *seriesID)}
	case "content-ratings":
		if *seriesID == 0 {
			return fmt.Errorf("tv: -i <series_id> is required for -a content-ratings")
		}
		handler = handlers.TVContentRatingsHandler{SeriesID: *seriesID}
		params = []string{fmt.Sprintf("id-%d", *seriesID)}
	case "credits":
		if *seriesID == 0 {
			return fmt.Errorf("tv: -i <series_id> is required for -a credits")
		}
		handler = handlers.TVCreditsHandler{SeriesID: *seriesID, Language: *language}
		params = []string{fmt.Sprintf("id-%d", *seriesID)}
	case "episode-groups":
		if *seriesID == 0 {
			return fmt.Errorf("tv: -i <series_id> is required for -a episode-groups")
		}
		handler = handlers.TVEpisodeGroupsHandler{SeriesID: *seriesID}
		params = []string{fmt.Sprintf("id-%d", *seriesID)}
	case "external-ids":
		if *seriesID == 0 {
			return fmt.Errorf("tv: -i <series_id> is required for -a external-ids")
		}
		handler = handlers.TVExternalIDsHandler{SeriesID: *seriesID}
		params = []string{fmt.Sprintf("id-%d", *seriesID)}
	default:
		return fmt.Errorf("tv: unknown action %q", *action)
	}

	result, err := handler.Handle(context.Background(), client)
	if err != nil {
		if !retried && tvSessionActions[*action] && isSessionInvalid(err) {
			if newSession, refreshErr := refreshSession(fs, config, client); refreshErr == nil {
				options.Session = newSession
				return execTVAttempt(fs, client, config, options, args, true)
			}
		}
		return err
	}

	return writeResult(fs, config, "tv", *action, result, params...)
}
