# `tv-seasons`

TV season details plus every other TMDB TV season endpoint: aggregate-credits, credits,
external-ids, images, translations, videos, and 🔒 account-states.

| action                   | example                                                              |
|--------------------------|-----------------------------------------------------------------------|
| `details`                | `tmdb-sync tv-seasons -a details -i 1399 -s 1`                         |
| `account-states` 🔒      | `tmdb-sync tv-seasons -a account-states -i 1399 -s 1`                  |
| `aggregate-credits`      | `tmdb-sync tv-seasons -a aggregate-credits -i 1399 -s 1`               |
| `credits`                | `tmdb-sync tv-seasons -a credits -i 1399 -s 1`                         |
| `external-ids`           | `tmdb-sync tv-seasons -a external-ids -i 1399 -s 1`                    |
| `images`                 | `tmdb-sync tv-seasons -a images -i 1399 -s 1`                          |
| `translations`           | `tmdb-sync tv-seasons -a translations -i 1399 -s 1`                    |
| `videos`                 | `tmdb-sync tv-seasons -a videos -i 1399 -s 1`                          |

Notes:
- `-i <series_id>` and `-s <season_number>` are required for every action.
- `-language` applies to `aggregate-credits`, `credits`, `images`, `videos` (TMDB doesn't accept
  it on the other actions).
- `-include-image-language` applies only to `images`.
- `account-states` is 🔒 — it establishes a v3 session on demand the same way the `movies` module
  does, including transparent re-login on a stale session.
- `tv-seasons` has no `Abbrev` shorter than its own name.
