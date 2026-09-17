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

// CollectionsCmd is the "collections" module.
var CollectionsCmd = &Command{
	Name:   "collections",
	Abbrev: "co",
	Short:  "collection details, images, translations",
	Exec:   execCollections,
}

func execCollections(fs afero.Fs, client *internal.Client, config *cfg.Config, _ *str.Options, args []string) error {
	flagSet := flag.NewFlagSet("collections", flag.ContinueOnError)
	action := flagSet.String("a", "", "action: details, images, translations (required)")
	collectionID := flagSet.Int64("i", 0, "collection id, required for all actions")
	language := flagSet.String("language", "", "language, used by -a images")
	includeImageLanguage := flagSet.String("include-image-language", "", "include-image-language, used by -a images")
	if err := flagSet.Parse(args); err != nil {
		return err
	}

	var handler handlers.Handler
	switch *action {
	case "":
		return fmt.Errorf("collections: -a is required (action: details, images, translations)")
	case "details":
		if *collectionID == 0 {
			return fmt.Errorf("collections: -i <collection_id> is required for -a details")
		}
		handler = handlers.CollectionsDetailsHandler{CollectionID: *collectionID}
	case "images":
		if *collectionID == 0 {
			return fmt.Errorf("collections: -i <collection_id> is required for -a images")
		}
		handler = handlers.CollectionsImagesHandler{CollectionID: *collectionID, Language: *language, IncludeImageLanguage: *includeImageLanguage}
	case "translations":
		if *collectionID == 0 {
			return fmt.Errorf("collections: -i <collection_id> is required for -a translations")
		}
		handler = handlers.CollectionsTranslationsHandler{CollectionID: *collectionID}
	default:
		return fmt.Errorf("collections: unknown action %q", *action)
	}

	result, err := handler.Handle(context.Background(), client)
	if err != nil {
		return err
	}

	params := []string{fmt.Sprintf("id-%d", *collectionID)}
	return writeResult(fs, config, "collections", *action, result, params...)
}
