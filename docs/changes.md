# `changes`

Movie, TV, and person ids changed within a date range.

| action   | example                                                        | output file            |
|----------|-----------------------------------------------------------------|---------------------------|
| `movie`  | `tmdb-sync changes -a movie`                                   | `changes_movie.json`     |
| `tv`     | `tmdb-sync changes -a tv`                                      | `changes_tv.json`        |
| `person` | `tmdb-sync changes -a person`                                  | `changes_person.json`    |

Notes:
- `-start-date`/`-end-date` (`YYYY-MM-DD`) are both optional. TMDB defaults to the last 24 hours
  if omitted, and caps any query to a 14-day range.
- All three actions walk every page up to `-pages-limit` (default: `pages_limit` from config,
  `0` = unlimited), same as `movies -a popular`.
- `ch` is the module's `Abbrev` — `tmdb-sync ch -a movie` also works.

Output file names: every flag you set (other than `-a` and `-pages-limit`) is appended to the name
as `<flag>-<value>`, so runs with different filters write different files. With no extra flags the
name is unchanged. Values with characters outside `a-z 0-9 _ . -` add a short hash suffix.
