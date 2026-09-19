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

## v4 (`-v4`)

| action    | example                                                  | output file                |
|-----------|-----------------------------------------------------------|----------------------------|
| `details` | `tmdb-sync lists -v4 -a details -i 8 -language en-US`     | `lists_details_id-8_v4.json` |
| `create` 🔒 | `tmdb-sync lists -v4 -a create -name "x" -language en -country US -public` | `lists_create_id-<id>_v4.json` |
| `update` 🔒 | `tmdb-sync lists -v4 -a update -i 8 -name "new" -public=false` | `lists_update_id-8_v4.json` |
| `delete` 🔒 | `tmdb-sync lists -v4 -a delete -i 8` | `lists_delete_id-8_v4.json` |
| `add-items` 🔒 | `tmdb-sync lists -v4 -a add-items -i 8 -item movie:100 -item tv:200` | `lists_add-items_id-8_v4.json` |
| `update-items` 🔒 | `tmdb-sync lists -v4 -a update-items -i 8 -item "movie:100:my comment"` | `lists_update-items_id-8_v4.json` |
| `remove-items` 🔒 | `tmdb-sync lists -v4 -a remove-items -i 8 -item movie:100 -item tv:200` | `lists_remove-items_id-8_v4.json` |
| `item-status` | `tmdb-sync lists -v4 -a item-status -i 8 -media-type tv -media-id 100` | `lists_item-status_id-8_media-100_v4.json` |
| `clear` 🔒 | `tmdb-sync lists -v4 -a clear -i 8` | `lists_clear_id-8_v4.json` |

- `-v4 -a details` walks every item page (`-pages-limit`, default from config, 0 = unlimited) and
  merges the items into `results`. `-language` and `-sort-by` (e.g. `original_order.asc`) are optional.
- Public lists need only `read_access_token`; if `auth -v4 -a login` was run, the user token is sent
  so private lists you own are readable too.
- `-v4 -a create` needs `auth -v4 -a login` first; `-name`, `-language` (ISO 639-1) and `-country`
  (ISO 3166-1) are required, `-description` and `-public` (default private) are optional.
- `-v4 -a update` needs `auth -v4 -a login` and at least one of `-name`, `-description`, `-public`
  (`-public=false` makes it private), `-sort-by`, `-backdrop-path`; only the flags you pass are sent.
- `-v4 -a delete` needs `auth -v4 -a login`. It calls `DELETE /4/list/{list_id}`; TMDB's reference
  documents the path as `/4/{list_id}`, which looks like a typo (to be confirmed against the live API).
- `-v4 -a add-items` needs `auth -v4 -a login` and one or more `-item movie:<id>` / `-item tv:<id>`;
  the output has TMDB's per-item `results`.
- `-v4 -a update-items` sets the comment of items already on the list: `-item movie:<id>:<comment>`
  (repeatable, comment required, may contain colons). `add-items` rejects comments.
- `-v4 -a remove-items` needs `auth -v4 -a login` and one or more `-item movie:<id>` / `-item tv:<id>`;
  the output has TMDB's per-item `results`. The ids are sent as a JSON body on the DELETE request.
- `-v4 -a item-status` takes `-media-type movie|tv` and `-media-id`; like `details` it works on public
  lists with only `read_access_token` and uses the user token when cached.
- `-v4 -a clear` needs `auth -v4 -a login`. It sends `GET /4/list/{list_id}/clear` as TMDB's reference
  documents it (a destructive GET; to be confirmed against the live API) and removes every item.
