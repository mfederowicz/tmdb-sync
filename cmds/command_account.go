package cmds

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"net/http"
	"slices"
	"strings"

	"github.com/mfederowicz/tmdb-sync/cfg"
	"github.com/mfederowicz/tmdb-sync/cli"
	"github.com/mfederowicz/tmdb-sync/consts"
	"github.com/mfederowicz/tmdb-sync/handlers"
	"github.com/mfederowicz/tmdb-sync/internal"
	"github.com/spf13/afero"

	"github.com/mfederowicz/tmdb-sync/str"
)

// refreshSession is a var (not a plain function call) so tests can stub it
// without triggering the real browser-approval flow.
var refreshSession = cli.CreateSessionInteractively

// isSessionInvalid reports whether err is a TMDB auth failure (HTTP 401),
// which is the only signal TMDB gives for an expired/revoked session_id —
// there's no dedicated "check my session" endpoint.
func isSessionInvalid(err error) bool {
	var errResp *str.ErrorResponse
	return errors.As(err, &errResp) && errResp.StatusCode == http.StatusUnauthorized
}

// accountV4Actions are the `account` actions implemented for -v4.
var accountV4Actions = []string{"lists", "favorite-movies", "favorite-tv", "rated-movies", "rated-tv"}

// AccountCmd is the "account" 🔒 module. Every action requires a v3 session,
// established on demand via cli.HandleToken.
var AccountCmd = &Command{
	Name:   "account",
	Abbrev: "acc",
	Short:  "🔒 account details, favorites, watchlist, lists, rated",
	Exec:   execAccount,
}

func execAccount(fs afero.Fs, client *internal.Client, config *cfg.Config, options *str.Options, args []string) error {
	return execAccountAttempt(fs, client, config, options, args, false)
}

