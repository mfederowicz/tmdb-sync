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

// tvSeasonsSessionActions are the actions that require a v3 session.
var tvSeasonsSessionActions = map[string]bool{
	"account-states": true,
}

// tvSeasonsActionsHelp lists every tv-seasons action, shared between the -a
// flag's usage string and the "-a is required" error so both stay in sync.
const tvSeasonsActionsHelp = "details, account-states, aggregate-credits, credits, external-ids, images"

// TVSeasonsCmd is the "tv-seasons" module.
var TVSeasonsCmd = &Command{
	Name:   "tv-seasons",
	Abbrev: "tv-seasons",
	Short:  "TV season details and more, plus 🔒 account-states",
	Exec:   execTVSeasons,
}

func execTVSeasons(fs afero.Fs, client *internal.Client, config *cfg.Config, options *str.Options, args []string) error {
	return execTVSeasonsAttempt(fs, client, config, options, args, false)
}

// execTVSeasonsAttempt is execTVSeasons's body, split out so a stale session
// (detected via isSessionInvalid) can trigger one transparent re-login and
// retry instead of failing outright.
func execTVSeasonsAttempt(fs afero.Fs, client *internal.Client, config *cfg.Config, options *str.Options, args []string, retried bool) error {
	flagSet := flag.NewFlagSet("tv-seasons", flag.ContinueOnError)
	action := flagSet.String("a", "", "action: "+tvSeasonsActionsHelp+" (required)")
	seriesID := flagSet.Int64("i", 0, "tv series id, required for -a "+tvSeasonsActionsHelp)
	seasonNumber := flagSet.Int("s", 0, "season number, required for -a "+tvSeasonsActionsHelp)
	language := flagSet.String("language", "", "ISO 639-1 language code, used by -a aggregate-credits, credits, images")
	includeImageLanguage := flagSet.String("include-image-language", "", "comma-separated language codes, used by -a images")
	if err := flagSet.Parse(args); err != nil {
		return err
	}

	if tvSeasonsSessionActions[*action] {
		if err := cli.HandleToken(fs, config, client, options); err != nil {
			return fmt.Errorf("tv-seasons: %w", err)
		}
	}

	var handler handlers.Handler
	var params []string
	switch *action {
	case "":
		return fmt.Errorf("tv-seasons: -a is required (action: %s)", tvSeasonsActionsHelp)
	case "details":
		if *seriesID == 0 {
			return fmt.Errorf("tv-seasons: -i <series_id> is required for -a details")
		}
		handler = handlers.TVSeasonsDetailsHandler{SeriesID: *seriesID, SeasonNumber: *seasonNumber}
		params = []string{fmt.Sprintf("id-%d", *seriesID), fmt.Sprintf("season-%d", *seasonNumber)}
	case "account-states":
		if *seriesID == 0 {
			return fmt.Errorf("tv-seasons: -i <series_id> is required for -a account-states")
		}
		handler = handlers.TVSeasonsAccountStatesHandler{SeriesID: *seriesID, SeasonNumber: *seasonNumber, SessionID: options.Session.SessionID}
		params = []string{fmt.Sprintf("id-%d", *seriesID), fmt.Sprintf("season-%d", *seasonNumber)}
	case "aggregate-credits":
		if *seriesID == 0 {
			return fmt.Errorf("tv-seasons: -i <series_id> is required for -a aggregate-credits")
		}
		handler = handlers.TVSeasonsAggregateCreditsHandler{SeriesID: *seriesID, SeasonNumber: *seasonNumber, Language: *language}
		params = []string{fmt.Sprintf("id-%d", *seriesID), fmt.Sprintf("season-%d", *seasonNumber)}
	case "credits":
		if *seriesID == 0 {
			return fmt.Errorf("tv-seasons: -i <series_id> is required for -a credits")
		}
		handler = handlers.TVSeasonsCreditsHandler{SeriesID: *seriesID, SeasonNumber: *seasonNumber, Language: *language}
		params = []string{fmt.Sprintf("id-%d", *seriesID), fmt.Sprintf("season-%d", *seasonNumber)}
	case "external-ids":
		if *seriesID == 0 {
			return fmt.Errorf("tv-seasons: -i <series_id> is required for -a external-ids")
		}
		handler = handlers.TVSeasonsExternalIDsHandler{SeriesID: *seriesID, SeasonNumber: *seasonNumber}
		params = []string{fmt.Sprintf("id-%d", *seriesID), fmt.Sprintf("season-%d", *seasonNumber)}
	case "images":
		if *seriesID == 0 {
			return fmt.Errorf("tv-seasons: -i <series_id> is required for -a images")
		}
		handler = handlers.TVSeasonsImagesHandler{SeriesID: *seriesID, SeasonNumber: *seasonNumber, Language: *language, IncludeImageLanguage: *includeImageLanguage}
		params = []string{fmt.Sprintf("id-%d", *seriesID), fmt.Sprintf("season-%d", *seasonNumber)}
	default:
		return fmt.Errorf("tv-seasons: unknown action %q (action: %s)", *action, tvSeasonsActionsHelp)
	}

	result, err := handler.Handle(context.Background(), client)
	if err != nil {
		if !retried && tvSeasonsSessionActions[*action] && isSessionInvalid(err) {
			if newSession, refreshErr := refreshSession(fs, config, client); refreshErr == nil {
				options.Session = newSession
				return execTVSeasonsAttempt(fs, client, config, options, args, true)
			}
		}
		return err
	}

	return writeResult(fs, config, "tv-seasons", *action, result, params...)
}
