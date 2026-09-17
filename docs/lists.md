# `lists`

List details, item status, and (🔒, requiring a v3 session) create/mutate.

| action         | example                                              | output file                             |
|----------------|-------------------------------------------------------|-------------------------------------------|
| `details`      | `tmdb-sync lists -a details -i 1`                     | `lists_details_id-1.json`                 |
| `item-status`  | `tmdb-sync lists -a item-status -i 1 -media-id 100`   | `lists_item-status_id-1_media-100.json`   |
| `create` 🔒    | `tmdb-sync lists -a create -name "watch later"`       | `lists_create_id-<list_id>.json`          |
| `add-movie` 🔒 | `tmdb-sync lists -a add-movie -i 1 -media-id 100`     | `lists_add-movie_id-1_media-100.json`     |
| `remove-movie` 🔒 | `tmdb-sync lists -a remove-movie -i 1 -media-id 100` | `lists_remove-movie_id-1_media-100.json` |
| `clear` 🔒     | `tmdb-sync lists -a clear -i 1`                       | `lists_clear_id-1.json`                   |
| `delete` 🔒    | `tmdb-sync lists -a delete -i 1`                      | `lists_delete_id-1.json`                  |

Notes:
- `-i <list_id>` is required for all actions except `create`.
- `-media-id <movie_id>` is required for `item-status`, `add-movie`, `remove-movie`.
- `-name <name>` is required for `create`; `-description` and `-language` are optional.
- `details` and `item-status` are public reads — no session required.
- `create`, `add-movie`, `remove-movie`, `clear`, `delete` are 🔒 — they establish a v3 session on
  demand the same way the `account` module does, including transparent re-login on a stale session.
- `clear` always passes TMDB's required `confirm=true` — there's no separate confirmation flag.
- `li` is the module's `Abbrev` — `tmdb-sync li -a details -i 1` also works.
