# AGENTS.md

Instructions for any coding agent (Claude Code, Codex, Copilot, etc.) working in this
repository. Tool-specific files (`CLAUDE.md`, ...) may add detail but must not contradict this
file; this file is the generic source of truth.

## What this project is

`tmdb-sync` is a Go CLI for The Movie Database (TMDB) API: TOML config, `<module> -a <action>`
command dispatch, one `internal/<x>_service.go` + `handlers/<x>_handler.go` + `str/<x>.go` per
API endpoint area.

## Current state

Docs are the source of truth for scope, package layout, and the endpoint checklist. Read these
before writing any code, in this order:

1. `docs/PRD.md` — scope, goals, non-goals, phases.
2. `docs/ARCHITECTURE.md` — package layout and the exact recipe for adding a module.
3. `docs/API_COVERAGE.md` — the checklist of every TMDB v3 endpoint; work off this, don't
   invent scope.
4. `docs/CLI.md` — user-facing command/config shape.

## Decisions already made — don't relitigate without asking

- TMDB **v3** API first, aiming for near-100% coverage, before touching v4 at all.
- v4 auth/endpoints are an explicitly separate, later phase and must be addable as new files
  (`internal/auth_v4_service.go`, a new `cli` flow) without modifying v3 code
  (`internal/auth_service.go`, `cli/session.go`).
- Don't redesign the established shape (e.g. don't switch to cobra/urfave-cli, don't rename the
  config format away from TOML) without asking first.
- One module per branch/PR; if a module has more than ~3 endpoints, split it into smaller
  branches instead of landing it all at once.
- Modules are implemented in TMDB's own reference-nav order (see `docs/API_COVERAGE.md`'s section
  order and `docs/CLI.md`'s module table).
- Default branch `main`, pushed to a public GitHub remote. Only create commits/pushes when
  explicitly asked to.

## Working on this repo

- Pick the next unchecked item in `docs/API_COVERAGE.md` (or whatever module the user names),
  follow the "Adding a new module" recipe in `docs/ARCHITECTURE.md` — it's the same four-file
  shape every time, plus two registries to update (the `Client` field list in
  `internal/client.go`, and `cmds.Commands` in `cmds/runtime.go`'s `init()`).
- Every new `internal/<x>_service.go` method needs an `httptest`-backed test alongside it —
  see `docs/ARCHITECTURE.md`'s testing-strategy section.
- Tick off `docs/API_COVERAGE.md` and add a `docs/<module>.md` reference entry once a module
  actually works end-to-end and is tested — don't mark things done speculatively.
- Run `make build`, `make test` (or `go build ./...`, `go test -race ./...`) before considering
  any change finished.
