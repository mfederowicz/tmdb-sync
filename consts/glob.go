// Package consts holds shared constants used across the app.
package consts

// generic constants
const (
	EmptyString = ""
	ZeroValue   = 0
	// X644 is the default file permission for written json files.
	X644 = 0o644
	// X600 is the file permission for written files holding secrets (tokens).
	X600 = 0o600
	// X700 is the directory permission for directories holding those secret files.
	X700 = 0o700
)

// usage strings for CLI flags
const (
	VerboseUsage = "verbose output"
	VersionUsage = "print version and exit"
	ConfigUsage  = "path to config file"
	DebugUsage   = "print the full request URL for each API call"
)

// TMDB media_type values, shared by modules that accept both movie and TV
// endpoints (e.g. account favorites/watchlist).
const (
	MediaTypeMovie = "movie"
	MediaTypeTV    = "tv"
)

// TMDB v4 sort_by values for the account favorites/rated/watchlist endpoints.
const (
	SortByCreatedAtAsc  = "created_at.asc"
	SortByCreatedAtDesc = "created_at.desc"
)
