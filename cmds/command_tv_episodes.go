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

// tvEpisodesSessionActions are the actions that require a v3 session.
var tvEpisodesSessionActions = map[string]bool{
	"account-states": true,
}

// tvEpisodesActionsHelp lists every tv-episodes action, shared between the
// -a flag's usage string and the "-a is required" error so both stay in
// sync.
const tvEpisodesActionsHelp = "details, account-states"

// TVEpisodesCmd is the "tv-episodes" module.
var TVEpisodesCmd = &Command{
	Name:   "tv-episodes",
	Abbrev: "tv-episodes",
	Short:  "TV episode details and more",
	Exec:   execTVEpisodes,
}

func execTVEpisodes(fs afero.Fs, client *internal.Client, config *cfg.Config, options *str.Options, args []string) error {
	return execTVEpisodesAttempt(fs, client, config, options, args, false)
}

// execTVEpisodesAttempt is execTVEpisodes's body, split out so a stale
// session (detected via isSessionInvalid) can trigger one transparent
// re-login and retry instead of failing outright.
func execTVEpisodesAttempt(fs afero.Fs, client *internal.Client, config *cfg.Config, options *str.Options, args []string, retried bool) error {
	flagSet := flag.NewFlagSet("tv-episodes", flag.ContinueOnError)
	action := flagSet.String("a", "", "action: "+tvEpisodesActionsHelp+" (required)")
	seriesID := flagSet.Int64("i", 0, "tv series id, required for -a "+tvEpisodesActionsHelp)
	seasonNumber := flagSet.Int("s", 0, "season number, required for -a "+tvEpisodesActionsHelp)
	episodeNumber := flagSet.Int("e", 0, "episode number, required for -a "+tvEpisodesActionsHelp)
	if err := flagSet.Parse(args); err != nil {
		return err
	}

	if tvEpisodesSessionActions[*action] {
		if err := cli.HandleToken(fs, config, client, options); err != nil {
			return fmt.Errorf("tv-episodes: %w", err)
		}
	}

	var handler handlers.Handler
	var params []string
	switch *action {
	case "":
		return fmt.Errorf("tv-episodes: -a is required (action: %s)", tvEpisodesActionsHelp)
	case "details":
		if *seriesID == 0 {
			return fmt.Errorf("tv-episodes: -i <series_id> is required for -a details")
		}
		handler = handlers.TVEpisodesDetailsHandler{SeriesID: *seriesID, SeasonNumber: *seasonNumber, EpisodeNumber: *episodeNumber}
		params = []string{fmt.Sprintf("id-%d", *seriesID), fmt.Sprintf("season-%d", *seasonNumber), fmt.Sprintf("episode-%d", *episodeNumber)}
	case "account-states":
		if *seriesID == 0 {
			return fmt.Errorf("tv-episodes: -i <series_id> is required for -a account-states")
		}
		handler = handlers.TVEpisodesAccountStatesHandler{SeriesID: *seriesID, SeasonNumber: *seasonNumber, EpisodeNumber: *episodeNumber, SessionID: options.Session.SessionID}
		params = []string{fmt.Sprintf("id-%d", *seriesID), fmt.Sprintf("season-%d", *seasonNumber), fmt.Sprintf("episode-%d", *episodeNumber)}
	default:
		return fmt.Errorf("tv-episodes: unknown action %q (action: %s)", *action, tvEpisodesActionsHelp)
	}

	result, err := handler.Handle(context.Background(), client)
	if err != nil {
		if !retried && tvEpisodesSessionActions[*action] && isSessionInvalid(err) {
			if newSession, refreshErr := refreshSession(fs, config, client); refreshErr == nil {
				options.Session = newSession
				return execTVEpisodesAttempt(fs, client, config, options, args, true)
			}
		}
		return err
	}

	return writeResult(fs, config, "tv-episodes", *action, result, params...)
}
