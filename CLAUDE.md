# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project status

`tmdb-sync` is a Go CLI for the TMDB (themoviedb.org) API. `docs/PRD.md`, `docs/ARCHITECTURE.md`,
`docs/API_COVERAGE.md`, and `docs/CLI.md` are the source of truth for scope, package layout, and
the endpoint checklist — read them before writing code.

Key decisions already made (don't relitigate without asking):
- v3 API fully, then v4 later, as an additive phase (see `docs/ARCHITECTURE.md`'s auth section).
- TOML config, `<module> -a <action>` command dispatch, one `internal/<x>_service.go` +
  `handlers/<x>_handler.go` + `str/<x>.go` per endpoint area (see `docs/ARCHITECTURE.md`).
- One module per branch/PR. Before starting a module with more than a couple endpoints, ask the
  user how many endpoints to tackle in this session/branch rather than assuming a fixed count —
  available token budget varies per session, so split into smaller branches when the user says so.
- Modules are implemented in TMDB's own reference-nav order (see `docs/API_COVERAGE.md`'s section
  order and `docs/CLI.md`'s module table).
- `main` is pushed to a public GitHub remote now — commit and push only when asked.

## Commands

- `make install` — vendor dependencies (`go mod vendor`)
- `make build` — build the binary, embedding version info via ldflags
- `make test` — `go test -v -race ./...`
- `make cover` — `go test -cover -coverprofile coverage.out ./...`
- `make linter` — `revive --config ./revive.toml --formatter friendly ./...` (requires `revive`)
- `make cleanup` — `gofmt -w` on all `.go` files
- `make clean` — removes `*.json` files in the repo root

Module path: `github.com/mfederowicz/tmdb-sync`, Go 1.21.

## Working on this repo

1. Check `docs/API_COVERAGE.md` for what's implemented vs. not, and pick up the next unchecked
   item (or whatever the user names).
2. Follow the "Adding a new module" recipe in `docs/ARCHITECTURE.md` — it's the same four-file
   shape for every endpoint, plus two registries to update (`Client` field list, `cmds.Commands`).
3. Tick off `docs/API_COVERAGE.md` and add a `docs/<module>.md` reference entry (per `docs/CLI.md`'s
   table) once a module actually works and is tested.
