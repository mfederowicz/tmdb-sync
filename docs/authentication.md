# `authentication`

Validate an API key, run the v3 request-token/session flow, and manage guest sessions.
`create-session` and `delete-session` are 🔒 — they read/write `session_path` from your config.

| action                 | example                                                | output file                          |
|------------------------|---------------------------------------------------------|---------------------------------------|
| `validate-key`         | `tmdb-sync authentication -a validate-key`             | `authentication_validate-key.json`     |
| `create-request-token` | `tmdb-sync authentication -a create-request-token`     | `authentication_create-request-token.json` |
| `create-session`       | `tmdb-sync authentication -a create-session`           | `authentication_create-session.json`   |
| `create-guest-session`  | `tmdb-sync authentication -a create-guest-session`     | `authentication_create-guest-session.json` |
| `delete-session`        | `tmdb-sync authentication -a delete-session -s <session_id>` | `authentication_delete-session.json` |

`create-session` opens a browser for TMDB's approval step (`cli.CreateSessionInteractively`) and
persists the resulting session to `session_path`. `delete-session` defaults `-s` to that persisted
session if omitted.

`auth` is the module's `Abbrev` — `tmdb-sync auth -a validate-key` also works.
