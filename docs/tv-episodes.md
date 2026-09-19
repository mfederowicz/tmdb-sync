# `tv-episodes`

TV episode details plus every other TMDB TV episode endpoint: credits, external-ids, images,
translations, videos, and 🔒 account-states/add-rating/delete-rating.

| action                   | example                                                                     |
|--------------------------|------------------------------------------------------------------------------|
| `details`                | `tmdb-sync tv-episodes -a details -i 1399 -s 1 -e 1`                         |
| `account-states` 🔒      | `tmdb-sync tv-episodes -a account-states -i 1399 -s 1 -e 1`                  |
| `credits`                | `tmdb-sync tv-episodes -a credits -i 1399 -s 1 -e 1`                         |
| `external-ids`           | `tmdb-sync tv-episodes -a external-ids -i 1399 -s 1 -e 1`                    |
| `images`                 | `tmdb-sync tv-episodes -a images -i 1399 -s 1 -e 1`                          |
| `translations`           | `tmdb-sync tv-episodes -a translations -i 1399 -s 1 -e 1`                    |
| `videos`                 | `tmdb-sync tv-episodes -a videos -i 1399 -s 1 -e 1`                          |
| `add-rating` 🔒          | `tmdb-sync tv-episodes -a add-rating -i 1399 -s 1 -e 1 -value 8.5`           |
| `delete-rating` 🔒       | `tmdb-sync tv-episodes -a delete-rating -i 1399 -s 1 -e 1`                   |

Notes:
- `-i <series_id>`, `-s <season_number>`, and `-e <episode_number>` are required for every action.
- `-language` applies to `credits`, `images`, `videos` (TMDB doesn't accept it on the other
  actions).
- `-include-image-language` applies only to `images`.
- `-value <rating>` (0.5-10.0, in 0.5 increments) is required for `add-rating`.
- `-guest` (with `add-rating`) rates as the guest session cached by `guest-sessions -a create`,
  ignoring any account session; it errors if no guest session is cached. `-guest-session-id <id>`
  picks a specific guest session instead.
- `account-states`, `add-rating`, and `delete-rating` are 🔒 — they establish a v3 session on
  demand the same way the `movies` module does, including transparent re-login on a stale session.
- `tv-episodes` has no `Abbrev` shorter than its own name.
