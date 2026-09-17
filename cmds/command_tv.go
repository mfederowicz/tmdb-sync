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
const tvActionsHelp = "details, account-states, add-rating, aggregate-credits, airing-today, alternative-titles, content-ratings, credits, delete-rating, episode-groups, external-ids, images, keywords, latest, lists, on-the-air, popular, recommendations, reviews, screened-theatrically, similar, top-rated, translations, videos, watch-providers"

// tvIDActionsHelp lists the tv actions that require -i, shared between the
// -i flag's usage string and the module doc.
const tvIDActionsHelp = "details, account-states, add-rating, aggregate-credits, alternative-titles, content-ratings, credits, delete-rating, episode-groups, external-ids, images, keywords, lists, recommendations, reviews, screened-theatrically, similar, translations, videos, watch-providers"

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
	language := flagSet.String("language", "", "ISO 639-1 language code, used by -a aggregate-credits, credits, images, lists, recommendations, reviews, similar, videos")
	includeImageLanguage := flagSet.String("include-image-language", "", "comma-separated language codes, used by -a images")
	pagesLimit := flagSet.Int("pages-limit", config.PagesLimit, "pages limit, used by -a airing-today, lists, on-the-air, popular, recommendations, reviews, similar, top-rated (default: pages_limit from config, 0 = unlimited)")
	value := flagSet.Float64("value", 0, "rating value (0.5-10.0, in 0.5 increments), required for -a add-rating")
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
	case "images":
		if *seriesID == 0 {
			return fmt.Errorf("tv: -i <series_id> is required for -a images")
		}
		handler = handlers.TVImagesHandler{SeriesID: *seriesID, Language: *language, IncludeImageLanguage: *includeImageLanguage}
		params = []string{fmt.Sprintf("id-%d", *seriesID)}
	case "keywords":
		if *seriesID == 0 {
			return fmt.Errorf("tv: -i <series_id> is required for -a keywords")
		}
		handler = handlers.TVKeywordsHandler{SeriesID: *seriesID}
		params = []string{fmt.Sprintf("id-%d", *seriesID)}
	case "latest":
		handler = handlers.TVLatestHandler{}
	case "airing-today":
		handler = handlers.TVAiringTodayHandler{PagesLimit: *pagesLimit}
	case "on-the-air":
		handler = handlers.TVOnTheAirHandler{PagesLimit: *pagesLimit}
	case "popular":
		handler = handlers.TVPopularHandler{PagesLimit: *pagesLimit}
	case "top-rated":
		handler = handlers.TVTopRatedHandler{PagesLimit: *pagesLimit}
	case "lists":
		if *seriesID == 0 {
			return fmt.Errorf("tv: -i <series_id> is required for -a lists")
		}
		handler = handlers.TVListsHandler{SeriesID: *seriesID, Language: *language, PagesLimit: *pagesLimit}
		params = []string{fmt.Sprintf("id-%d", *seriesID)}
	case "recommendations":
		if *seriesID == 0 {
			return fmt.Errorf("tv: -i <series_id> is required for -a recommendations")
		}
		handler = handlers.TVRecommendationsHandler{SeriesID: *seriesID, Language: *language, PagesLimit: *pagesLimit}
		params = []string{fmt.Sprintf("id-%d", *seriesID)}
	case "reviews":
		if *seriesID == 0 {
			return fmt.Errorf("tv: -i <series_id> is required for -a reviews")
		}
		handler = handlers.TVReviewsHandler{SeriesID: *seriesID, Language: *language, PagesLimit: *pagesLimit}
		params = []string{fmt.Sprintf("id-%d", *seriesID)}
	case "screened-theatrically":
		if *seriesID == 0 {
			return fmt.Errorf("tv: -i <series_id> is required for -a screened-theatrically")
		}
		handler = handlers.TVScreenedTheatricallyHandler{SeriesID: *seriesID}
		params = []string{fmt.Sprintf("id-%d", *seriesID)}
	case "similar":
		if *seriesID == 0 {
			return fmt.Errorf("tv: -i <series_id> is required for -a similar")
		}
		handler = handlers.TVSimilarHandler{SeriesID: *seriesID, Language: *language, PagesLimit: *pagesLimit}
		params = []string{fmt.Sprintf("id-%d", *seriesID)}
	case "translations":
		if *seriesID == 0 {
			return fmt.Errorf("tv: -i <series_id> is required for -a translations")
		}
		handler = handlers.TVTranslationsHandler{SeriesID: *seriesID}
		params = []string{fmt.Sprintf("id-%d", *seriesID)}
	case "videos":
		if *seriesID == 0 {
			return fmt.Errorf("tv: -i <series_id> is required for -a videos")
		}
		handler = handlers.TVVideosHandler{SeriesID: *seriesID, Language: *language}
		params = []string{fmt.Sprintf("id-%d", *seriesID)}
	case "watch-providers":
		if *seriesID == 0 {
			return fmt.Errorf("tv: -i <series_id> is required for -a watch-providers")
		}
		handler = handlers.TVWatchProvidersHandler{SeriesID: *seriesID}
		params = []string{fmt.Sprintf("id-%d", *seriesID)}
	case "add-rating":
		if *seriesID == 0 {
			return fmt.Errorf("tv: -i <series_id> is required for -a add-rating")
		}
		if *value == 0 {
			return fmt.Errorf("tv: -value <rating> is required for -a add-rating")
		}
		handler = handlers.TVAddRatingHandler{SeriesID: *seriesID, SessionID: options.Session.SessionID, Value: *value}
		params = []string{fmt.Sprintf("id-%d", *seriesID)}
	case "delete-rating":
		if *seriesID == 0 {
			return fmt.Errorf("tv: -i <series_id> is required for -a delete-rating")
		}
		handler = handlers.TVDeleteRatingHandler{SeriesID: *seriesID, SessionID: options.Session.SessionID}
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
