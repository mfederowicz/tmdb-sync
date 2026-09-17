package internal

import (
	"context"
	"fmt"
	"net/http"

	"github.com/mfederowicz/tmdb-sync/str"
)

// CompaniesService handles communication with the /company endpoints of the TMDB API.
type CompaniesService Service

// GetCompany fetches details for a single company by TMDB id.
//
// Api docs: https://developer.themoviedb.org/reference/company-details
func (s *CompaniesService) GetCompany(ctx context.Context, companyID int64) (*str.Company, *str.Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("company/%d", companyID), nil)
	if err != nil {
		return nil, nil, err
	}

	company := new(str.Company)
	resp, err := s.client.Do(ctx, req, company)
	if err != nil {
		return nil, resp, err
	}

	return company, resp, nil
}

// GetCompanyAlternativeNames fetches the alternative names for a single company by TMDB id.
//
// Api docs: https://developer.themoviedb.org/reference/company-alternative-names
func (s *CompaniesService) GetCompanyAlternativeNames(ctx context.Context, companyID int64) (*str.AlternativeNames, *str.Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("company/%d/alternative_names", companyID), nil)
	if err != nil {
		return nil, nil, err
	}

	names := new(str.AlternativeNames)
	resp, err := s.client.Do(ctx, req, names)
	if err != nil {
		return nil, resp, err
	}

	return names, resp, nil
}

// GetCompanyImages fetches the logos for a single company by TMDB id.
//
// Api docs: https://developer.themoviedb.org/reference/company-images
func (s *CompaniesService) GetCompanyImages(ctx context.Context, companyID int64) (*str.Images, *str.Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("company/%d/images", companyID), nil)
	if err != nil {
		return nil, nil, err
	}

	images := new(str.Images)
	resp, err := s.client.Do(ctx, req, images)
	if err != nil {
		return nil, resp, err
	}

	return images, resp, nil
}
