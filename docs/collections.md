# `collections`

Collection details, images, and translations.

| action         | example                                              | output file                |
|----------------|-------------------------------------------------------|----------------------------|
| `details`      | `tmdb-sync collections -a details -i 10`             | `collections_details.json`      |
| `images`       | `tmdb-sync collections -a images -i 10`              | `collections_images.json`       |
| `translations` | `tmdb-sync collections -a translations -i 10`        | `collections_translations.json` |

Notes:
- `-i <collection_id>` is required for all three actions.
- `-language`/`-include-image-language` are optional, used only by `-a images`.
- `co` is the module's `Abbrev` — `tmdb-sync co -a details -i 10` also works.
