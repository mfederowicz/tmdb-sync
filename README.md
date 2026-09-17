# tmdb-sync

A Go CLI for [The Movie Database (TMDB)](https://www.themoviedb.org) API — modular, TOML-configured,
aiming for full TMDB v3 API coverage, one module at a time.

## Status

🚧 **In progress**, one module at a time. Implemented so far:

- [`certifications`](docs/certifications.md) ✅ — movie/tv certifications, by country.
- [`changes`](docs/changes.md) ✅ — movie/tv/person change lists.
- [`collections`](docs/collections.md) ✅ — collection details, images, translations.
- [`companies`](docs/companies.md) ✅ — company details, alternative names, images.
- [`configuration`](docs/configuration.md) ✅ — image base URLs/sizes, countries, jobs, languages,
  primary translations, timezones.
- [`credits`](docs/credits.md) ✅ — credit details.
- [`discover`](docs/discover.md) ✅ — discover movies/TV shows by filter.
- [`find`](docs/find.md) ✅ — find TMDB movies/TV/people by an external id (IMDb, TVDB, ...).
- `movies` 🚧 — `details`, `popular` implemented; the rest of the module is still planned.

See:

- [`docs/PRD.md`](docs/PRD.md) — goals, non-goals, scope phases, success criteria.
- [`docs/ARCHITECTURE.md`](docs/ARCHITECTURE.md) — package layout and how new modules get added.
- [`docs/API_COVERAGE.md`](docs/API_COVERAGE.md) — checklist of every TMDB v3 endpoint, tracked
  as work lands.
- [`docs/CLI.md`](docs/CLI.md) — full command/config reference.
- [`AGENTS.md`](AGENTS.md) — instructions for coding agents working in this repo.

## Goal

Near-100% coverage of the TMDB **v3** API (movies, tv, people, search, discover, trending,
account, etc.) from a single Go binary configured with one TOML file, with TMDB **v4** auth and
v4-only endpoints planned as a later, additive phase.

## Install & usage

```sh
go install github.com/mfederowicz/tmdb-sync@latest
```

```toml
# ~/tmdb-sync.toml
api_key = "your-v3-api-key"
session_path = "~/.config/tmdb-sync/session.json"
```

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
