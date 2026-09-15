package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
)

// CertificationsTVHandler handles `certifications -a tv`.
type CertificationsTVHandler struct{}

// Handle fetches TV certifications, by country.
func (CertificationsTVHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	certifications, _, err := client.Certifications.GetTVCertifications(ctx)
	if err != nil {
		return nil, err
	}
	return certifications, nil
}
