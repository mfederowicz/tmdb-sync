package cmds

import (
	"context"
	"flag"
	"fmt"

	"github.com/mfederowicz/tmdb-sync/cfg"
	"github.com/mfederowicz/tmdb-sync/handlers"
	"github.com/mfederowicz/tmdb-sync/internal"
	"github.com/mfederowicz/tmdb-sync/str"

	"github.com/spf13/afero"
)

// TVEpisodeGroupsCmd is the "tv-episode-groups" module.
var TVEpisodeGroupsCmd = &Command{
	Name:   "tv-episode-groups",
	Abbrev: "tv-episode-groups",
	Short:  "TV episode group details",
	Exec:   execTVEpisodeGroups,
}

func execTVEpisodeGroups(fs afero.Fs, client *internal.Client, config *cfg.Config, _ *str.Options, args []string) error {
	flagSet := flag.NewFlagSet("tv-episode-groups", flag.ContinueOnError)
	action := flagSet.String("a", "", "action: details (required)")
	id := flagSet.String("i", "", "episode group id, required for all actions")
	if err := flagSet.Parse(args); err != nil {
		return err
	}

	var handler handlers.Handler
	switch *action {
	case "":
		return fmt.Errorf("tv-episode-groups: -a is required (action: details)")
	case "details":
		if *id == "" {
			return fmt.Errorf("tv-episode-groups: -i <episode_group_id> is required for -a details")
		}
		handler = handlers.TVEpisodeGroupDetailsHandler{ID: *id}
	default:
		return fmt.Errorf("tv-episode-groups: unknown action %q", *action)
	}

	result, err := handler.Handle(context.Background(), client)
	if err != nil {
		return err
	}

	params := []string{fmt.Sprintf("id-%s", *id)}
	return writeResult(fs, config, "tv-episode-groups", *action, result, params...)
}
