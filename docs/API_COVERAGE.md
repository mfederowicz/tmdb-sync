# TMDB v3 API Coverage Checklist

Tracks progress toward "near-100% v3 coverage" (`PRD.md`). Check an item off only once it has:
a `cmds`/`handlers` entry, an `internal` service method with an `httptest`-backed test, and (for
non-trivial response shapes) a `docs/<module>.md` reference entry.

Endpoint names/paths are from the TMDB v3 API reference (developer.themoviedb.org) at time of
writing — re-verify against the live reference before implementing each one, in case TMDB has
added/deprecated endpoints since.

Account endpoints (marked 🔒) require a v3 session (`session_id`) — implemented in phase 4,
after the session-auth plumbing from phase 2 exists but before any account command is wired up.

## Account 🔒
- [x] Details — `GET /account/{account_id}`
- [x] Add to Watchlist — `POST /account/{account_id}/watchlist`
- [x] Add/Remove Favorite — `POST /account/{account_id}/favorite`
- [x] Get Favorite Movies — `GET /account/{account_id}/favorite/movies`
- [x] Get Favorite TV Shows — `GET /account/{account_id}/favorite/tv`
- [x] Get Lists — `GET /account/{account_id}/lists`
- [x] Get Rated Movies — `GET /account/{account_id}/rated/movies`
- [x] Get Rated TV Shows — `GET /account/{account_id}/rated/tv`
- [x] Get Rated TV Episodes — `GET /account/{account_id}/rated/tv/episodes`
- [x] Get Watchlist Movies — `GET /account/{account_id}/watchlist/movies`
- [x] Get Watchlist TV Shows — `GET /account/{account_id}/watchlist/tv`

