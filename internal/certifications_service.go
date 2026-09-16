package internal

import (
	"context"
	"net/http"

	"github.com/mfederowicz/tmdb-sync/str"
)

// CertificationsService handles the TMDB /certification endpoints.
type CertificationsService Service

// GetMovieCertifications returns movie certifications, by country.
//
// Api docs: https://developer.themoviedb.org/reference/certification-movie-list
func (s *CertificationsService) GetMovieCertifications(ctx context.Context) (*str.Certifications, *str.Response, error) {
	return s.getCertifications(ctx, "certification/movie/list")
}

// GetTVCertifications returns TV certifications, by country.
//
// Api docs: https://developer.themoviedb.org/reference/certification-tv-list
func (s *CertificationsService) GetTVCertifications(ctx context.Context) (*str.Certifications, *str.Response, error) {
	return s.getCertifications(ctx, "certification/tv/list")
}

func (s *CertificationsService) getCertifications(ctx context.Context, path string) (*str.Certifications, *str.Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	certifications := new(str.Certifications)
	resp, err := s.client.Do(ctx, req, certifications)
	if err != nil {
		return nil, resp, err
	}

	return certifications, resp, nil
}
