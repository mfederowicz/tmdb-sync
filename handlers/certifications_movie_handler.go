package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
)

// CertificationsMovieHandler handles `certifications -a movie`.
type CertificationsMovieHandler struct{}

// Handle fetches movie certifications, by country.
func (CertificationsMovieHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	certifications, _, err := client.Certifications.GetMovieCertifications(ctx)
	if err != nil {
		return nil, err
	}
	return certifications, nil
}
