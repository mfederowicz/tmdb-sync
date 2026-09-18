# `guest-sessions`

Create a guest session, and read back its rated movies, TV shows, and TV episodes.

| action              | example                                                    | output file                                  |
|---------------------|-------------------------------------------------------------|-------------------------------------------------|
| `create`            | `tmdb-sync guest-sessions -a create`                       | `guest-sessions_create.json`                     |
| `rated-movies`      | `tmdb-sync guest-sessions -a rated-movies`                  | `guest-sessions_rated-movies_id-<id>.json`       |
| `rated-tv`          | `tmdb-sync guest-sessions -a rated-tv`                       | `guest-sessions_rated-tv_id-<id>.json`           |
| `rated-tv-episodes` | `tmdb-sync guest-sessions -a rated-tv-episodes`              | `guest-sessions_rated-tv-episodes_id-<id>.json`  |

Notes:
- `create` takes no flags; it caches the returned `guest_session_id` to
  `~/.config/tmdb-sync/guest_session.json` (path configurable via `guest_session_path` in the TOML
  config), the same way `account -a details` caches `account.json`.
- `-i <guest_session_id>` is optional for `rated-movies`, `rated-tv`, and `rated-tv-episodes` — omit
  it to reuse the id cached by the last `-a create`, or pass it explicitly to use a different guest
  session.
- `-language`, `-sort-by` (`created_at.asc`/`created_at.desc`), and `-pages-limit` are optional for
  the three `rated-*` actions.
- `gs` is the module's `Abbrev` — `tmdb-sync gs -a create` also works.
