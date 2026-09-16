# tmdb-sync — Architecture

If in doubt about a pattern, follow the existing shape rather than inventing something new — see
"Adding a new module" below.

## Repo / git conventions

- Default branch is `main`, pushed to a public GitHub remote.
- One module per branch/PR; if a module has more than ~3 endpoints, split it into smaller
  branches rather than landing it all at once.
- Only commit/push when explicitly asked.

## Package layout

- `consts/` — shared constants (`glob.go`, `keys.go`, `numbers.go`): zero values, default
  per_page/pages_limit, file perms, header names, flag usage strings.
- `cfg/` — `config.go`: TOML-backed `Config` struct + `InitConfig`/`ReadConfigFromFile`/
  `DefaultConfig`/`normalizeConfig` (flag-merge pattern: defaults < config file < CLI flags).
  `options.go`: `OptionsFromConfig` turns a `*Config` into request headers (`Authorization:
  Bearer <read_access_token>`, or `api_key` for the query string) and loads any previously
  persisted `str.Session` from `session_path`.
- `str/` — plain data structs, no logic beyond trivial helpers (e.g. `Session.Valid()`):
  `options.go`, `response.go` (wraps `*http.Response`), `errors.go` (TMDB's
  `{status_code, status_message, success}` error body), `request_token.go`, `session.go`, and one
  file per API resource shape (`movie.go`, `movies.go`, `configuration.go`, ...).
- `internal/` — the TMDB API client.
  - `client.go`: `Client` wraps `*http.Client`, `BaseURL` (`https://api.themoviedb.org/3/`),
    `AuthURL` (reserved, unused until the v4 phase), a `headers map[string]any`, and a `common
    Service` used to build every API-area service via the `(*XService)(&c.common)` conversion
    pattern (mirrors `google/go-github`). `NewRequest` builds JSON requests; appends `api_key` as
    a query param if present in headers. `Do`/`BareDo` decode JSON responses into `str.Response`
    and map non-2xx into `str.ErrorResponse` / `AbuseRateLimitError` (429, or a request
    short-circuited locally before a known `RateLimitReset`).
  - `service.go`: `Service{client *Client}` — the shared base every `*Service` type converts from.
  - `auth_service.go`: TMDB **v3** auth only (`CreateRequestToken`, `CreateSession` against
    `authentication/token/new` / `authentication/session/new`), exposed as `Client.Auth`.
  - `<x>_service.go` per module (`movies_service.go`, `configuration_service.go`, ...), exposed as
    `Client.<X>`.
- `cli/` — interactive helpers: `session.go` (`HandleToken`: checks `options.Session.Valid()`,
  else runs the v3 request-token → browser-approval → session flow and persists the result),
  `browser.go` (best-effort OS-specific browser launch + "press enter" wait), `version.go`.
- `cmds/` + `handlers/` — command dispatch.
  - `cmds.Command{Name, Abbrev, Short, Exec}`; `cmds.ModulesRuntime` resolves the first CLI arg to
    a command by exact `Name` or exact `Abbrev` match only (no prefix matching). `Commands` (the registry) is built
    inside `runtime.go`'s `init()`, **not** a package-level var literal — a literal creates an
    initialization-order cycle with `HelpCmd`/`HelpFunc` (`HelpFunc` prints `Commands`).
  - Each `cmds/command_<module>.go` owns its own `flag.FlagSet` for that module's action args
    (`-a details`, `-i <id>`, `-p <page>`, ...) and picks a `handlers.Handler` to run.
  - `handlers.Handler` is `Handle(ctx, client) (any, error)` — one small struct per action
    (`MoviesDetailsHandler{MovieID}`, `MoviesPopularHandler{Page}`, ...), calling exactly one
    `internal` service method. `cmds/output.go`'s `writeResult` then saves the result as JSON (see
    Output below) instead of dumping it to stdout.
