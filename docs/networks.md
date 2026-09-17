# `networks`

Network details, alternative names, and images.

| action              | example                                          | output file                         |
|---------------------|---------------------------------------------------|--------------------------------------|
| `details`           | `tmdb-sync networks -a details -i 1`              | `networks_details_id-1.json`         |
| `alternative-names` | `tmdb-sync networks -a alternative-names -i 1`    | `networks_alternative-names_id-1.json` |
| `images`            | `tmdb-sync networks -a images -i 1`               | `networks_images_id-1.json`          |

Notes:
- `-i <network_id>` is required for all three actions.
- `nw` is the module's `Abbrev` — `tmdb-sync nw -a details -i 1` also works.