// execAccountAttempt is execAccount's body, split out so a stale session
// (detected via isSessionInvalid) can trigger one transparent re-login and
// retry instead of failing outright.
func execAccountAttempt(fs afero.Fs, client *internal.Client, config *cfg.Config, options *str.Options, args []string, retried bool) error {
	flagSet := flag.NewFlagSet("account", flag.ContinueOnError)
	action := flagSet.String("a", "", "action: details, add-watchlist, add-favorite, favorite-movies, favorite-tv, lists, rated-movies, rated-tv, rated-tv-episodes, watchlist-movies, watchlist-tv (required)")
	accountID := flagSet.Int64("i", 0, "account id (optional: omit to self-resolve via the session and cache it)")
	mediaType := flagSet.String("media-type", "", "media type: movie, tv - required for -a add-watchlist, add-favorite")
	mediaID := flagSet.Int64("media-id", 0, "movie/tv id - required for -a add-watchlist, add-favorite")
	watchlist := flagSet.Bool("watchlist", true, "used by -a add-watchlist: true adds, false removes")
	favorite := flagSet.Bool("favorite", true, "used by -a add-favorite: true adds, false removes")
	v3 := flagSet.Bool("v3", false, "use the v3 API (default)")
	v4 := flagSet.Bool("v4", false, "use the v4 API: needs `auth -v4 -a login` first; actions: lists, favorite-movies, favorite-tv, rated-movies, rated-tv")
	pagesLimit := flagSet.Int("pages-limit", config.PagesLimit, "pages limit, used by -a favorite-movies, favorite-tv, lists, rated-movies, rated-tv, rated-tv-episodes, watchlist-movies, watchlist-tv (default: pages_limit from config, 0 = unlimited)")
	if err := flagSet.Parse(args); err != nil {
		return err
	}

	version, err := resolveAPIVersion(*v3, *v4)
	if err != nil {
		return fmt.Errorf("account: %w", err)
	}
	v4Mode := version == apiV4

	var id int64
	if v4Mode {
		if *action != "" && !slices.Contains(accountV4Actions, *action) {
			return fmt.Errorf("account: action %q is not available with -v4 (v4 actions: %s)", *action, strings.Join(accountV4Actions, ", "))
		}
		if *action != "" && (options.AccessTokenV4 == nil || options.AccessTokenV4.AccessToken == "" || options.AccessTokenV4.AccountID == "") {
			return fmt.Errorf("account: no v4 access token cached at %s, run `auth -v4 -a login` first", config.AccessTokenPath)
		}
	} else {
		if err := cli.HandleToken(fs, config, client, options); err != nil {
			return fmt.Errorf("account: %w", err)
		}
		id = *accountID
		if id == 0 && options.Account != nil && options.Account.ID != 0 {
			id = options.Account.ID
		}
	}

	var handler handlers.Handler
	var params []string
	switch *action {
	case "":
		return fmt.Errorf("account: -a is required (action: details, add-watchlist, add-favorite, favorite-movies, favorite-tv, lists, rated-movies, rated-tv, rated-tv-episodes, watchlist-movies, watchlist-tv)")
	case "details":
		handler = handlers.AccountDetailsHandler{AccountID: id, SessionID: options.Session.SessionID}
		if id != 0 {
			params = []string{fmt.Sprintf("id-%d", id)}
		}
	case "add-watchlist":
		if id == 0 {
			return fmt.Errorf("account: -i <account_id> is required for -a add-watchlist (or run -a details once to cache it)")
		}
		if *mediaType != consts.MediaTypeMovie && *mediaType != consts.MediaTypeTV {
			return fmt.Errorf("account: -media-type must be movie or tv for -a add-watchlist")
		}
		if *mediaID == 0 {
			return fmt.Errorf("account: -media-id is required for -a add-watchlist")
		}
		handler = handlers.AccountWatchlistHandler{
			AccountID: id,
			SessionID: options.Session.SessionID,
			MediaType: *mediaType,
			MediaID:   *mediaID,
			Watchlist: *watchlist,
		}
		params = []string{fmt.Sprintf("id-%d", id), *mediaType, fmt.Sprintf("media-%d", *mediaID)}
	case "add-favorite":
		if id == 0 {
			return fmt.Errorf("account: -i <account_id> is required for -a add-favorite (or run -a details once to cache it)")
		}
		if *mediaType != consts.MediaTypeMovie && *mediaType != consts.MediaTypeTV {
			return fmt.Errorf("account: -media-type must be movie or tv for -a add-favorite")
		}
		if *mediaID == 0 {
			return fmt.Errorf("account: -media-id is required for -a add-favorite")
		}
		handler = handlers.AccountFavoriteHandler{
			AccountID: id,
			SessionID: options.Session.SessionID,
			MediaType: *mediaType,
			MediaID:   *mediaID,
			Favorite:  *favorite,
		}
		params = []string{fmt.Sprintf("id-%d", id), *mediaType, fmt.Sprintf("media-%d", *mediaID)}
	case "favorite-movies":
		if v4Mode {
			handler = handlers.AccountFavoriteMoviesHandler{V4: true, AccessToken: options.AccessTokenV4.AccessToken, V4AccountID: options.AccessTokenV4.AccountID, PagesLimit: *pagesLimit}
			params = []string{"v4"}
			break
		}
		if id == 0 {
			return fmt.Errorf("account: -i <account_id> is required for -a favorite-movies (or run -a details once to cache it)")
		}
		handler = handlers.AccountFavoriteMoviesHandler{AccountID: id, SessionID: options.Session.SessionID, PagesLimit: *pagesLimit}
		params = []string{fmt.Sprintf("id-%d", id)}
	case "favorite-tv":
		if v4Mode {
			handler = handlers.AccountFavoriteTVHandler{V4: true, AccessToken: options.AccessTokenV4.AccessToken, V4AccountID: options.AccessTokenV4.AccountID, PagesLimit: *pagesLimit}
			params = []string{"v4"}
			break
		}
		if id == 0 {
			return fmt.Errorf("account: -i <account_id> is required for -a favorite-tv (or run -a details once to cache it)")
		}
		handler = handlers.AccountFavoriteTVHandler{AccountID: id, SessionID: options.Session.SessionID, PagesLimit: *pagesLimit}
		params = []string{fmt.Sprintf("id-%d", id)}
	case "lists":
		if v4Mode {
			handler = handlers.AccountListsHandler{V4: true, AccessToken: options.AccessTokenV4.AccessToken, V4AccountID: options.AccessTokenV4.AccountID, PagesLimit: *pagesLimit}
			params = []string{"v4"}
			break
		}
		if id == 0 {
			return fmt.Errorf("account: -i <account_id> is required for -a lists (or run -a details once to cache it)")
		}
		handler = handlers.AccountListsHandler{AccountID: id, SessionID: options.Session.SessionID, PagesLimit: *pagesLimit}
		params = []string{fmt.Sprintf("id-%d", id)}
	case "rated-movies":
		if v4Mode {
			handler = handlers.AccountRatedMoviesHandler{V4: true, AccessToken: options.AccessTokenV4.AccessToken, V4AccountID: options.AccessTokenV4.AccountID, PagesLimit: *pagesLimit}
			params = []string{"v4"}
			break
		}
		if id == 0 {
			return fmt.Errorf("account: -i <account_id> is required for -a rated-movies (or run -a details once to cache it)")
		}
		handler = handlers.AccountRatedMoviesHandler{AccountID: id, SessionID: options.Session.SessionID, PagesLimit: *pagesLimit}
		params = []string{fmt.Sprintf("id-%d", id)}
	case "rated-tv":
		if v4Mode {
			handler = handlers.AccountRatedTVHandler{V4: true, AccessToken: options.AccessTokenV4.AccessToken, V4AccountID: options.AccessTokenV4.AccountID, PagesLimit: *pagesLimit}
			params = []string{"v4"}
			break
		}
		if id == 0 {
			return fmt.Errorf("account: -i <account_id> is required for -a rated-tv (or run -a details once to cache it)")
		}
		handler = handlers.AccountRatedTVHandler{AccountID: id, SessionID: options.Session.SessionID, PagesLimit: *pagesLimit}
		params = []string{fmt.Sprintf("id-%d", id)}
	case "rated-tv-episodes":
		if id == 0 {
			return fmt.Errorf("account: -i <account_id> is required for -a rated-tv-episodes (or run -a details once to cache it)")
		}
		handler = handlers.AccountRatedTVEpisodesHandler{AccountID: id, SessionID: options.Session.SessionID, PagesLimit: *pagesLimit}
		params = []string{fmt.Sprintf("id-%d", id)}
	case "watchlist-movies":
		if id == 0 {
			return fmt.Errorf("account: -i <account_id> is required for -a watchlist-movies (or run -a details once to cache it)")
		}
		handler = handlers.AccountWatchlistMoviesHandler{AccountID: id, SessionID: options.Session.SessionID, PagesLimit: *pagesLimit}
		params = []string{fmt.Sprintf("id-%d", id)}
	case "watchlist-tv":
		if id == 0 {
			return fmt.Errorf("account: -i <account_id> is required for -a watchlist-tv (or run -a details once to cache it)")
		}
		handler = handlers.AccountWatchlistTVHandler{AccountID: id, SessionID: options.Session.SessionID, PagesLimit: *pagesLimit}
		params = []string{fmt.Sprintf("id-%d", id)}
	default:
		return fmt.Errorf("account: unknown action %q", *action)
	}

	result, err := handler.Handle(context.Background(), client)
	if err != nil {
		if !retried && isSessionInvalid(err) {
			if newSession, refreshErr := refreshSession(fs, config, client); refreshErr == nil {
				options.Session = newSession
				return execAccountAttempt(fs, client, config, options, args, true)
			}
		}
		return err
	}

	if account, ok := result.(*str.Account); ok {
		if err := cfg.WriteAccount(fs, config.AccountPath, account); err != nil {
			return fmt.Errorf("account: cache account details: %w", err)
		}
		if len(params) == 0 {
			params = []string{fmt.Sprintf("id-%d", account.ID)}
		}
	}

	return writeResult(fs, config, "account", *action, result, params...)
}
