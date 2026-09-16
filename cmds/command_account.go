package cmds

import (
	"context"
	"flag"
	"fmt"

	"github.com/mfederowicz/tmdb-sync/cfg"
	"github.com/mfederowicz/tmdb-sync/cli"
	"github.com/mfederowicz/tmdb-sync/consts"
	"github.com/mfederowicz/tmdb-sync/handlers"
	"github.com/mfederowicz/tmdb-sync/internal"
	"github.com/spf13/afero"

	"github.com/mfederowicz/tmdb-sync/str"
)

// AccountCmd is the "account" 🔒 module. Every action requires a v3 session,
// established on demand via cli.HandleToken.
var AccountCmd = &Command{
	Name:   "account",
	Abbrev: "acc",
	Short:  "🔒 account details, favorites, watchlist, lists, rated",
	Exec:   execAccount,
}

func execAccount(fs afero.Fs, client *internal.Client, config *cfg.Config, options *str.Options, args []string) error {
	flagSet := flag.NewFlagSet("account", flag.ContinueOnError)
	action := flagSet.String("a", "", "action: details, add-watchlist, add-favorite, favorite-movies, favorite-tv, lists, rated-movies (required)")
	accountID := flagSet.Int64("i", 0, "account id (optional: omit to self-resolve via the session and cache it)")
	mediaType := flagSet.String("media-type", "", "media type: movie, tv - required for -a add-watchlist, add-favorite")
	mediaID := flagSet.Int64("media-id", 0, "movie/tv id - required for -a add-watchlist, add-favorite")
	watchlist := flagSet.Bool("watchlist", true, "used by -a add-watchlist: true adds, false removes")
	favorite := flagSet.Bool("favorite", true, "used by -a add-favorite: true adds, false removes")
	pagesLimit := flagSet.Int("pages-limit", config.PagesLimit, "pages limit, used by -a favorite-movies, favorite-tv, lists, rated-movies (default: pages_limit from config, 0 = unlimited)")
	if err := flagSet.Parse(args); err != nil {
		return err
	}

	if err := cli.HandleToken(fs, config, client, options); err != nil {
		return fmt.Errorf("account: %w", err)
	}

	id := *accountID
	if id == 0 && options.Account != nil && options.Account.ID != 0 {
		id = options.Account.ID
	}

	var handler handlers.Handler
	var params []string
	switch *action {
	case "":
		return fmt.Errorf("account: -a is required (action: details, add-watchlist, add-favorite, favorite-movies, favorite-tv, lists, rated-movies)")
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
		if id == 0 {
			return fmt.Errorf("account: -i <account_id> is required for -a favorite-movies (or run -a details once to cache it)")
		}
		handler = handlers.AccountFavoriteMoviesHandler{AccountID: id, SessionID: options.Session.SessionID, PagesLimit: *pagesLimit}
		params = []string{fmt.Sprintf("id-%d", id)}
	case "favorite-tv":
		if id == 0 {
			return fmt.Errorf("account: -i <account_id> is required for -a favorite-tv (or run -a details once to cache it)")
		}
		handler = handlers.AccountFavoriteTVHandler{AccountID: id, SessionID: options.Session.SessionID, PagesLimit: *pagesLimit}
		params = []string{fmt.Sprintf("id-%d", id)}
	case "lists":
		if id == 0 {
			return fmt.Errorf("account: -i <account_id> is required for -a lists (or run -a details once to cache it)")
		}
		handler = handlers.AccountListsHandler{AccountID: id, SessionID: options.Session.SessionID, PagesLimit: *pagesLimit}
		params = []string{fmt.Sprintf("id-%d", id)}
	case "rated-movies":
		if id == 0 {
			return fmt.Errorf("account: -i <account_id> is required for -a rated-movies (or run -a details once to cache it)")
		}
		handler = handlers.AccountRatedMoviesHandler{AccountID: id, SessionID: options.Session.SessionID, PagesLimit: *pagesLimit}
		params = []string{fmt.Sprintf("id-%d", id)}
	default:
		return fmt.Errorf("account: unknown action %q", *action)
	}

	result, err := handler.Handle(context.Background(), client)
	if err != nil {
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
