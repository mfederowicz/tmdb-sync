# `trending`

List trending movies, TV shows, and people.

| action  | example                              | output file                     |
|---------|--------------------------------------|----------------------------------|
| `all`   | `tmdb-sync trending -a all -w day`   | `trending_all_window-day.json`   |
| `movie` | `tmdb-sync trending -a movie -w week`| `trending_movie_window-week.json`|
| `tv`    | `tmdb-sync trending -a tv -w day`    | `trending_tv_window-day.json`    |
| `person`| `tmdb-sync trending -a person -w week`| `trending_person_window-week.json`|

Notes:
- `-w <day|week>` selects the time window; defaults to `day`.
- Results are paginated; `-pages-limit` caps how many pages are walked (default: `pages_limit`
  from config, `0` = unlimited), same as other list actions.
- `tr` is the module's `Abbrev` — `tmdb-sync tr -a movie -w day` also works.
- `all` results are movie/tv/person shapes merged into one struct, distinguished by
  `media_type`; fields not relevant to a given result's media type are left zero.
