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

Notes:
- `add-watchlist`/`add-favorite` default to adding (`-watchlist`/`-favorite` default `true`);
  pass `-watchlist=false` or `-favorite=false` to remove instead.
- `favorite-movies`, `favorite-tv`, `lists`, `rated-movies`, `rated-tv` all walk every page up to
  `-pages-limit` (default: `pages_limit` from config, `0` = unlimited), same as `movies -a popular`.
- `acc` is the module's `Abbrev` — `tmdb-sync acc -a details` also works.
