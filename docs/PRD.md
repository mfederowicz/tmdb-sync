# tmdb-sync — Product Requirements Document

## Summary

`tmdb-sync` is a Go CLI tool that talks to [The Movie Database (TMDB)](https://www.themoviedb.org)
API: TOML config, `<module> -a <action>` command dispatch, one `internal/<x>_service.go` +
`handlers/<x>_handler.go` + `str/<x>.go` per API area.

## Goals

- Near-100% coverage of the **TMDB v3 API**: every public read endpoint (movies, tv, people,
  search, discover, trending, genres, keywords, companies, collections, networks, certifications,
  watch providers, changes, reviews, credits, configuration) and every v3 session-authenticated
  account endpoint (favorites, watchlist, ratings, lists).
- A single self-contained Go binary, installable via `go install`, configured with one TOML file.
- Output that's easy to pipe/script against (JSON to stdout per command).
- An architecture where adding TMDB API coverage is mechanical: copying an existing
  service/handler/struct triple for a new endpoint, not redesigning plumbing.

## Non-goals (for v1 / this phase)

- Anything beyond TMDB's v3 and v4 APIs. v4 (auth, account, lists) was originally a later phase
  and is now implemented — see [Roadmap](#roadmap).
- A GUI, TUI, or daemon/sync-loop mode — this is a one-shot CLI, invoked per command.
- Write-heavy "sync" workflows (scrobbling, check-ins) — TMDB's account API is small
  (favorites/watchlist/ratings/lists only); no equivalent exists on TMDB.

## Target user

The tool's author, using it personally/via scripts to pull TMDB data (movie/TV metadata,
discovery, search) and manage their TMDB account (watchlist, ratings, lists) from the command
line or from other tooling that shells out to it.

## Scope phases

1. **Docs** (this phase): PRD, architecture doc, API coverage checklist, CLI reference — agree on
   shape before writing code, since the last attempt jumped to code first and had to be redone.
2. **Skeleton**: client, config, command dispatch, v3 session auth plumbing (no account endpoints
   wired to commands yet) — one smoke-test module (`configuration`) end to end.
3. **v3 read-only coverage**: movies, tv (+ seasons/episodes), people, search, discover, trending,
   genres, keywords, companies, collections, networks, certifications, watch providers, changes,
   reviews, credits — module by module, tracked in `API_COVERAGE.md`.
4. **v3 account coverage**: favorites, watchlist, ratings, lists — using the session auth built
   in phase 2.
5. **v4 phase** (done, see [v4 notes](#v4-notes)): v4 auth, v4-only endpoints, added alongside
   v3 code, not replacing it.

## Success criteria

- Every endpoint in `API_COVERAGE.md` is checked off with a working command, a passing test, and
  (if it returns a non-trivial shape) a `docs/<module>.md` reference entry.
- `make build`, `make test`, `make lint` all pass at every phase boundary.
- A fresh clone + `go install` + a valid `tmdb-sync.toml` can run any implemented command against
  the real TMDB API.

## Roadmap

All phases are complete: every v3 and v4 endpoint in `API_COVERAGE.md` is checked off (the one
exception is `GET /keyword/{id}/movies`, deprecated by TMDB and left unchecked). See
`API_COVERAGE.md` for the endpoint list and `ARCHITECTURE.md` for how new modules get added
without touching existing ones. Remaining work is maintenance: bug fixes and keeping up with
TMDB API changes.

## v4 notes

v4 is 21 endpoints (auth 3, account 9, lists 9 — see `API_COVERAGE.md`'s `## v4 —` sections) on
base `/4/`, complementing v3 rather than replacing it.

**Version selection.** `-v3` / `-v4` flags on the modules that overlap (`account`, `lists`), like
`ping -4/-6`; **v3 is the default**. Giving both is an error. The version is a property of the
invocation, not a global mode (`auth_version` in the config is superseded by the flags). A v4-only
action without `-v4` (e.g. `account -a recommended-movies`) errors with a hint instead of
switching silently, and vice versa. v4 needs `read_access_token` plus the user token from
`auth -v4 -a login`; v4 requests never carry `api_key`.

**Verified against the live API:** the delete-list path is `/4/list/{list_id}` (the reference
prints `/4/{list_id}`, a typo) and clear-list is a `GET`, as documented.
