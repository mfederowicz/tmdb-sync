# `configuration`

TMDB's own reference data: image CDN base URLs/sizes, and the lookup lists (countries, jobs,
languages, primary translations, timezones) TMDB uses across its other endpoints.

| action                  | example                                              | output file                              |
|--------------------------|--------------------------------------------------------|--------------------------------------------|
| `details`                | `tmdb-sync configuration -a details`                  | `configuration_details.json`               |
| `countries`               | `tmdb-sync configuration -a countries`                | `configuration_countries.json`             |
| `jobs`                    | `tmdb-sync configuration -a jobs`                     | `configuration_jobs.json`                  |
| `languages`               | `tmdb-sync configuration -a languages`                | `configuration_languages.json`             |
| `primary-translations`    | `tmdb-sync configuration -a primary-translations`     | `configuration_primary-translations.json`  |
| `timezones`               | `tmdb-sync configuration -a timezones`                | `configuration_timezones.json`             |

`config` is the module's `Abbrev` — `tmdb-sync config -a details` also works.
