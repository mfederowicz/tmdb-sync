package internal

import (
	"context"
	"net/http"

	"github.com/mfederowicz/tmdb-sync/str"
)

// ConfigurationService handles the TMDB /configuration endpoint.
type ConfigurationService Service

// GetAPIConfiguration returns TMDB's image base URLs/sizes and change keys.
//
// Api docs: https://developer.themoviedb.org/reference/configuration-details
func (s *ConfigurationService) GetAPIConfiguration(ctx context.Context) (*str.Configuration, *str.Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, "configuration", nil)
	if err != nil {
		return nil, nil, err
	}

	config := new(str.Configuration)
	resp, err := s.client.Do(ctx, req, config)
	if err != nil {
		return nil, resp, err
	}

	return config, resp, nil
}
