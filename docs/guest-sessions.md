# `guest-sessions`

Create a guest session, and read back its rated movies, TV shows, and TV episodes.

| action              | example                                                    | output file                                  |
|---------------------|-------------------------------------------------------------|-------------------------------------------------|
| `create`            | `tmdb-sync guest-sessions -a create`                       | `guest-sessions_create.json`                     |
| `rated-movies`      | `tmdb-sync guest-sessions -a rated-movies -i abc123`        | `guest-sessions_rated-movies_id-abc123.json`     |
| `rated-tv`          | `tmdb-sync guest-sessions -a rated-tv -i abc123`             | `guest-sessions_rated-tv_id-abc123.json`         |
| `rated-tv-episodes` | `tmdb-sync guest-sessions -a rated-tv-episodes -i abc123`    | `guest-sessions_rated-tv-episodes_id-abc123.json` |

Notes:
- `create` takes no flags; it returns a `guest_session_id` (and its expiry) to pass as `-i` to the
  `rated-*` actions below.
- `-i <guest_session_id>` is required for `rated-movies`, `rated-tv`, and `rated-tv-episodes`.
- `-language`, `-sort-by` (`created_at.asc`/`created_at.desc`), and `-pages-limit` are optional for
  the three `rated-*` actions.
- `gs` is the module's `Abbrev` — `tmdb-sync gs -a create` also works.
