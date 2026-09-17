package cli

import (
	"fmt"
	"runtime/debug"
)

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
	}

	return fmt.Sprintf("Version:        %s\nCommit:         %s\nBuilt           %s by %s", v, c, d, by)
}
