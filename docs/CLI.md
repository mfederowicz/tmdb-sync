# tmdb-sync — CLI reference

> Describes the intended CLI shape (see `PRD.md`/`ARCHITECTURE.md`). Nothing is implemented yet
> as of this doc — update the examples below as each module actually lands, and check it off in
> `API_COVERAGE.md`.

## Install

```sh
go install github.com/mfederowicz/tmdb-sync@latest
```

## Configuration

TOML file, default path `~/tmdb-sync.toml`, overridable with `-c`:

```toml
# one of these two is required
api_key = "your-v3-api-key"
# read_access_token = "your-v4-read-access-token"   # used as a Bearer header instead of api_key

session_path = "~/.config/tmdb-sync/session.json"   # created by the auth flow, only needed for account commands
per_page = 20
pages_limit = 10
```

## Flags

- `-c <path>` — path to the config file (default `~/tmdb-sync.toml`)
- `-v` — verbose output
- `-version` — print version and exit

## Command shape

```sh
tmdb-sync <module> -a <action> [module-specific flags]
```

`<module>` may be an unambiguous prefix of a module name (mirrors trakt-sync). Running with no
args, or `tmdb-sync help`, lists all modules.

## Planned modules (see API_COVERAGE.md for the full endpoint list per module)

| module          | actions (examples)                          |
|-----------------|----------------------------------------------|
| `configuration` | `details`, `countries`, `languages`, `jobs`, `timezones`, `primary-translations` |
| `movies`        | `details -i <id>`, `popular -p <page>`, `top-rated`, `now-playing`, `upcoming`, `credits -i <id>`, `videos -i <id>`, ... |
| `tv`            | `details -i <id>`, `popular -p <page>`, `top-rated`, `on-the-air`, `airing-today`, ... |
| `tv-seasons`    | `details -i <id> -s <season_number>`, ... |
| `tv-episodes`   | `details -i <id> -s <season_number> -e <episode_number>`, ... |
| `people`        | `details -i <id>`, `popular -p <page>`, `movie-credits -i <id>`, `tv-credits -i <id>`, ... |
| `search`        | `movie -q <query>`, `tv -q <query>`, `person -q <query>`, `multi -q <query>`, `collection -q <query>`, `company -q <query>`, `keyword -q <query>` |
| `discover`      | `movie [filters]`, `tv [filters]` |
| `trending`      | `all -w <day\|week>`, `movie -w <day\|week>`, `tv -w <day\|week>`, `person -w <day\|week>` |
| `genres`        | `movie`, `tv` |
| `keywords`      | `details -i <id>` |
| `companies`     | `details -i <id>`, `images -i <id>` |
| `collections`   | `details -i <id>`, `images -i <id>`, `translations -i <id>` |
| `networks`      | `details -i <id>`, `images -i <id>` |
| `certifications`| `movie`, `tv` |
| `watch-providers`| `regions`, `movie`, `tv` |
| `credits`       | `details -i <credit_id>` |
| `reviews`       | `details -i <review_id>` |
| `changes`       | `movie`, `tv`, `person` |
| `find`          | `-i <external_id> --source <imdb_id\|tvdb_id\|...>` |
| `account`       | 🔒 `details`, `favorites`, `watchlist`, `rated`, `lists`, `add-favorite`, `add-watchlist` |
| `lists`         | `details -i <id>`, `create`, `add-movie`, `remove-movie`, `clear`, `delete` |

🔒 = requires a v3 session; the first 🔒 command run triggers the browser-approval flow
(`cli/session.go`) and persists the session to `session_path`.

## Output

Every command prints one JSON document to stdout (the raw TMDB response shape, or a thin
pagination wrapper) — pipe to `jq` or similar for further processing.
