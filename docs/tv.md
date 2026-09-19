# `tv`

TV series details plus every other TMDB TV series endpoint: lists, credits, images, videos,
watch providers, and 🔒 account-states/add-rating/delete-rating.

| action                   | example                                                        |
|--------------------------|-----------------------------------------------------------------|
| `details`                | `tmdb-sync tv -a details -i 1399`                                |
| `popular`                | `tmdb-sync tv -a popular`                                        |
| `account-states` 🔒      | `tmdb-sync tv -a account-states -i 1399`                         |
| `aggregate-credits`      | `tmdb-sync tv -a aggregate-credits -i 1399`                      |
| `airing-today`           | `tmdb-sync tv -a airing-today`                                   |
| `alternative-titles`     | `tmdb-sync tv -a alternative-titles -i 1399`                     |
| `content-ratings`        | `tmdb-sync tv -a content-ratings -i 1399`                        |
| `credits`                | `tmdb-sync tv -a credits -i 1399`                                |
| `episode-groups`         | `tmdb-sync tv -a episode-groups -i 1399`                         |
| `external-ids`           | `tmdb-sync tv -a external-ids -i 1399`                           |
| `images`                 | `tmdb-sync tv -a images -i 1399`                                 |
| `keywords`               | `tmdb-sync tv -a keywords -i 1399`                               |
| `latest`                 | `tmdb-sync tv -a latest`                                         |
| `lists`                  | `tmdb-sync tv -a lists -i 1399`                                  |
| `on-the-air`             | `tmdb-sync tv -a on-the-air`                                     |
| `recommendations`        | `tmdb-sync tv -a recommendations -i 1399`                        |
| `reviews`                | `tmdb-sync tv -a reviews -i 1399`                                |
| `screened-theatrically`  | `tmdb-sync tv -a screened-theatrically -i 1399`                  |
| `similar`                | `tmdb-sync tv -a similar -i 1399`                                |
| `top-rated`              | `tmdb-sync tv -a top-rated`                                      |
| `translations`           | `tmdb-sync tv -a translations -i 1399`                           |
| `videos`                 | `tmdb-sync tv -a videos -i 1399`                                 |
| `watch-providers`        | `tmdb-sync tv -a watch-providers -i 1399`                        |
| `add-rating` 🔒          | `tmdb-sync tv -a add-rating -i 1399 -value 8.5`                  |
| `delete-rating` 🔒       | `tmdb-sync tv -a delete-rating -i 1399`                          |

Notes:
- `-i <series_id>` is required for every action except `airing-today`, `latest`, `on-the-air`,
  `popular`, and `top-rated`.
- `-pages-limit` (default: `pages_limit` from config, `0` = unlimited) applies to every paginated
  action: `airing-today`, `lists`, `on-the-air`, `popular`, `recommendations`, `reviews`,
  `similar`, `top-rated`.
- `-language` applies to `aggregate-credits`, `credits`, `images`, `lists`, `recommendations`,
  `reviews`, `similar`, `videos` (TMDB doesn't accept it on the other paginated/latest actions).
- `-include-image-language` applies only to `images`.
- `-value <rating>` (0.5-10.0, in 0.5 increments) is required for `add-rating`.
- `-guest` (with `add-rating`) rates as the guest session cached by `guest-sessions -a create`,
  ignoring any account session; it errors if no guest session is cached. `-guest-session-id <id>`
  picks a specific guest session instead.
- `account-states`, `add-rating`, and `delete-rating` are 🔒 — they establish a v3 session on
  demand the same way the `movies` module does, including transparent re-login on a stale session.
- `tv` has no `Abbrev` shorter than its own name.
