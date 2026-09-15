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

// CertificationsCmd is the "certifications" module.
var CertificationsCmd = &Command{
	Name:   "certifications",
	Abbrev: "cert",
	Short:  "movie/tv certifications, by country",
	Exec:   execCertifications,
}

func execCertifications(fs afero.Fs, client *internal.Client, config *cfg.Config, _ *str.Options, args []string) error {
	flagSet := flag.NewFlagSet("certifications", flag.ContinueOnError)
	action := flagSet.String("a", "movie", "action: movie, tv")
	if err := flagSet.Parse(args); err != nil {
		return err
	}

	var handler handlers.Handler
	switch *action {
	case "movie":
		handler = handlers.CertificationsMovieHandler{}
	case "tv":
		handler = handlers.CertificationsTVHandler{}
	default:
		return fmt.Errorf("certifications: unknown action %q", *action)
	}

	result, err := handler.Handle(context.Background(), client)
	if err != nil {
		return err
	}

	return writeResult(fs, config, "certifications", *action, result)
}
