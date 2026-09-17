# `discover`

Discover movies and TV shows by filter (paginated; every action fetches
`min(total_pages, pages_limit)` pages and merges them, per `CLI.md`'s Pagination section).

| action  | example                                                      | output file            |
|---------|-----------------------------------------------------------------|--------------------------|
| `movie` | `tmdb-sync discover -a movie -sort-by popularity.desc -with-genres 28` | `discover_movie.json` |
| `tv`    | `tmdb-sync discover -a tv -year 2020`                        | `discover_tv.json`     |

Filter flags (all optional, apply to both actions except where noted):

- `-sort-by <value>` — e.g. `popularity.desc`
- `-language <ISO 639-1>`
- `-region <ISO 3166-1>` — movie only
- `-include-adult`
- `-year <int>` — `primary_release_year` for movie, `first_air_date_year` for tv
- `-with-genres <comma-separated ids>`
- `-vote-average-gte <float>` / `-vote-average-lte <float>`
- `-with-watch-providers <comma-separated ids>`
- `-watch-region <ISO 3166-1>`
- `-pages-limit <int>` — default: `pages_limit` from config, `0` = unlimited

Notes:
- `d` is the module's `Abbrev` — `tmdb-sync d -a movie ...` also works.
- More TMDB discover filters exist beyond this set; add them to `uri.DiscoverMovieOptions`/
  `uri.DiscoverTVOptions` and the corresponding command flags as needed.
