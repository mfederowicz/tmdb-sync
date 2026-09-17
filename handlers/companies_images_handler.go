package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
)

// CompaniesImagesHandler handles `companies -a images -i <company_id>`.
type CompaniesImagesHandler struct {
	CompanyID int64
}

// Handle fetches the images for a single company.
func (h CompaniesImagesHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	images, _, err := client.Companies.GetCompanyImages(ctx, h.CompanyID)
	if err != nil {
		return nil, err
	}
	return images, nil
}
