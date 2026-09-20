# tmdb-sync

A Go CLI for [The Movie Database (TMDB)](https://www.themoviedb.org) API — modular, TOML-configured,
covering the TMDB v3 API and the v4 auth, account and lists endpoints.

## Status

✅ **Feature-complete**: every v3 and v4 endpoint is implemented, except the TMDB-deprecated
`GET /keyword/{id}/movies`. Modules:

- [`account`](docs/account.md) ✅ — 🔒 details, favorites, watchlist, lists, rated (v3 session or `-v4`),
  plus v4 recommendations.
- [`auth`](docs/auth.md) ✅ — v4 login: request token, access token, logout; v3 logout.
- [`certifications`](docs/certifications.md) ✅ — movie/tv certifications, by country.
- [`changes`](docs/changes.md) ✅ — movie/tv/person change lists.
- [`collections`](docs/collections.md) ✅ — collection details, images, translations.
- [`companies`](docs/companies.md) ✅ — company details, alternative names, images.
- [`configuration`](docs/configuration.md) ✅ — image base URLs/sizes, countries, jobs, languages,
  primary translations, timezones.
- [`credits`](docs/credits.md) ✅ — credit details.
- [`discover`](docs/discover.md) ✅ — discover movies/TV shows by filter.
- [`find`](docs/find.md) ✅ — find TMDB movies/TV/people by an external id (IMDb, TVDB, ...).
- [`genres`](docs/genres.md) ✅ — official movie/tv genre lists.
- [`guest-sessions`](docs/guest-sessions.md) ✅ — guest session rated movies, TV shows, TV episodes.
- [`keywords`](docs/keywords.md) ✅ — keyword details.
- [`lists`](docs/lists.md) ✅ — list details, item status, and 🔒 create/add-movie/remove-movie/
  clear/delete (v3 session or `-v4`), plus v4 update and batch add/update/remove items.
- [`movies`](docs/movies.md) ✅ — full movie module: details, lists, credits, images, videos,
  watch providers, and 🔒 account-states/add-rating/delete-rating.
- [`networks`](docs/networks.md) ✅ — network details, alternative names, images.
- [`people`](docs/people.md) ✅ — person details, credits, images, translations, popular/latest.
- [`reviews`](docs/reviews.md) ✅ — review details.
- [`search`](docs/search.md) ✅ — search collections, companies, keywords, movies, multi, people,
  TV shows.
- [`trending`](docs/trending.md) ✅ — trending movies, TV shows, and people, by day or week.
- [`tv`](docs/tv.md) ✅ — full TV series module: details, lists, credits, images, videos,
  watch providers, and 🔒 account-states/add-rating/delete-rating.
- [`tv-seasons`](docs/tv-seasons.md) ✅ — TV season details, aggregate-credits, credits,
  external-ids, images, translations, videos, and 🔒 account-states.
- [`tv-episodes`](docs/tv-episodes.md) ✅ — TV episode details, credits, external-ids, images,
  translations, videos, and 🔒 account-states/add-rating/delete-rating.
- [`tv-episode-groups`](docs/tv-episode-groups.md) ✅ — TV episode group details by group id.
- [`watch-providers`](docs/watch-providers.md) ✅ — available regions, movie/tv watch provider lists.

See:

- [`docs/PRD.md`](docs/PRD.md) — goals, non-goals, scope phases, success criteria.
- [`docs/ARCHITECTURE.md`](docs/ARCHITECTURE.md) — package layout and how new modules get added.
- [`docs/API_COVERAGE.md`](docs/API_COVERAGE.md) — checklist of every TMDB v3 endpoint, tracked
  as work lands.
- [`docs/CLI.md`](docs/CLI.md) — full command/config reference.
- [`AGENTS.md`](AGENTS.md) — instructions for coding agents working in this repo.

## Goal

Near-100% coverage of the TMDB **v3** API (movies, tv, people, search, discover, trending,
account, etc.) from a single Go binary configured with one TOML file, plus TMDB **v4** auth,
account and lists endpoints (selected with `-v4`).

## Install & usage

```sh
go install github.com/mfederowicz/tmdb-sync@latest
```

Create a TOML config file (default `~/tmdb-sync.toml`, override with `-c <path>`):

```toml
# ~/tmdb-sync.toml — at least one of api_key / read_access_token is required
api_key = "your-v3-api-key"
# read_access_token = "your-v4-read-access-token"
```

Everything else is optional and has a default. A full example with every field:

```toml
api_key = "your-v3-api-key"
read_access_token = ""
session_path = "~/.config/tmdb-sync/session.json"
account_path = "~/.config/tmdb-sync/account.json"
guest_session_path = "~/.config/tmdb-sync/guest_session.json"
access_token_path = "~/.config/tmdb-sync/access_token.json"
output_dir = ""
pages_limit = 10
verbose = false
debug = false
```

### Config fields

| field | default | description |
|-------|---------|-------------|
| `api_key` | — | TMDB **v3 API key**. Sent as the `api_key` query parameter on v3 requests. |
| `read_access_token` | — | TMDB **v4 API Read Access Token**. Sent as an `Authorization: Bearer` header, and required for `-v4` commands (v4 requests never carry `api_key`). If both this and `api_key` are set, the token is used. |
| `session_path` | `~/.config/tmdb-sync/session.json` | Where the v3 user session id is cached (used by 🔒 account commands). Must end in `json`. |
| `account_path` | `~/.config/tmdb-sync/account.json` | Cache of the resolved account details (id, username) for the v3 session. |
| `guest_session_path` | `~/.config/tmdb-sync/guest_session.json` | Where a guest session is cached (`guest-sessions` module). |
| `access_token_path` | `~/.config/tmdb-sync/access_token.json` | Where the v4 **user** access token and account object id are stored by `auth -v4 -a login`. Not the same thing as `read_access_token`. |
| `output_dir` | `""` (current directory) | Directory result JSON files are written to. `~` is expanded. |
| `pages_limit` | `10` | Maximum number of pages walked by list commands using `-all`. `0` means unlimited (bounded by TMDB's `total_pages`); negative values are an error. |
| `verbose` | `false` | Verbose output (same as `-v`; the flag overrides the file). |
| `debug` | `false` | Debug output (same as `-d` / `-debug`). |
| `per_page` | `0` | Accepted in the file, but not currently used by any command. |
| `auth_version` | `"v3"` | Accepted in the file, but superseded by the per-command `-v3` / `-v4` flags. |
| `config_path` | `~/tmdb-sync.toml` | Path of the config file itself; effectively set only via `-c`. |

`~` is expanded in every path field. The session, account, guest-session and access-token files
are created automatically (with their parent directory) with owner-only permissions (`0600`),
so you normally never write them by hand. They hold credentials: keep them out of version control.

### Which credentials do I need?

| you want to | needs |
|-------------|-------|
| read public v3 data (movies, tv, search, ...) | `api_key` **or** `read_access_token` |
| v3 🔒 account commands (favorites, watchlist, ratings, lists) | the above, plus a v3 session in `session_path` |
| v4 (`-v4`) commands | `read_access_token`, plus the user token in `access_token_path` from `auth -v4 -a login` |

### Command-line flags

| flag | description |
|------|-------------|
| `-c <path>` | config file path (default `~/tmdb-sync.toml`) |
| `-v` | verbose output |
| `-d`, `-debug` | debug output |
| `-version` | print version and exit |
| `-v3` / `-v4` | API version for modules that exist in both (`account`, `lists`, `auth`); v3 is the default |

Module-specific flags (`-a <action>`, `-i`, `-s`, ...) are documented in
[`docs/CLI.md`](docs/CLI.md) and each module's page.

```sh
tmdb-sync certifications -a movie
```

See [`docs/CLI.md`](docs/CLI.md) for the full command reference, kept in sync as modules land.

## Development

```sh
make install   # go mod vendor
make build     # build the binary
make test      # go test -v -race ./...
make cover     # coverage report
make linter    # revive
```
