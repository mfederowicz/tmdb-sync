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
