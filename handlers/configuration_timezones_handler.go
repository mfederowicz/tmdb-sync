package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
)

// ConfigurationTimezonesHandler handles `configuration -a timezones`.
type ConfigurationTimezonesHandler struct{}

// Handle fetches TMDB's list of timezones.
func (ConfigurationTimezonesHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	timezones, _, err := client.Configuration.GetTimezones(ctx)
	if err != nil {
		return nil, err
	}
	return timezones, nil
}
