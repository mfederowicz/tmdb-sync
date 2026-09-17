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

// peopleActionsHelp lists every people action, shared between the -a flag's
// usage string and the "-a is required" error so both stay in sync.
const peopleActionsHelp = "details"

// peopleIDActionsHelp lists the people actions that require -i.
const peopleIDActionsHelp = "details"

// PeopleCmd is the "people" module.
var PeopleCmd = &Command{
	Name:   "people",
	Abbrev: "p",
	Short:  "person details",
	Exec:   execPeople,
}

func execPeople(fs afero.Fs, client *internal.Client, config *cfg.Config, _ *str.Options, args []string) error {
	flagSet := flag.NewFlagSet("people", flag.ContinueOnError)
	action := flagSet.String("a", "", "action: "+peopleActionsHelp+" (required)")
	personID := flagSet.Int64("i", 0, "person id, required for -a "+peopleIDActionsHelp)
	if err := flagSet.Parse(args); err != nil {
		return err
	}

	var handler handlers.Handler
	switch *action {
	case "":
		return fmt.Errorf("people: -a is required (action: %s)", peopleActionsHelp)
	case "details":
		if *personID == 0 {
			return fmt.Errorf("people: -i <person_id> is required for -a details")
		}
		handler = handlers.PeopleDetailsHandler{PersonID: *personID}
	default:
		return fmt.Errorf("people: unknown action %q", *action)
	}

	result, err := handler.Handle(context.Background(), client)
	if err != nil {
		return err
	}

	params := []string{fmt.Sprintf("id-%d", *personID)}
	return writeResult(fs, config, "people", *action, result, params...)
}
