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

### Changed

### Fixed

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
