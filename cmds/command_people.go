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
const peopleActionsHelp = "details, combined-credits, external-ids, images, latest, movie-credits, popular, tv-credits"

// peopleIDActionsHelp lists the people actions that require -i.
const peopleIDActionsHelp = "details, combined-credits, external-ids, images, movie-credits, tv-credits"

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
	pagesLimit := flagSet.Int("pages-limit", config.PagesLimit, "pages limit, used by -a popular (default: pages_limit from config, 0 = unlimited)")
	if err := flagSet.Parse(args); err != nil {
		return err
	}

	var handler handlers.Handler
	var params []string
	switch *action {
	case "":
		return fmt.Errorf("people: -a is required (action: %s)", peopleActionsHelp)
	case "details":
		if *personID == 0 {
			return fmt.Errorf("people: -i <person_id> is required for -a details")
		}
		handler = handlers.PeopleDetailsHandler{PersonID: *personID}
		params = []string{fmt.Sprintf("id-%d", *personID)}
	case "combined-credits":
		if *personID == 0 {
			return fmt.Errorf("people: -i <person_id> is required for -a combined-credits")
		}
		handler = handlers.PeopleCombinedCreditsHandler{PersonID: *personID}
		params = []string{fmt.Sprintf("id-%d", *personID)}
	case "external-ids":
		if *personID == 0 {
			return fmt.Errorf("people: -i <person_id> is required for -a external-ids")
		}
		handler = handlers.PeopleExternalIDsHandler{PersonID: *personID}
		params = []string{fmt.Sprintf("id-%d", *personID)}
	case "images":
		if *personID == 0 {
			return fmt.Errorf("people: -i <person_id> is required for -a images")
		}
		handler = handlers.PeopleImagesHandler{PersonID: *personID}
		params = []string{fmt.Sprintf("id-%d", *personID)}
	case "latest":
		handler = handlers.PeopleLatestHandler{}
	case "movie-credits":
		if *personID == 0 {
			return fmt.Errorf("people: -i <person_id> is required for -a movie-credits")
		}
		handler = handlers.PeopleMovieCreditsHandler{PersonID: *personID}
		params = []string{fmt.Sprintf("id-%d", *personID)}
	case "popular":
		handler = handlers.PeoplePopularHandler{PagesLimit: *pagesLimit}
	case "tv-credits":
		if *personID == 0 {
			return fmt.Errorf("people: -i <person_id> is required for -a tv-credits")
		}
		handler = handlers.PeopleTVCreditsHandler{PersonID: *personID}
		params = []string{fmt.Sprintf("id-%d", *personID)}
	default:
		return fmt.Errorf("people: unknown action %q", *action)
	}

	result, err := handler.Handle(context.Background(), client)
	if err != nil {
		return err
	}

	return writeResult(fs, config, "people", *action, result, params...)
}
