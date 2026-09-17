# `companies`

Company details, alternative names, and images.

| action              | example                                            | output file                     |
|---------------------|-----------------------------------------------------|----------------------------------|
| `details`           | `tmdb-sync companies -a details -i 1`              | `companies_details.json`             |
| `alternative-names` | `tmdb-sync companies -a alternative-names -i 1`    | `companies_alternative-names.json`   |
| `images`            | `tmdb-sync companies -a images -i 1`               | `companies_images.json`              |

Notes:
- `-i <company_id>` is required for all three actions.
- `cp` is the module's `Abbrev` — `tmdb-sync cp -a details -i 1` also works.
