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

// NetworksCmd is the "networks" module.
var NetworksCmd = &Command{
	Name:   "networks",
	Abbrev: "nw",
	Short:  "network details, alternative names, images",
	Exec:   execNetworks,
}

func execNetworks(fs afero.Fs, client *internal.Client, config *cfg.Config, _ *str.Options, args []string) error {
	flagSet := flag.NewFlagSet("networks", flag.ContinueOnError)
	action := flagSet.String("a", "", "action: details, alternative-names, images (required)")
	networkID := flagSet.Int64("i", 0, "network id, required for all actions")
	if err := flagSet.Parse(args); err != nil {
		return err
	}

	var handler handlers.Handler
	switch *action {
	case "":
		return fmt.Errorf("networks: -a is required (action: details, alternative-names, images)")
	case "details":
		if *networkID == 0 {
			return fmt.Errorf("networks: -i <network_id> is required for -a details")
		}
		handler = handlers.NetworksDetailsHandler{NetworkID: *networkID}
	case "alternative-names":
		if *networkID == 0 {
			return fmt.Errorf("networks: -i <network_id> is required for -a alternative-names")
		}
		handler = handlers.NetworksAlternativeNamesHandler{NetworkID: *networkID}
	case "images":
		if *networkID == 0 {
			return fmt.Errorf("networks: -i <network_id> is required for -a images")
		}
		handler = handlers.NetworksImagesHandler{NetworkID: *networkID}
	default:
		return fmt.Errorf("networks: unknown action %q", *action)
	}

	result, err := handler.Handle(context.Background(), client)
	if err != nil {
		return err
	}

	params := []string{fmt.Sprintf("id-%d", *networkID)}
	return writeResult(fs, config, "networks", *action, result, params...)
}
