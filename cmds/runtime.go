package cmds

import (
	"flag"

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
		AccountCmd,
		CertificationsCmd,
		ChangesCmd,
		CollectionsCmd,
		CompaniesCmd,
		ConfigurationCmd,
		CreditsCmd,
		DiscoverCmd,
		FindCmd,
		GenresCmd,
		KeywordsCmd,
		ListsCmd,
		MoviesCmd,
		NetworksCmd,
		PeopleCmd,
		ReviewsCmd,
		SearchCmd,
		TrendingCmd,
		TVCmd,
		TVSeasonsCmd,
		TVEpisodesCmd,
		TVEpisodeGroupsCmd,
		WatchProvidersCmd,
		HelpCmd,
	}
}

// runtime consts
const (
	foundOne = 1
	notFound = 0
)

func runFoundedModule(cmd *Command, fs afero.Fs, client *internal.Client, config *cfg.Config, options *str.Options, args []string) bool {
	if err := cmd.Exec(fs, client, config, options, args); err != nil {
		printer.Println(err)
		return false
	}
	return true
}

// ModulesRuntime core function for process commands. It returns false when
// the command failed (unknown module or a module-level error), so main can
// exit with a non-zero status.
func ModulesRuntime(args []string, fs afero.Fs, config *cfg.Config, client *internal.Client, options *str.Options) bool {
	sub, args := args[notFound], args[foundOne:]

	for _, cmd := range Commands {
		if sub == cmd.Name || sub == cmd.Abbrev {
			return runFoundedModule(cmd, fs, client, config, options, args)
		}
	}

	printer.Printf("error: unknown command %q\n\n", sub)
	flag.Usage()
	return false
}