- `printer/` — thin `fmt`-over-`io.Writer` wrapper; `printer.Stdout` is swappable so tests can
  capture CLI output without touching global state beyond a package var. Used only for short
  status lines now (e.g. `writeResult`'s "wrote <path>"), not for dumping API responses.
- `uri/` — query-string helpers (`AddPage`, `SanitizeURL` to redact `api_key` from error/log
  output).
- `writer/` — `BuildFilename(module, action, params...)` (e.g.
  `BuildFilename("movies", "details", "id-550")` → `movies_details_id-550.json`) and
  `WriteJSON(fs, dir, filename, v)` (creates `dir` if needed, writes indented JSON via afero).
  `cmds/output.go`'s `writeResult` is the one call site every command action should go through.
- `main.go` — flags (`-v`, `-version`, `-c`), `cfg.InitConfig` → `cfg.OptionsFromConfig` →
  `internal.NewClient` → `client.UpdateHeaders` → `cmds.ModulesRuntime`. **Does not** call
  `cli.HandleToken` unconditionally — public v3 endpoints don't need a session; only a future
  account-specific command should trigger the browser-approval flow, on demand.

## Adding a new module (the repeatable unit of work)

For an endpoint like `GET /tv/{tv_id}`:
1. `str/tv.go` — response struct.
2. `internal/tv_service.go` — `TVService.GetTV(ctx, id) (*str.TV, *str.Response, error)`, wired
   onto `Client` in `client.go`'s field list + `initialize()`.
3. `handlers/tv_details_handler.go` — `TVDetailsHandler{TVID}.Handle(...)`.
4. `cmds/command_tv.go` (new, or extend if a `tv` command already exists) — flag parsing + switch
   on `-a`, add `TVCmd` to `Commands` in `cmds/runtime.go`'s `init()`.
5. Tests: `internal/tv_service_test.go` (httptest-backed), extend `cmds` dispatch tests if a new
   top-level command was added.
6. Tick it off in `API_COVERAGE.md`; add a `docs/tv.md` entry if the shape is non-trivial.

No existing file needs to change except the two registries (`Client` field list +
`cmds.Commands`).

## Pagination: always capped, never a dedicated "GetAll" function

TMDB does **not** put pagination info in HTTP headers (unlike some other APIs) — list endpoints
return `page`, `total_pages`, `total_results` in the JSON body itself. TMDB also fixes the page
*size* at 20 items server-side; `cfg.Config.PerPage` is reserved/unused because there's no API
lever for it. The only real control is how many *pages* to walk, `cfg.Config.PagesLimit`.

The generic building block is `internal.FetchAllPages[T]` (`internal/paginate.go`): given a
`fetch(ctx, page) (PageResult[T], error)` closure, it walks pages 1..N until either TMDB reports
no more pages (`total_pages`) or `pagesLimit` (0 = unlimited) is hit, and returns the concatenated
`[]T`.

**There is no separate `Get<X>` (single page) + `GetAll<X>` (every page) pair, and no `-all`
flag.** Each list endpoint gets exactly one public method that always applies the cap — see
`internal/movies_service.go`: `getPopularMoviesPage` (unexported, one HTTP call) is the plumbing;
`GetPopularMovies(ctx, pagesLimit)` is the only exported entry point, built on `FetchAllPages`, and
always returns however many pages `min(total_pages, pagesLimit)` allows. A command
(`cmds/command_movies.go`) just passes `config.PagesLimit` through — no extra flag needed. Set
`pages_limit = 0` in config for "fetch every page TMDB has".

## Output: JSON files, not stdout dumps

Every command action ends by calling `writeResult(fs, config, module, action, result, params...)`
(`cmds/output.go`), which:
1. Builds a filename via `writer.BuildFilename(module, action, params...)` —
   `<module>_<action>[_<param>...].json`, e.g. `movies_details_id-550.json`,
   `movies_popular_page-1.json`, or `movies_popular_all.json` for a `-all` fetch.
2. Writes indented JSON under `config.OutputDir` (default: current directory) via
   `writer.WriteJSON`.
3. Prints a one-line confirmation (`wrote <path>`) instead of the raw JSON — the file is the
   product, not the stdout stream.

`params` are the action's identifying arguments as `key-value` strings (`fmt.Sprintf("id-%d",
movieID)`, `fmt.Sprintf("page-%d", page)`, or the literal `"all"`) — pass whatever makes two
different invocations of the same action not collide on one filename.

## Auth model (v3 now, v4 later without rewriting)

TMDB v3 has its own request-token/session flow, independent of v4:
1. `GET /3/authentication/token/new` → `request_token`.
2. User approves at `https://www.themoviedb.org/authenticate/{request_token}` (browser, via
   `cli/browser.go`).
3. `POST /3/authentication/session/new` (with the approved `request_token`) → `session_id`.
4. Session persisted to `session_path` (JSON file).

`Client.Auth` (v3) and a future `Client.AuthV4` are separate fields/services/files by design —
adding v4 means new files (`internal/auth_v4_service.go`, a new `cli` flow, a `cfg.AuthVersion`
switch already reserved in `Config`), never edits to the v3 files above.

## Error handling

- Typed errors: `str.ErrorResponse` (TMDB's JSON error body) for normal 4xx/5xx, and
  `internal.AbuseRateLimitError` for 429 / locally-known rate-limit windows.
- `internal.CheckResponse` is the single place that maps an `*http.Response` to one of these (or
  `nil` for 2xx) — new service methods don't need their own error mapping.

## Testing strategy

- `cfg`: table-driven tests over `ReadConfigFromFile`/`MergeConfigs`/`normalizeConfig` using
  `afero.NewMemMapFs()` — no real filesystem I/O.
- `internal`: `httptest.Server`-backed tests per service method (request shape, header/query auth,
  JSON decoding, error mapping). This is the main safety net per new module.
- `cmds`: dispatch resolution (`ModulesRuntime` exact name/abbrev matching, unknown command) using a
  fake `Commands` slice and `printer.Stdout` swapped to a buffer.
- No integration tests against the real TMDB API (would require a real key); rely on `httptest`
  mocks plus manual smoke testing during development.
