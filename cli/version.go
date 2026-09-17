package cli

import (
	"fmt"
	"regexp"
	"runtime/debug"
)

// pseudoVersionRE matches Go pseudo-versions, e.g. v0.9.1-0.20260917155637-65660c458192,
// capturing the embedded timestamp and abbreviated commit hash.
var pseudoVersionRE = regexp.MustCompile(`-(\d{14})-([0-9a-f]{12})$`)

// version info, overridden at build time via -ldflags (see Makefile).
var (
	date    = "unknown"
	builtBy = "unknown"
	version = "dev"
	commit  = "unknown"
)

// GenAppVersion returns a human readable version string.
func GenAppVersion() string {
	v, c, d, by := version, commit, date, builtBy

	if v == "dev" {
		if info, ok := debug.ReadBuildInfo(); ok {
			if info.Main.Version != "" && info.Main.Version != "(devel)" {
				v = info.Main.Version
			}
			for _, s := range info.Settings {
				switch s.Key {
				case "vcs.revision":
					c = s.Value
				case "vcs.time":
					d = s.Value
				}
			}
			by = "go install"
		}

		if c == "unknown" || d == "unknown" {
			if m := pseudoVersionRE.FindStringSubmatch(v); m != nil {
				if c == "unknown" {
					c = m[2]
				}
				if d == "unknown" {
					d = m[1]
				}
			}
		}
	}

	return fmt.Sprintf("Version:        %s\nCommit:         %s\nBuilt           %s by %s", v, c, d, by)
}
