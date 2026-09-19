# `account` 🔒

Account details, favorites, watchlist, custom lists, and rated media. Every action requires a
v3 session — the first call triggers the browser-approval flow (`cli.HandleToken`) and persists
the session to `session_path` (default `~/.config/tmdb-sync/session.json`).

`-i <account_id>` is optional everywhere: omit it and `details` self-resolves the account via the
session (TMDB's session-only `GET /account`), then caches it to `account_path` (default
`~/.config/tmdb-sync/account.json`). Later invocations of any action reuse that cached id without
repeating `-i`.

| action            | example                                                                              | output file                                  |
|-------------------|---------------------------------------------------------------------------------------|-------------------------------------------------|
| `details`         | `tmdb-sync account -a details`                                                       | `account_details_id-<account_id>.json`          |
| `add-watchlist`   | `tmdb-sync account -a add-watchlist -media-type movie -media-id 550`                | `account_add-watchlist_id-<account_id>_movie_media-550.json` |
| `add-favorite`    | `tmdb-sync account -a add-favorite -media-type movie -media-id 550`                 | `account_add-favorite_id-<account_id>_movie_media-550.json`  |
| `favorite-movies` | `tmdb-sync account -a favorite-movies`                                              | `account_favorite-movies_id-<account_id>.json`  |
| `favorite-tv`     | `tmdb-sync account -a favorite-tv`                                                   | `account_favorite-tv_id-<account_id>.json`      |
| `lists`           | `tmdb-sync account -a lists`                                                         | `account_lists_id-<account_id>.json`            |
| `rated-movies`    | `tmdb-sync account -a rated-movies`                                                  | `account_rated-movies_id-<account_id>.json`     |
| `rated-tv`        | `tmdb-sync account -a rated-tv`                                                      | `account_rated-tv_id-<account_id>.json`         |
| `rated-tv-episodes` | `tmdb-sync account -a rated-tv-episodes`                                           | `account_rated-tv-episodes_id-<account_id>.json` |
| `watchlist-movies` | `tmdb-sync account -a watchlist-movies`                                             | `account_watchlist-movies_id-<account_id>.json` |
| `watchlist-tv`    | `tmdb-sync account -a watchlist-tv`                                                  | `account_watchlist-tv_id-<account_id>.json`     |

Notes:
- `add-watchlist`/`add-favorite` default to adding (`-watchlist`/`-favorite` default `true`);
  pass `-watchlist=false` or `-favorite=false` to remove instead.
- `favorite-movies`, `favorite-tv`, `lists`, `rated-movies`, `rated-tv`, `rated-tv-episodes`,
  `watchlist-movies`, `watchlist-tv` all walk every page up to `-pages-limit` (default:
  `pages_limit` from config, `0` = unlimited), same as `movies -a popular`.
- `acc` is the module's `Abbrev` — `tmdb-sync acc -a details` also works.
- If TMDB reports the cached session as no longer valid (HTTP 401 — e.g. it expired or was
  revoked from the TMDB website), the action transparently re-runs the browser-approval login
  flow and retries once, instead of failing outright.

## v4 (`-v4`)

`account -v4 -a <action>` uses the TMDB v4 API instead. It needs `read_access_token` in the config
and a cached user access token from `tmdb-sync auth -v4 -a login` (no v3 session, no `-i`: the v4
`account_object_id` is read from the cached token file). Output files carry a `_v4` suffix and
`-pages-limit` works as in v3. Combining `-v3` and `-v4` is an error.

| action  | example                             | output file            |
|---------|-------------------------------------|------------------------|
| `lists` | `tmdb-sync account -v4 -a lists`    | `account_lists_v4.json` |
| `favorite-movies` | `tmdb-sync account -v4 -a favorite-movies` | `account_favorite-movies_v4.json` |
| `favorite-tv` | `tmdb-sync account -v4 -a favorite-tv` | `account_favorite-tv_v4.json` |
| `rated-movies` | `tmdb-sync account -v4 -a rated-movies` | `account_rated-movies_v4.json` |
| `rated-tv` | `tmdb-sync account -v4 -a rated-tv` | `account_rated-tv_v4.json` |
| `recommended-movies` | `tmdb-sync account -v4 -a recommended-movies` | `account_recommended-movies_v4.json` |

`recommended-movies` exists only in v4; without `-v4` it fails with a hint.

Note: v4 rated items carry the user's rating as `account_rating` (`{created_at, value}`) instead
of v3's flat `rating`.
