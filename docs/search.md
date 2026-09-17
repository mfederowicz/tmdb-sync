# `search`

Search TMDB for collections, companies, keywords, movies, people, and TV shows.

| action       | example                                       | output file                       |
|--------------|------------------------------------------------|------------------------------------|
| `collections`| `tmdb-sync search -a collections -query harry` | `search_collections_query-harry.json` |
| `companies`  | `tmdb-sync search -a companies -query pixar`   | `search_companies_query-pixar.json`   |
| `keywords`   | `tmdb-sync search -a keywords -query space`    | `search_keywords_query-space.json`    |
| `movies`     | `tmdb-sync search -a movies -query matrix`     | `search_movies_query-matrix.json`     |
| `multi`      | `tmdb-sync search -a multi -query matrix`      | `search_multi_query-matrix.json`      |
| `people`     | `tmdb-sync search -a people -query keanu`      | `search_people_query-keanu.json`      |
| `tv`         | `tmdb-sync search -a tv -query "breaking bad"` | `search_tv_query-breaking bad.json`   |

Notes:
- `-query <query>` is required for all actions.
- `-language`, `-include-adult` apply to `collections`, `movies`, `multi`, `people`, and `tv`.
- `-region` applies to `collections` and `movies`.
- `-year`, `-primary-release-year` apply to `movies`; `-year`, `-first-air-date-year` apply to `tv`.
- Results are paginated; `-pages-limit` caps how many pages are walked (default: `pages_limit`
  from config, `0` = unlimited), same as other list actions.
- `sr` is the module's `Abbrev` — `tmdb-sync sr -a movies -query matrix` also works.
- `multi` results are movie/tv/person shapes merged into one struct, distinguished by
  `media_type`; fields not relevant to a given result's media type are left zero.
