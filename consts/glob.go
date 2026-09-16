// Package consts holds shared constants used across the app.
package consts

// generic constants
const (
	EmptyString = ""
	ZeroValue   = 0
	// X644 is the default file permission for written json files.
	X644 = 0o644
)

// usage strings for CLI flags
const (
	VerboseUsage = "verbose output"
	VersionUsage = "print version and exit"
	ConfigUsage  = "path to config file"
)

// TMDB media_type values, shared by modules that accept both movie and TV
// endpoints (e.g. account favorites/watchlist).
const (
	MediaTypeMovie = "movie"
	MediaTypeTV    = "tv"
)
