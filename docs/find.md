# `find`

Find TMDB movies/TV/people by an external id (IMDb, TVDB, ...).

| action  | example                                                  | output file       |
|---------|-----------------------------------------------------------|-------------------|
| `by-id` | `tmdb-sync find -a by-id -i tt0137523 --source imdb_id` | `find_by-id_id-tt0137523.json` |

Flags:

- `-i <external_id>` — required
- `--source <external_source>` — required, e.g. `imdb_id`, `tvdb_id`, `facebook_id`, `tiktok_id`
- `--language <ISO 639-1>` — optional

Response includes `movie_results`, `tv_results`, `person_results`, `tv_episode_results`, and
`tv_season_results` — TMDB returns whichever of these the external id actually matches.

Notes:
- `f` is the module's `Abbrev` — `tmdb-sync f -a by-id ...` also works.
