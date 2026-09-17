# `people`

Person details, credits, images, translations, and the popular/latest lists.

| action              | example                                     | output file                         |
|---------------------|----------------------------------------------|---------------------------------------|
| `details`           | `tmdb-sync people -a details -i 31`          | `people_details_id-31.json`           |
| `combined-credits`  | `tmdb-sync people -a combined-credits -i 31` | `people_combined-credits_id-31.json`  |
| `external-ids`      | `tmdb-sync people -a external-ids -i 31`     | `people_external-ids_id-31.json`      |
| `images`            | `tmdb-sync people -a images -i 31`           | `people_images_id-31.json`            |
| `latest`            | `tmdb-sync people -a latest`                 | `people_latest.json`                  |
| `movie-credits`     | `tmdb-sync people -a movie-credits -i 31`    | `people_movie-credits_id-31.json`     |
| `popular`           | `tmdb-sync people -a popular`                | `people_popular.json`                 |
| `tv-credits`        | `tmdb-sync people -a tv-credits -i 31`       | `people_tv-credits_id-31.json`        |
| `translations`      | `tmdb-sync people -a translations -i 31`     | `people_translations_id-31.json`      |

Notes:
- `-i <person_id>` is required for every action except `latest` and `popular`.
- `popular` is paginated like `movies -a popular` — it walks pages up to `-pages-limit`
  (default: `pages_limit` from config, `0` = unlimited).
- `p` is the module's `Abbrev` — `tmdb-sync p -a details -i 31` also works.
