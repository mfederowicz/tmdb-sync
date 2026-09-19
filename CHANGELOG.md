# Changelog

All notable changes to this project are documented here, in the
[Keep a Changelog](https://keepachangelog.com/en/1.1.0/) format.

Versioning follows [SemVer](https://semver.org/), applied loosely while the project is pre-1.0:

- **Minor** (`0.X.0`) — a module lands (new endpoints/commands), or other user-facing behavior
  changes (e.g. authentication/rate-limit handling).
- **Patch** (`0.1.X`) — bug fixes, CI/tooling changes, docs-only changes, refactors with no
  behavior change.
- **Major** stays `0` until v3 coverage is near-complete and the CLI surface is considered stable.

A version is tagged once a module (or a meaningful fix) is done — there's no fixed release
schedule.

## [Unreleased]

### Added

- `auth -a logout` (v3, no `-v4`): invalidates the cached v3 session at TMDB and deletes the
  session and account cache files; errors if no session is cached. The result is written as
  `auth_logout_v3.json`. `auth -v4 -a logout` is unchanged.

### Changed

- The cached v3 session, account, guest-session and v4 access-token files are now written
  owner-only (`0600`), and an existing world-readable file is tightened on the next write.

### Fixed

- `account` with a missing or unknown `-a` no longer starts the browser login flow before reporting
  the error.
- A 401 on a `-v4` `account` or `lists` call no longer triggers the v3 login flow (which could not
  fix it); the error now suggests `auth -v4 -a login`.
- When the automatic re-login after a 401 fails, the error now includes why, instead of only the
  original 401.
- `pages_limit = 0` in the config file now means "all pages", as documented, instead of being
  ignored in favour of the default of 10. A negative `pages_limit` is rejected with an error.
- `output_dir` now expands a leading `~`, instead of creating a directory literally named `~`.
  `~other/...` is no longer rewritten to `<home>/other/...` in any config path.
- Login no longer fails on a fresh machine: the parent directory of the session, account,
  guest-session and access-token cache files is created if missing.
- API client: response bodies of failed (non-2xx) requests are now closed, so failed calls no
  longer leak connections.
- A 429 error no longer prints the request URL with the `api_key` query parameter in clear.
- A 429 with a `Retry-After` header now arms the client's rate-limit guard, so later requests in
  the same run stop instead of hammering the API.
- Error responses with a non-JSON body (e.g. a proxy's HTML page) now keep their HTTP status code,
  so a 401 still triggers the session re-login.
- v4 requests made with a user access token no longer require `read_access_token` in the config.

## [0.27.0] - 2026-09-19

### Added

- `account -v4` (needs `-v4`, `read_access_token` and a cached user token from `auth -v4 -a login`;
  no v3 session and no `-i`): `-a lists`, `favorite-movies`, `favorite-tv`, `rated-movies`,
  `rated-tv`, `watchlist-movies`, `watchlist-tv`, and the v4-only `recommended-movies` and
  `recommended-tv`. Results are written as `account_<action>_v4.json`.
- `account -v4` accepts `-language <code>` (favorites, rated, recommended, watchlist) and
  `-sort-by created_at.asc|created_at.desc` (favorites, rated, watchlist); using one where the
  endpoint doesn't take it, or without `-v4`, is an error.
- `lists -v4` (needs `-v4` and `read_access_token`; everything except `details`
  also needs a cached user token from `auth -v4 -a login`): `-a details` (walks every item page,
  `-language`, `-sort-by`, `-pages-limit`), `item-status` (`-media-type movie|tv`, `-media-id`),
  and 🔒 `create` (`-name`, `-language`, `-country`, `-description`, `-public`), `update`
  (`-name`, `-description`, `-public[=false]`, `-sort-by`, `-backdrop-path`), `delete`, `clear`,
  and the batch actions `add-items`, `update-items` (comment per item) and `remove-items`, which
  take repeatable `-item movie:<id>` / `-item tv:<id>[:<comment>]`. Results are written as
  `lists_<action>_..._v4.json`.

### Changed

- Internal refactor, no behavior change: the v4 auth service methods, tests and types now live in
  the v3 auth files (`Client.Auth.*V4`, `str.RequestTokenV4`/`AccessTokenV4`) instead of separate
  `*_v4` files, matching how v4 account is laid out.

## [0.26.0] - 2026-09-19

### Added

- `auth` module (v4 only, needs `-v4` and `read_access_token`): `-a request-token`,
  `-a access-token -request-token <token>`, `-a login` (browser approval flow), and `-a logout`.
  The user access token and v4 `account_id` are cached (mode `0600`) at the new
  `access_token_path` config key (default `~/.config/tmdb-sync/access_token.json`), never in
  `output_dir`.
- v4 client plumbing: `NewRequestV4` targets `https://api.themoviedb.org/4/` with the bearer token
  only (no `api_key`), and shared `-v3`/`-v4` flag resolution for the upcoming v4 `account`/`lists`.

## [0.25.0] - 2026-09-19

### Added

- `-d`/`-debug` now also prints the outgoing request headers (sorted, with the `Authorization`
  value redacted).
- `movies`, `tv`, and `tv-episodes` `-a add-rating` accept a new `-guest` flag that rates as the
  guest session cached by `guest-sessions -a create`, ignoring any account session; errors when no
  guest session is cached.

### Changed

- `-d`/`-debug` request URL line now includes the `api_key` query parameter, shown as `REDACTED`.

## [0.24.0] - 2026-09-18

### Added

- `-d`/`-debug` CLI flags and `debug` config option: print the full resolved request URL (before
  the API key is appended) to stdout for each outbound API call.
- `movies`, `tv`, and `tv-episodes` `-a add-rating` accept a new `-guest-session-id` flag, sending
  TMDB's `guest_session_id` query parameter instead of `session_id` so a rating can be attached to
  a guest session; falls back to the id cached by `guest-sessions -a create` when no account
  session is configured.

### Fixed

- `guest-sessions -a rated-movies/rated-tv/rated-tv-episodes` always returned a 404, since there
  was no way to actually record a rating against a guest session (rating endpoints only ever sent
  `session_id`, never `guest_session_id`).

## [0.23.0] - 2026-09-18

### Added

- `guest-sessions` module: `create`, `rated-movies`, `rated-tv`, `rated-tv-episodes`.

### Changed

- `authentication`'s `CreateGuestSession` is now wired in via `guest-sessions -a create`, which
  caches the returned id to `~/.config/tmdb-sync/guest_session.json` (`guest_session_path` in
  config); `-i` is now optional on `guest-sessions`' `rated-*` actions, falling back to that cache.

## [0.22.0] - 2026-09-18

### Added

- `watch-providers` module: `available-regions`, `movie-providers`, `tv-providers`.

## [0.21.0] - 2026-09-18

### Added

- `tv-episode-groups` module: `details`.

## [0.20.0] - 2026-09-18

### Added

- `tv-episodes` module: `details`, `account-states`, `credits`, `external-ids`, `images`,
  `translations`, `videos`, `add-rating`, `delete-rating`.

## [0.19.0] - 2026-09-18

### Added

- `tv-seasons` module: `details`, `account-states`, `aggregate-credits`, `credits`,
  `external-ids`, `images`, `translations`, `videos`.

### Changed

- Wrapped two overly long action-help consts in `cmds/command_tv.go` to satisfy revive's
  line-length lint rule.

## [0.18.0] - 2026-09-18

### Added

- `tv` module: `details`, `account-states`, `aggregate-credits`, `alternative-titles`,
  `content-ratings`, `credits`, `episode-groups`, `external-ids`, `images`, `keywords`, `latest`,
  `lists`, `recommendations`, `reviews`, `screened-theatrically`, `similar`, `translations`,
  `videos`, `watch-providers`, `popular`, `top-rated`, `on-the-air`, `airing-today`,
  `add-rating`, and `delete-rating` actions.

## [0.17.0] - 2026-09-17

### Added

- `trending` module: `all`, `movie`, `tv`, and `person` actions, with `-w <day|week>` time-window
  support.

## [0.16.0] - 2026-09-17

### Added

- `search` module: `collections`, `companies`, `keywords`, `movies`, `multi`, `people`, and `tv`
  actions.

## [0.15.0] - 2026-09-17

### Added

- `reviews` module: `details` action.

## [0.14.0] - 2026-09-17

### Added

- `people` module: `details`, `combined-credits`, `external-ids`, `images`, `latest`,
  `movie-credits`, `popular` (paginated), `tv-credits`, `translations` actions.

## [0.13.0] - 2026-09-17

### Added

- `networks` module: `details`, `alternative-names`, `images` actions.

## [0.12.1] - 2026-09-17

### Fixed

- `movies`: fixed two `revive` line-length warnings on the `-a`/`-i` flag usage strings and the
  "-a is required" error (extracted into shared consts), and an unused-receiver warning on
  `MoviesLatestHandler.Handle`. No behavior change.

## [0.12.0] - 2026-09-17

### Added

- `movies` module, complete: `account-states` 🔒, `alternative-titles`, `credits`, `external-ids`,
  `images`, `keywords`, `latest`, `lists`, `now-playing`, `recommendations`, `release-dates`,
  `reviews`, `similar`, `top-rated`, `translations`, `upcoming`, `videos`, `watch-providers`,
  `add-rating` 🔒, `delete-rating` 🔒 — completes the module's full 20/20 endpoint coverage
  (`details` and `popular` already existed). `add-rating`/`delete-rating` use the existing
  v3-session plumbing directly, without needing the guest-session flow deferred for the future
  rating/guest-sessions checkpoint.

## [0.11.0] - 2026-09-17

### Added

- `lists` module, complete: `details` and `item-status` (public reads); `create`, `add-movie`,
  `remove-movie`, `clear`, `delete` (🔒, requiring a v3 session, with the same transparent
  re-login on a stale session as the `account` module).

### Changed

- Replaced the fixed per-PR module-size threshold with asking the user how many endpoints to
  tackle per session/branch (see `CLAUDE.md`), since available token budget varies per session.

## [0.10.0] - 2026-09-17

### Added

- `keywords` module: `details` action (`GET /keyword/{keyword_id}`), returning a keyword's id
  and name. TMDB's "Movies by keyword" endpoint is deprecated in favor of `discover -a movie`
  with `with_keywords` and is not implemented.

## [0.9.0] - 2026-09-17

### Added

- `genres` module: `movie` and `tv` actions (`GET /genre/movie/list`, `GET /genre/tv/list`),
  returning TMDB's official genre lists.

## [0.8.0] - 2026-09-17

### Added

- `find` module: `by-id` action (`GET /find/{external_id}`), looking up an external id (IMDb,
  TVDB, ...) and returning matching movie/tv/person/episode/season results.

## [0.7.0] - 2026-09-17

### Added

- `discover` module: `movie` and `tv` actions (`GET /discover/movie`, `GET /discover/tv`) with
  filters for sort order, language/region, release/air year, genres, vote average range, watch
  providers, and watch region.

### Changed

- `uri.AddQuery`'s query-string builder now supports float-valued filter fields.

## [0.6.0] - 2026-09-17

### Added

- `credits` module: `details` action (`GET /credit/{credit_id}`).

## [0.5.0] - 2026-09-17

### Added

- `collections` module: `details`, `images`, `translations` actions.
- `companies` module: `details`, `alternative-names`, `images` actions.
- Shared `str.Image`/`str.Images` and `str.Translation`/`str.Translations` response types, and
  `uri.ImagesOptions`, reusable by future per-resource image/translation endpoints.

### Fixed

- `str.CollectionPart.GenreIds` renamed to `GenreIDs` to follow Go naming conventions for
  initialisms.

## [0.4.0] - 2026-09-17

### Added

- `changes` module: `movie`, `tv`, `person` actions, with optional `-start-date`/`-end-date`
  filters and `-pages-limit` pagination.
- `main.go` validates the configured `api_key`/`read_access_token` at startup and exits with a
  clear error before running any command if it's invalid.
- `account` actions now automatically re-run the login flow and retry once on an HTTP 401
  (expired/revoked session), instead of failing outright.

## [0.3.0] - 2026-09-17

### Added

- `account` 🔒 module: `rated-tv-episodes`, `watchlist-movies`, `watchlist-tv` actions — completes
  the module's full 11/11 endpoint coverage.

## [0.2.0] - 2026-09-17

### Added

- `account` 🔒 module: `details`, `add-watchlist`, `add-favorite`, `favorite-movies`,
  `favorite-tv`, `lists`, `rated-movies`, `rated-tv` actions.
- `-i <account_id>` is optional on every `account` action: omitted, `details` self-resolves the
  account via the session and caches it to `account_path` (default
  `~/.config/tmdb-sync/account.json`), so later invocations don't need to repeat it.
- `docs/account.md` reference page.

### Changed

- `consts.MediaTypeMovie`/`MediaTypeTV` added and reused in place of repeated string literals.

## [0.1.0] - 2026-09-16

### Added

- `configuration` module: `countries`, `jobs`, `languages`, `primary-translations`, `timezones`
  actions (in addition to the existing `details`).
- `-pages-limit` flag on `movies -a popular` to override `pages_limit` per-invocation.
- `docs/certifications.md`, `docs/configuration.md` reference pages.

### Changed

- Every module now requires an explicit `-a`; a missing/unknown action is an error instead of
  silently falling back to a default action.
- Module resolution now requires an exact module name or abbreviation match instead of an
  unambiguous prefix.
- `main` exits with status 1 on any command failure (unknown/ambiguous module, missing `-a`, or a
  module-level error) instead of always exiting 0.

### Fixed

- Removed the `authentication` CLI module (command, handlers, and its doc page); it had no real
  standalone use since `main` never auto-triggers the session flow — a future 🔒 module calls the
  internal auth plumbing (`internal/auth_service.go`, `cli/session.go`) directly instead.
