package cmds

import (
	"flag"
	"strings"

	"github.com/mfederowicz/tmdb-sync/cfg"
	"github.com/mfederowicz/tmdb-sync/internal"
	"github.com/mfederowicz/tmdb-sync/printer"
	"github.com/mfederowicz/tmdb-sync/str"

	"github.com/spf13/afero"
)

// Commands is the list of all registered modules, populated in init() to
// avoid an initialization-order cycle with HelpCmd (which prints Commands).
var Commands []*Command

func init() {
	Commands = []*Command{
		CertificationsCmd,
		ConfigurationCmd,
		MoviesCmd,
		HelpCmd,
	}
}

// runtime consts
const (
	foundOne = 1
	notFound = 0
)

func runFoundedModule(cmd *Command, fs afero.Fs, client *internal.Client, config *cfg.Config, options *str.Options, args []string) {
	if err := cmd.Exec(fs, client, config, options, args); err != nil {
		printer.Println(err)
	}
}

// ModulesRuntime core function for process commands
func ModulesRuntime(args []string, fs afero.Fs, config *cfg.Config, client *internal.Client, options *str.Options) {
	var found []*Command
	sub, args := args[notFound], args[foundOne:]

find:
	for _, cmd := range Commands {
		if sub == cmd.Abbrev {
			found = []*Command{cmd}
			break find
		}
		if strings.HasPrefix(cmd.Name, sub) {
			found = append(found, cmd)
		}
	}

	switch cnt := len(found); cnt {
	case foundOne:
		runFoundedModule(found[0], fs, client, config, options, args)
	case notFound:
		printer.Printf("error: unknown command %q\n\n", sub)
		flag.Usage()
	default:
		printer.Printf("error: non-unique command prefix %q (matched %d commands)\n\n", sub, cnt)
		flag.Usage()
	}
}
