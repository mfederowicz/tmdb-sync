package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
)

// CompaniesDetailsHandler handles `companies -a details -i <company_id>`.
type CompaniesDetailsHandler struct {
	CompanyID int64
}

// Handle fetches details for a single company.
func (h CompaniesDetailsHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	company, _, err := client.Companies.GetCompany(ctx, h.CompanyID)
	if err != nil {
		return nil, err
	}
	return company, nil
}
