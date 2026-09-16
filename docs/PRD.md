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

- TMDB **v4** API and v4 auth (request_token → browser approval → v4 access_token, v4 lists,
  v4 account). Explicitly a later phase — see [Roadmap](#roadmap). The v3 phase's architecture
  must not need to change to accommodate it later (see `ARCHITECTURE.md`).
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
5. **v4 phase** (out of scope until v3 is ~100%): v4 auth, v4-only endpoints, added alongside v3
   code, not replacing it.

## Success criteria

- Every endpoint in `API_COVERAGE.md` is checked off with a working command, a passing test, and
  (if it returns a non-trivial shape) a `docs/<module>.md` reference entry.
- `make build`, `make test`, `make lint` all pass at every phase boundary.
- A fresh clone + `go install` + a valid `tmdb-sync.toml` can run any implemented command against
  the real TMDB API.

## Roadmap

See `API_COVERAGE.md` for the concrete, checkable endpoint list, and `ARCHITECTURE.md` for how
new modules get added without touching existing ones.
