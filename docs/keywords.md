# `keywords`

Keyword details.

| action    | example                              | output file                    |
|-----------|---------------------------------------|---------------------------------|
| `details` | `tmdb-sync keywords -a details -i 1701` | `keywords_details_id-1701.json` |

Notes:
- `-i <keyword_id>` is required.
- `kw` is the module's `Abbrev` — `tmdb-sync kw -a details -i 1701` also works.
- TMDB's "Movies by keyword" endpoint (`GET /keyword/{keyword_id}/movies`) is deprecated in favor
  of `discover -a movie -with-keywords <id>` and is not implemented here.
