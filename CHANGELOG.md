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
