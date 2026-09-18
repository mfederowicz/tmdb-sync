# `watch-providers`

Available watch-provider regions, plus the movie and TV watch provider lists.

| action              | example                                                 | output file                            |
|---------------------|----------------------------------------------------------|-----------------------------------------|
| `available-regions` | `tmdb-sync watch-providers -a available-regions`        | `watch-providers_available-regions.json` |
| `movie-providers`   | `tmdb-sync watch-providers -a movie-providers`           | `watch-providers_movie-providers.json`   |
| `tv-providers`      | `tmdb-sync watch-providers -a tv-providers`              | `watch-providers_tv-providers.json`      |

Notes:
- `-language` (ISO 639-1) is optional for all three actions.
- `-watch-region` (ISO 3166-1) is optional and only used by `movie-providers`/`tv-providers`.
- `wp` is the module's `Abbrev` — `tmdb-sync wp -a available-regions` also works.
