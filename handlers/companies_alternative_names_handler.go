package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
)

// CompaniesAlternativeNamesHandler handles `companies -a alternative-names -i <company_id>`.
type CompaniesAlternativeNamesHandler struct {
	CompanyID int64
}

// Handle fetches the alternative names for a single company.
func (h CompaniesAlternativeNamesHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	names, _, err := client.Companies.GetCompanyAlternativeNames(ctx, h.CompanyID)
	if err != nil {
		return nil, err
	}
	return names, nil
}
