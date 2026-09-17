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

// CompaniesCmd is the "companies" module.
var CompaniesCmd = &Command{
	Name:   "companies",
	Abbrev: "cp",
	Short:  "company details, alternative names, images",
	Exec:   execCompanies,
}

func execCompanies(fs afero.Fs, client *internal.Client, config *cfg.Config, _ *str.Options, args []string) error {
	flagSet := flag.NewFlagSet("companies", flag.ContinueOnError)
	action := flagSet.String("a", "", "action: details, alternative-names, images (required)")
	companyID := flagSet.Int64("i", 0, "company id, required for all actions")
	if err := flagSet.Parse(args); err != nil {
		return err
	}

	var handler handlers.Handler
	switch *action {
	case "":
		return fmt.Errorf("companies: -a is required (action: details, alternative-names, images)")
	case "details":
		if *companyID == 0 {
			return fmt.Errorf("companies: -i <company_id> is required for -a details")
		}
		handler = handlers.CompaniesDetailsHandler{CompanyID: *companyID}
	case "alternative-names":
		if *companyID == 0 {
			return fmt.Errorf("companies: -i <company_id> is required for -a alternative-names")
		}
		handler = handlers.CompaniesAlternativeNamesHandler{CompanyID: *companyID}
	case "images":
		if *companyID == 0 {
			return fmt.Errorf("companies: -i <company_id> is required for -a images")
		}
		handler = handlers.CompaniesImagesHandler{CompanyID: *companyID}
	default:
		return fmt.Errorf("companies: unknown action %q", *action)
	}

	result, err := handler.Handle(context.Background(), client)
	if err != nil {
		return err
	}

	params := []string{fmt.Sprintf("id-%d", *companyID)}
	return writeResult(fs, config, "companies", *action, result, params...)
}
