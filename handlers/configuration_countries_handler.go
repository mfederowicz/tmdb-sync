package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
)

// ConfigurationCountriesHandler handles `configuration -a countries`.
type ConfigurationCountriesHandler struct{}

// Handle fetches TMDB's list of countries.
func (ConfigurationCountriesHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	countries, _, err := client.Configuration.GetCountries(ctx)
	if err != nil {
		return nil, err
	}
	return countries, nil
}
