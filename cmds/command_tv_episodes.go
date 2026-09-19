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
	"add-rating":     true,
	"delete-rating":  true,
}

// tvEpisodesActionsHelp lists every tv-episodes action, shared between the
// -a flag's usage string and the "-a is required" error so both stay in
// sync.
const tvEpisodesActionsHelp = "details, account-states, add-rating, credits, delete-rating, external-ids, images, translations, videos"

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
	language := flagSet.String("language", "", "ISO 639-1 language code, used by -a credits, images, videos")
	includeImageLanguage := flagSet.String("include-image-language", "", "comma-separated language codes, used by -a images")
	value := flagSet.Float64("value", 0, "rating value (0.5-10.0, in 0.5 increments), required for -a add-rating")
	guestSessionID := flagSet.String("guest-session-id", "", "guest session id (alternative to account session, used by -a add-rating): omit to use the one cached by `guest-sessions -a create`, if no account session is set up (or pass -guest)")
	guest := flagSet.Bool("guest", false, "rate as the guest session cached by `guest-sessions -a create`, ignoring the account session (used by -a add-rating)")
	if err := flagSet.Parse(args); err != nil {
		return err
	}

	ratingGuestSessionID, err := resolveRatingGuestSessionID(*action, *guestSessionID, *guest, options)
	if err != nil {
		return fmt.Errorf("tv-episodes: %w", err)
	}
	usingGuestSession := ratingGuestSessionID != ""

	if tvEpisodesSessionActions[*action] && !usingGuestSession {
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
	case "credits":
		if *seriesID == 0 {
			return fmt.Errorf("tv-episodes: -i <series_id> is required for -a credits")
		}
		handler = handlers.TVEpisodesCreditsHandler{SeriesID: *seriesID, SeasonNumber: *seasonNumber, EpisodeNumber: *episodeNumber, Language: *language}
		params = []string{fmt.Sprintf("id-%d", *seriesID), fmt.Sprintf("season-%d", *seasonNumber), fmt.Sprintf("episode-%d", *episodeNumber)}
	case "external-ids":
		if *seriesID == 0 {
			return fmt.Errorf("tv-episodes: -i <series_id> is required for -a external-ids")
		}
		handler = handlers.TVEpisodesExternalIDsHandler{SeriesID: *seriesID, SeasonNumber: *seasonNumber, EpisodeNumber: *episodeNumber}
		params = []string{fmt.Sprintf("id-%d", *seriesID), fmt.Sprintf("season-%d", *seasonNumber), fmt.Sprintf("episode-%d", *episodeNumber)}
	case "images":
		if *seriesID == 0 {
			return fmt.Errorf("tv-episodes: -i <series_id> is required for -a images")
		}
		handler = handlers.TVEpisodesImagesHandler{SeriesID: *seriesID, SeasonNumber: *seasonNumber, EpisodeNumber: *episodeNumber, Language: *language, IncludeImageLanguage: *includeImageLanguage}
		params = []string{fmt.Sprintf("id-%d", *seriesID), fmt.Sprintf("season-%d", *seasonNumber), fmt.Sprintf("episode-%d", *episodeNumber)}
	case "translations":
		if *seriesID == 0 {
			return fmt.Errorf("tv-episodes: -i <series_id> is required for -a translations")
		}
		handler = handlers.TVEpisodesTranslationsHandler{SeriesID: *seriesID, SeasonNumber: *seasonNumber, EpisodeNumber: *episodeNumber}
		params = []string{fmt.Sprintf("id-%d", *seriesID), fmt.Sprintf("season-%d", *seasonNumber), fmt.Sprintf("episode-%d", *episodeNumber)}
	case "videos":
		if *seriesID == 0 {
			return fmt.Errorf("tv-episodes: -i <series_id> is required for -a videos")
		}
		handler = handlers.TVEpisodesVideosHandler{SeriesID: *seriesID, SeasonNumber: *seasonNumber, EpisodeNumber: *episodeNumber, Language: *language}
		params = []string{fmt.Sprintf("id-%d", *seriesID), fmt.Sprintf("season-%d", *seasonNumber), fmt.Sprintf("episode-%d", *episodeNumber)}
	case "add-rating":
		if *seriesID == 0 {
			return fmt.Errorf("tv-episodes: -i <series_id> is required for -a add-rating")
		}
		if *value == 0 {
			return fmt.Errorf("tv-episodes: -value <rating> is required for -a add-rating")
		}
		sessionID := ""
		if !usingGuestSession {
			sessionID = options.Session.SessionID
		}
		handler = handlers.TVEpisodesAddRatingHandler{SeriesID: *seriesID, SeasonNumber: *seasonNumber, EpisodeNumber: *episodeNumber, SessionID: sessionID, GuestSessionID: ratingGuestSessionID, Value: *value}
		params = []string{fmt.Sprintf("id-%d", *seriesID), fmt.Sprintf("season-%d", *seasonNumber), fmt.Sprintf("episode-%d", *episodeNumber)}
	case "delete-rating":
		if *seriesID == 0 {
			return fmt.Errorf("tv-episodes: -i <series_id> is required for -a delete-rating")
		}
		handler = handlers.TVEpisodesDeleteRatingHandler{SeriesID: *seriesID, SeasonNumber: *seasonNumber, EpisodeNumber: *episodeNumber, SessionID: options.Session.SessionID}
		params = []string{fmt.Sprintf("id-%d", *seriesID), fmt.Sprintf("season-%d", *seasonNumber), fmt.Sprintf("episode-%d", *episodeNumber)}
	default:
		return fmt.Errorf("tv-episodes: unknown action %q (action: %s)", *action, tvEpisodesActionsHelp)
	}

	result, err := handler.Handle(context.Background(), client)
	if err != nil {
		if !retried && tvEpisodesSessionActions[*action] && !usingGuestSession && isSessionInvalid(err) {
			if newSession, refreshErr := refreshSession(fs, config, client); refreshErr == nil {
				options.Session = newSession
				return execTVEpisodesAttempt(fs, client, config, options, args, true)
			}
		}
		return err
	}

	return writeResult(fs, config, "tv-episodes", *action, result, params...)
}
