package cli

import "fmt"

// version info, overridden at build time via -ldflags (see Makefile).
var (
	date    = "unknown"
	builtBy = "unknown"
	version = "dev"
)

// GenAppVersion returns a human readable version string.
func GenAppVersion() string {
	return fmt.Sprintf("tmdb-sync %s (built %s by %s)", version, date, builtBy)
}
