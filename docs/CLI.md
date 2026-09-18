# tmdb-sync — CLI reference

> Describes the intended CLI shape (see `PRD.md`/`ARCHITECTURE.md`). `certifications`,
> `configuration`, and `movies` are implemented as of this doc (see
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

`<module>` must be a module's exact name or its abbreviation (see the table below). Running with
no args, or `tmdb-sync help`, lists all modules.

## Modules (see API_COVERAGE.md for the full endpoint list per module)

Ordered to match TMDB's reference nav. ✅ = implemented; the rest are planned.

| module            | actions (examples)                          |
|-------------------|----------------------------------------------|
| `account`         | 🔒 `details`, `favorites`, `watchlist`, `rated`, `lists`, `add-favorite`, `add-watchlist` |
| `certifications` ✅ ([docs](certifications.md)) | `movie`, `tv` |
| `changes` ✅ ([docs](changes.md)) | `movie`, `tv`, `person` |
| `collections` ✅ ([docs](collections.md)) | `details -i <id>`, `images -i <id>`, `translations -i <id>` |
| `companies` ✅ ([docs](companies.md)) | `details -i <id>`, `alternative-names -i <id>`, `images -i <id>` |
| `configuration` ✅ ([docs](configuration.md)) | `details`, `countries`, `languages`, `jobs`, `timezones`, `primary-translations` |
| `credits` ✅ ([docs](credits.md)) | `details -i <credit_id>` |
| `discover` ✅ ([docs](discover.md)) | `movie [filters]`, `tv [filters]` |
| `find` ✅ ([docs](find.md)) | `by-id -i <external_id> --source <imdb_id\|tvdb_id\|...>` |
| `genres` ✅ ([docs](genres.md)) | `movie`, `tv` |
| `guest-sessions` ✅ ([docs](guest-sessions.md)) | `rated-movies -i <guest_session_id>`, `rated-tv -i <guest_session_id>`, `rated-tv-episodes -i <guest_session_id>` |
| `keywords` ✅ ([docs](keywords.md)) | `details -i <id>` |
| `lists` ✅ ([docs](lists.md)) | `details -i <id>`, `item-status -i <id> -media-id <movie_id>`, `create -name <name>` 🔒, `add-movie -i <id> -media-id <id>` 🔒, `remove-movie -i <id> -media-id <id>` 🔒, `clear -i <id>` 🔒, `delete -i <id>` 🔒 |
| `movies` ✅ ([docs](movies.md)) | `details -i <id>`, `popular` (paginated, see below), `top-rated`, `now-playing`, `upcoming`, `credits -i <id>`, `videos -i <id>`, ... full list in `docs/movies.md` |
| `networks` ✅ ([docs](networks.md)) | `details -i <id>`, `alternative-names -i <id>`, `images -i <id>` |
| `people` ✅ ([docs](people.md)) | `details -i <id>`, `popular` (paginated, see below), `latest`, `movie-credits -i <id>`, `tv-credits -i <id>`, ... full list in `docs/people.md` |
| `reviews` ✅ ([docs](reviews.md)) | `details -i <review_id>` |
| `search` ✅ ([docs](search.md)) | `collections -query <query>`, `companies -query <query>`, `keywords -query <query>`, `movies -query <query>`, `multi -query <query>`, `people -query <query>`, `tv -query <query>` |
| `trending` ✅ ([docs](trending.md)) | `all -w <day\|week>`, `movie -w <day\|week>`, `tv -w <day\|week>`, `person -w <day\|week>` |
| `tv` ✅ ([docs](tv.md)) | `details -i <id>`, `popular` (paginated, see below), `top-rated`, `on-the-air`, `airing-today`, `credits -i <id>`, `videos -i <id>`, ... full list in `docs/tv.md` |
| `tv-seasons` ✅ ([docs](tv-seasons.md)) | `details -i <id> -s <season_number>`, `account-states -i <id> -s <season_number>` 🔒, `aggregate-credits -i <id> -s <season_number>`, `credits -i <id> -s <season_number>`, `external-ids -i <id> -s <season_number>`, `images -i <id> -s <season_number>`, `translations -i <id> -s <season_number>`, `videos -i <id> -s <season_number>` |
| `tv-episodes` ✅ ([docs](tv-episodes.md)) | `details -i <id> -s <season_number> -e <episode_number>`, `account-states -i <id> -s <season_number> -e <episode_number>` 🔒, `credits -i <id> -s <season_number> -e <episode_number>`, `external-ids -i <id> -s <season_number> -e <episode_number>`, `images -i <id> -s <season_number> -e <episode_number>`, `translations -i <id> -s <season_number> -e <episode_number>`, `videos -i <id> -s <season_number> -e <episode_number>`, `add-rating -i <id> -s <season_number> -e <episode_number> -value <rating>` 🔒, `delete-rating -i <id> -s <season_number> -e <episode_number>` 🔒 |
| `tv-episode-groups` ✅ ([docs](tv-episode-groups.md)) | `details -i <episode_group_id>` |
| `watch-providers` ✅ ([docs](watch-providers.md)) | `available-regions`, `movie-providers`, `tv-providers` |

🔒 = requires a v3 session. There's no standalone `authentication` command — a 🔒 module (like
`account`) triggers the request-token → browser-approval → session flow itself, on demand
(`cli.HandleToken`/`cli.CreateSessionInteractively` in `cli/session.go`), and persists the result
to `session_path` for subsequent runs to reuse.

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
