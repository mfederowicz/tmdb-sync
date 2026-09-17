# `reviews`

Review details.

| action    | example                                | output file                  |
|-----------|------------------------------------------|-------------------------------|
| `details` | `tmdb-sync reviews -a details -i abc123` | `reviews_details_id-abc123.json` |

Notes:
- `-i <review_id>` is required.
- `rv` is the module's `Abbrev` — `tmdb-sync rv -a details -i abc123` also works.