## Authentication
Used internally by the session flow (`internal/auth_service.go`, `cli/session.go`), not exposed as
a CLI action. `CreateRequestToken`/`CreateSession` back the `account` module's login flow
(`cli.HandleToken`); `ValidateKey` backs a startup credentials preflight (`main.go`), and a 401
from any `account` call now transparently re-triggers login via `cli.CreateSessionInteractively`
(`cmds/command_account.go`'s `execAccountAttempt`). `DeleteSession` and `CreateGuestSession` are
implemented and unit-tested but **not usable/wired-in yet** — marked unchecked deliberately (not
an oversight). They're blocked on a future checkpoint: a `rating` module (movie/TV/episode rating
+ reading back guest ratings) that would give `CreateGuestSession` a real caller; `DeleteSession`
has no organic caller identified at all (TMDB has no equivalent delete for guest sessions — they
just expire after 60 min idle — and there's no planned explicit account-logout feature either).
Revisit both once that checkpoint is reached.
- [ ] Create Guest Session — `GET /authentication/guest_session/new`
- [x] Create Request Token — `GET /authentication/token/new`
- [x] Create Session — `POST /authentication/session/new`
- [ ] Delete Session (logout) — `DELETE /authentication/session`
- [x] Validate Key — `GET /authentication`

## Certifications
- [x] Movie Certifications — `GET /certification/movie/list`
- [x] TV Certifications — `GET /certification/tv/list`

## Changes
- [x] Movie Change List — `GET /movie/changes`
- [x] People Change List — `GET /person/changes`
- [x] TV Change List — `GET /tv/changes`

## Collections
- [x] Details — `GET /collection/{collection_id}`
- [x] Images — `GET /collection/{collection_id}/images`
- [x] Translations — `GET /collection/{collection_id}/translations`

## Companies
- [x] Details — `GET /company/{company_id}`
- [x] Alternative Names — `GET /company/{company_id}/alternative_names`
- [x] Images — `GET /company/{company_id}/images`

## Configuration
- [x] Details — `GET /configuration`
- [x] Countries — `GET /configuration/countries`
- [x] Jobs — `GET /configuration/jobs`
- [x] Languages — `GET /configuration/languages`
- [x] Primary Translations — `GET /configuration/primary_translations`
- [x] Timezones — `GET /configuration/timezones`

## Credits
- [x] Details — `GET /credit/{credit_id}`

## Discover
- [x] Movie — `GET /discover/movie`
- [x] TV — `GET /discover/tv`

## Find
- [x] By external ID — `GET /find/{external_id}`

## Genres
- [x] Movie List — `GET /genre/movie/list`
- [x] TV List — `GET /genre/tv/list`

## Guest Sessions
- [ ] Rated Movies — `GET /guest_session/{guest_session_id}/rated/movies`
- [ ] Rated TV Shows — `GET /guest_session/{guest_session_id}/rated/tv`
- [ ] Rated TV Episodes — `GET /guest_session/{guest_session_id}/rated/tv/episodes`

## Keywords
- [x] Details — `GET /keyword/{keyword_id}`
- [ ] Movies by keyword — `GET /keyword/{keyword_id}/movies` *(deprecated by TMDB — "Use
      `/discover/movie` with `with_keywords` instead"; not implemented)*

## Lists (v3) 🔒 for mutation, public for read
- [x] Details — `GET /list/{list_id}`
- [x] Check Item Status — `GET /list/{list_id}/item_status`
- [x] Create — `POST /list`
- [x] Add Movie — `POST /list/{list_id}/add_item`
- [x] Remove Movie — `POST /list/{list_id}/remove_item`
- [x] Clear — `POST /list/{list_id}/clear`
- [x] Delete — `DELETE /list/{list_id}`

## Movies
- [x] Details — `GET /movie/{movie_id}`
- [x] Account States 🔒 — `GET /movie/{movie_id}/account_states`
- [x] Alternative Titles — `GET /movie/{movie_id}/alternative_titles`
- [x] Credits — `GET /movie/{movie_id}/credits`
- [x] External IDs — `GET /movie/{movie_id}/external_ids`
- [x] Images — `GET /movie/{movie_id}/images`
- [x] Keywords — `GET /movie/{movie_id}/keywords`
- [x] Latest — `GET /movie/latest`
- [x] Lists — `GET /movie/{movie_id}/lists`
- [x] Now Playing — `GET /movie/now_playing`
- [x] Popular — `GET /movie/popular`
- [x] Recommendations — `GET /movie/{movie_id}/recommendations`
- [x] Release Dates — `GET /movie/{movie_id}/release_dates`
- [x] Reviews — `GET /movie/{movie_id}/reviews`
- [x] Similar — `GET /movie/{movie_id}/similar`
- [x] Top Rated — `GET /movie/top_rated`
- [x] Translations — `GET /movie/{movie_id}/translations`
- [x] Upcoming — `GET /movie/upcoming`
- [x] Videos — `GET /movie/{movie_id}/videos`
- [x] Watch Providers — `GET /movie/{movie_id}/watch/providers`
- [x] Add Rating 🔒 — `POST /movie/{movie_id}/rating`
- [x] Delete Rating 🔒 — `DELETE /movie/{movie_id}/rating`

## Networks
- [x] Details — `GET /network/{network_id}`
- [x] Alternative Names — `GET /network/{network_id}/alternative_names`
- [x] Images — `GET /network/{network_id}/images`

## People
- [x] Details — `GET /person/{person_id}`
- [x] Combined Credits — `GET /person/{person_id}/combined_credits`
- [x] External IDs — `GET /person/{person_id}/external_ids`
- [x] Images — `GET /person/{person_id}/images`
- [x] Latest — `GET /person/latest`
- [x] Movie Credits — `GET /person/{person_id}/movie_credits`
- [x] Popular — `GET /person/popular`
- [x] TV Credits — `GET /person/{person_id}/tv_credits`
- [x] Translations — `GET /person/{person_id}/translations`

## Reviews
- [x] Details — `GET /review/{review_id}`

## Search
- [x] Collections — `GET /search/collection`
- [x] Companies — `GET /search/company`
- [x] Keywords — `GET /search/keyword`
- [x] Movies — `GET /search/movie`
- [x] Multi — `GET /search/multi`
- [x] People — `GET /search/person`
- [x] TV Shows — `GET /search/tv`

## Trending
- [x] All — `GET /trending/all/{time_window}`
- [x] Movies — `GET /trending/movie/{time_window}`
- [x] People — `GET /trending/person/{time_window}`
- [x] TV Shows — `GET /trending/tv/{time_window}`

## TV Series
- [x] Details — `GET /tv/{series_id}`
- [x] Account States 🔒 — `GET /tv/{series_id}/account_states`
- [x] Aggregate Credits — `GET /tv/{series_id}/aggregate_credits`
- [x] Alternative Titles — `GET /tv/{series_id}/alternative_titles`
- [x] Content Ratings — `GET /tv/{series_id}/content_ratings`
- [x] Credits — `GET /tv/{series_id}/credits`
- [x] Episode Groups — `GET /tv/{series_id}/episode_groups`
- [x] External IDs — `GET /tv/{series_id}/external_ids`
- [x] Images — `GET /tv/{series_id}/images`
- [x] Keywords — `GET /tv/{series_id}/keywords`
- [x] Latest — `GET /tv/latest`
- [x] Lists — `GET /tv/{series_id}/lists`
- [x] Recommendations — `GET /tv/{series_id}/recommendations`
- [x] Reviews — `GET /tv/{series_id}/reviews`
- [x] Screened Theatrically — `GET /tv/{series_id}/screened_theatrically`
- [x] Similar — `GET /tv/{series_id}/similar`
- [x] Translations — `GET /tv/{series_id}/translations`
- [x] Videos — `GET /tv/{series_id}/videos`
- [x] Watch Providers — `GET /tv/{series_id}/watch/providers`
- [x] Popular — `GET /tv/popular`
- [x] Top Rated — `GET /tv/top_rated`
- [x] On The Air — `GET /tv/on_the_air`
- [x] Airing Today — `GET /tv/airing_today`
- [x] Add Rating 🔒 — `POST /tv/{series_id}/rating`
- [x] Delete Rating 🔒 — `DELETE /tv/{series_id}/rating`

## TV Seasons
- [x] Details — `GET /tv/{series_id}/season/{season_number}`
- [x] Account States 🔒 — `GET /tv/{series_id}/season/{season_number}/account_states`
- [x] Aggregate Credits — `GET /tv/{series_id}/season/{season_number}/aggregate_credits`
- [x] Credits — `GET /tv/{series_id}/season/{season_number}/credits`
- [x] External IDs — `GET /tv/{series_id}/season/{season_number}/external_ids`
- [x] Images — `GET /tv/{series_id}/season/{season_number}/images`
- [x] Translations — `GET /tv/{series_id}/season/{season_number}/translations`
- [x] Videos — `GET /tv/{series_id}/season/{season_number}/videos`

## TV Episodes
- [x] Details — `GET /tv/{series_id}/season/{season_number}/episode/{episode_number}`
- [x] Account States 🔒 — `.../episode/{episode_number}/account_states`
- [x] Credits — `.../episode/{episode_number}/credits`
- [x] External IDs — `.../episode/{episode_number}/external_ids`
- [x] Images — `.../episode/{episode_number}/images`
- [x] Translations — `.../episode/{episode_number}/translations`
- [ ] Videos — `.../episode/{episode_number}/videos`
- [ ] Add Rating 🔒 — `POST .../episode/{episode_number}/rating`
- [ ] Delete Rating 🔒 — `DELETE .../episode/{episode_number}/rating`

## TV Episode Groups
- [ ] Details — `GET /tv/episode_group/{id}`

## Watch Providers
- [ ] Available Regions — `GET /watch/providers/regions`
- [ ] Movie Providers — `GET /watch/providers/movie`
- [ ] TV Providers — `GET /watch/providers/tv`
