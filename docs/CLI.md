# tmdb-sync — CLI reference

> Describes the intended CLI shape (see `PRD.md`/`ARCHITECTURE.md`). `authentication`,
> `certifications`, `configuration`, and `movies` are implemented as of this doc (see
> `API_COVERAGE.md`, marked ✅ below); the rest of the module table is still planned — update
> examples as each module actually lands. The table is ordered to match TMDB's own reference nav
> (developer.themoviedb.org/reference) category-for-category, so the CLI's module list maps 1:1
> onto TMDB's docs.

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
output_dir = ""        # where result JSON files are written; "" = current directory
per_page = 20
pages_limit = 10       # cap on pages walked by -all; 0 = unlimited (bounded by TMDB's total_pages)
```

## Flags

- `-c <path>` — path to the config file (default `~/tmdb-sync.toml`)
- `-v` — verbose output
- `-version` — print version and exit

## Command shape

```sh
tmdb-sync <module> -a <action> [module-specific flags]
```

`<module>` may be an unambiguous prefix of a module name. Running with no
args, or `tmdb-sync help`, lists all modules.

## Modules (see API_COVERAGE.md for the full endpoint list per module)

Ordered to match TMDB's reference nav. ✅ = implemented; the rest are planned.

| module            | actions (examples)                          |
|-------------------|----------------------------------------------|
| `account`         | 🔒 `details`, `favorites`, `watchlist`, `rated`, `lists`, `add-favorite`, `add-watchlist` |
| `authentication` ✅| `validate-key`, `create-request-token`, `create-session` (🔒 interactive browser approval), `create-guest-session`, `delete-session -s <session_id>` |
| `certifications` ✅| `movie`, `tv` |
| `changes`         | `movie`, `tv`, `person` |
| `collections`     | `details -i <id>`, `images -i <id>`, `translations -i <id>` |
| `companies`       | `details -i <id>`, `images -i <id>` |
| `configuration` ✅ | `details`, `countries`, `languages`, `jobs`, `timezones`, `primary-translations` |
| `credits`         | `details -i <credit_id>` |
| `discover`        | `movie [filters]`, `tv [filters]` |
| `find`            | `-i <external_id> --source <imdb_id\|tvdb_id\|...>` |
| `genres`          | `movie`, `tv` |
| `keywords`        | `details -i <id>` |
| `lists`           | `details -i <id>`, `create`, `add-movie`, `remove-movie`, `clear`, `delete` |
| `movies` ✅        | `details -i <id>`, `popular` (paginated, see below), `top-rated`, `now-playing`, `upcoming`, `credits -i <id>`, `videos -i <id>`, ... |
| `networks`        | `details -i <id>`, `images -i <id>` |
| `people`          | `details -i <id>`, `popular`, `movie-credits -i <id>`, `tv-credits -i <id>`, ... |
| `reviews`         | `details -i <review_id>` |
| `search`          | `movie -q <query>`, `tv -q <query>`, `person -q <query>`, `multi -q <query>`, `collection -q <query>`, `company -q <query>`, `keyword -q <query>` |
| `trending`        | `all -w <day\|week>`, `movie -w <day\|week>`, `tv -w <day\|week>`, `person -w <day\|week>` |
| `tv`              | `details -i <id>`, `popular`, `top-rated`, `on-the-air`, `airing-today`, ... |
| `tv-seasons`      | `details -i <id> -s <season_number>`, ... |
| `tv-episodes`     | `details -i <id> -s <season_number> -e <episode_number>`, ... |
| `watch-providers` | `regions`, `movie`, `tv` |

🔒 = requires a v3 session. `authentication -a create-session` runs the browser-approval flow
(`cli/session.go`) and persists the session to `session_path`; other 🔒 modules (like `account`)
read that persisted session rather than triggering the flow themselves.

## Pagination

TMDB list endpoints paginate via `page`/`total_pages` fields in the JSON body (there's no
pagination info in HTTP headers, unlike some other APIs), and TMDB fixes the page size at 20 items —
there's no API lever to change it. Every list action (e.g. `movies -a popular`) always fetches
`min(total_pages, pages_limit)` pages and merges them into one result — no separate flag needed:

```sh
tmdb-sync movies -a popular            # up to pages_limit pages (from config), merged
```

Set `pages_limit = 0` in config to fetch every page TMDB has for that list.

## Output

Every command writes its result as one JSON file, named `<module>_<action>[_<params>].json` (e.g.
`movies_details_id-550.json`, `movies_popular.json`), under `output_dir` (default: the current
directory), and prints a one-line `wrote <path>` confirmation — the file is the product, not the
terminal output. Read/process the file with `jq`, a script, etc.
