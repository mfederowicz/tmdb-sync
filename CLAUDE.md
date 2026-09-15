# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project status

`tmdb-sync` is a from-scratch Go CLI for the TMDB (themoviedb.org) API, meant to mirror the
architecture of the author's existing [`trakt-sync`](https://github.com/mfederowicz/trakt-sync).

A first code skeleton was written and then **deliberately deleted** to redo documentation first —
see `docs/PRD.md`, `docs/ARCHITECTURE.md`, `docs/API_COVERAGE.md`, and `docs/CLI.md`. Those four
docs are the source of truth for scope, package layout, and the endpoint checklist. Read them
before writing code. Currently the repo contains only `go.mod` and `Makefile` — no Go source.

Key decisions already made (don't relitigate without asking):
- v3 API fully, then v4 later, as an additive phase (see `docs/ARCHITECTURE.md`'s auth section).
- Architecture and CLI shape mirror trakt-sync closely (TOML config, `<module> -a <action>`
  dispatch, one `internal/<x>_service.go` + `handlers/<x>_handler.go` + `str/<x>.go` per endpoint
  area).
- Local git repo only, default branch `main`, no remote yet. Only commit when asked.

## Commands

- `make install` — vendor dependencies (`go mod vendor`)
- `make build` — build the binary, embedding version info via ldflags
- `make test` — `go test -v -race ./...`
- `make cover` — `go test -cover -coverprofile coverage.out ./...`
- `make linter` — `revive --config ./revive.toml --formatter friendly ./...` (requires `revive`;
  no `revive.toml` present yet)
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
