# `guest-sessions`

A guest session's rated movies, TV shows, and TV episodes.

| action              | example                                                    | output file                                  |
|---------------------|-------------------------------------------------------------|-------------------------------------------------|
| `rated-movies`      | `tmdb-sync guest-sessions -a rated-movies -i abc123`        | `guest-sessions_rated-movies_id-abc123.json`     |
| `rated-tv`          | `tmdb-sync guest-sessions -a rated-tv -i abc123`             | `guest-sessions_rated-tv_id-abc123.json`         |
| `rated-tv-episodes` | `tmdb-sync guest-sessions -a rated-tv-episodes -i abc123`    | `guest-sessions_rated-tv-episodes_id-abc123.json` |

Notes:
- `-i <guest_session_id>` is required for all three actions.
- `-language`, `-sort-by` (`created_at.asc`/`created_at.desc`), and `-pages-limit` are optional for
  all three actions.
- `gs` is the module's `Abbrev` — `tmdb-sync gs -a rated-movies -i abc123` also works.
- There's no CLI action to create a guest session yet — see `docs/API_COVERAGE.md`'s Authentication
  section for why `CreateGuestSession` isn't wired in.
