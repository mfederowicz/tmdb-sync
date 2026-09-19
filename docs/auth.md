# `auth` (v4 only)

TMDB v4 login: create a user access token in the browser and cache it for the v4 `account` and
`lists` actions. Every action needs `-v4` (without it the command errors with a hint), and a
`read_access_token` in the config — `api_key` alone is not enough for v4.

| action          | example                                                             | output file                |
|-----------------|---------------------------------------------------------------------|----------------------------|
| `login`         | `tmdb-sync auth -v4 -a login`                                        | none (token cached)        |
| `request-token` | `tmdb-sync auth -v4 -a request-token [-redirect-to <url>]`           | `auth_request-token_v4.json` |
| `access-token`  | `tmdb-sync auth -v4 -a access-token -request-token <approved token>` | none (token cached)        |
| `logout`        | `tmdb-sync auth -v4 -a logout`                                       | `auth_logout_v4.json`      |
| `logout` (v3)   | `tmdb-sync auth -a logout`                                           | `auth_logout_v3.json`      |

Notes:
- `login` runs the whole flow: create a request token, open
  `https://www.themoviedb.org/auth/access?request_token=…` in the browser, wait for Enter, then
  exchange the approved token for a user access token.
- `request-token` and `access-token` are the two halves of `login`, for scripted use: approve the
  token from `request-token` in the browser yourself, then pass it to `access-token`.
- The user access token and the v4 `account_id` (`account_object_id`) are saved to
  `~/.config/tmdb-sync/access_token.json` (`access_token_path` in the TOML config), mode `0600`. They
  are never written to `output_dir`.
- `logout` invalidates the cached token at TMDB and deletes the file; it errors if no token is
  cached. Log in again with `-a login`.
- v3 `logout` (no `-v4`, the only v3 action of this module) invalidates the cached v3 session at TMDB
  and deletes `session_path` and `account_path`; it errors if no session is cached. The next 🔒
  action logs in again on demand. TMDB has no equivalent for guest sessions, they just expire.
- `au` is the module's `Abbrev` — `tmdb-sync au -v4 -a login` also works.
