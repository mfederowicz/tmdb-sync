package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
)

// ConfigurationDetailsHandler handles `configuration -a details`.
type ConfigurationDetailsHandler struct{}

// Handle fetches TMDB's API configuration.
func (ConfigurationDetailsHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	config, _, err := client.Configuration.GetAPIConfiguration(ctx)
	if err != nil {
		return nil, err
	}
	return config, nil
}
