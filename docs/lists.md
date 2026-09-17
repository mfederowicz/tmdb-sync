# `lists`

List details, item status. 🚧 Only reads are implemented so far; `create`, `add-movie`,
`remove-movie`, `clear`, and `delete` (all 🔒, requiring a v3 session) are planned in a later PR.

| action        | example                                       | output file                          |
|---------------|------------------------------------------------|---------------------------------------|
| `details`     | `tmdb-sync lists -a details -i 1`               | `lists_details_id-1.json`             |
| `item-status` | `tmdb-sync lists -a item-status -i 1 -media-id 100` | `lists_item-status_id-1_media-100.json` |

Notes:
- `-i <list_id>` is required for all actions.
- `-media-id <movie_id>` is required for `-a item-status`.
- `li` is the module's `Abbrev` — `tmdb-sync li -a details -i 1` also works.
- Both actions are public reads — no session required.
