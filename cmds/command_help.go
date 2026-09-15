package cmds

import (
	"github.com/mfederowicz/tmdb-sync/cfg"
	"github.com/mfederowicz/tmdb-sync/internal"
	"github.com/mfederowicz/tmdb-sync/printer"
	"github.com/mfederowicz/tmdb-sync/str"

	"github.com/spf13/afero"
)

// HelpCmd prints the list of available modules.
var HelpCmd = &Command{
	Name:   "help",
	Abbrev: "h",
	Short:  "show this help",
	Exec:   execHelp,
}

func execHelp(_ afero.Fs, _ *internal.Client, _ *cfg.Config, _ *str.Options, _ []string) error {
	HelpFunc(nil)
	return nil
}

// HelpFunc prints usage information listing all registered commands.
func HelpFunc(_ *Command) {
	printer.Println("Usage: tmdb-sync [-c config] [-v] <module> [args...]")
	printer.Println("")
	printer.Println("Modules:")
	for _, cmd := range Commands {
		printer.Printf("  %-16s %s\n", cmd.Name, cmd.Short)
	}
}
