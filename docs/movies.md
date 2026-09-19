# `movies`

Movie details plus every other TMDB movie endpoint: lists, credits, images, videos,
watch providers, and 🔒 account-states/add-rating/delete-rating.

| action              | example                                                     |
|----------------------|--------------------------------------------------------------|
| `details`            | `tmdb-sync movies -a details -i 550`                          |
| `popular`            | `tmdb-sync movies -a popular`                                  |
| `account-states` 🔒  | `tmdb-sync movies -a account-states -i 550`                    |
| `alternative-titles` | `tmdb-sync movies -a alternative-titles -i 550 -country US`    |
| `credits`            | `tmdb-sync movies -a credits -i 550`                           |
| `external-ids`       | `tmdb-sync movies -a external-ids -i 550`                      |
| `images`             | `tmdb-sync movies -a images -i 550`                            |
| `keywords`           | `tmdb-sync movies -a keywords -i 550`                          |
| `latest`             | `tmdb-sync movies -a latest`                                   |
| `lists`              | `tmdb-sync movies -a lists -i 550`                             |
| `now-playing`        | `tmdb-sync movies -a now-playing`                               |
| `recommendations`    | `tmdb-sync movies -a recommendations -i 550`                   |
| `release-dates`      | `tmdb-sync movies -a release-dates -i 550`                     |
| `reviews`            | `tmdb-sync movies -a reviews -i 550`                           |
| `similar`            | `tmdb-sync movies -a similar -i 550`                           |
| `top-rated`          | `tmdb-sync movies -a top-rated`                                 |
| `translations`       | `tmdb-sync movies -a translations -i 550`                      |
| `upcoming`           | `tmdb-sync movies -a upcoming`                                  |
| `videos`             | `tmdb-sync movies -a videos -i 550`                             |
| `watch-providers`    | `tmdb-sync movies -a watch-providers -i 550`                    |
| `add-rating` 🔒      | `tmdb-sync movies -a add-rating -i 550 -value 8.5`             |
| `delete-rating` 🔒   | `tmdb-sync movies -a delete-rating -i 550`                      |

Notes:
- `-i <movie_id>` is required for every action except `popular`, `latest`, `now-playing`,
  `top-rated`, and `upcoming`.
- `-pages-limit` (default: `pages_limit` from config, `0` = unlimited) applies to every paginated
  action: `popular`, `lists`, `now-playing`, `recommendations`, `reviews`, `similar`, `top-rated`,
  `upcoming`.
- `-language` applies to `credits`, `images`, `videos`, and every paginated action above except
  `popular`, `now-playing`, `top-rated`, `upcoming` (TMDB doesn't accept it there).
- `-country` applies only to `alternative-titles`.
- `-include-image-language` applies only to `images`.
- `-value <rating>` (0.5-10.0, in 0.5 increments) is required for `add-rating`.
- `-guest` (with `add-rating`) rates as the guest session cached by `guest-sessions -a create`,
  ignoring any account session; it errors if no guest session is cached. `-guest-session-id <id>`
  picks a specific guest session instead.
- `account-states`, `add-rating`, and `delete-rating` are 🔒 — they establish a v3 session on
  demand the same way the `account` module does, including transparent re-login on a stale
  session. `add-rating`/`delete-rating` use the existing v3-session plumbing directly; they don't
  need the guest-session flow deferred for the future rating/guest-sessions checkpoint.
- `m` is the module's `Abbrev` — `tmdb-sync m -a details -i 550` also works.
