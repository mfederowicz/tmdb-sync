package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
)

// ConfigurationPrimaryTranslationsHandler handles `configuration -a primary-translations`.
type ConfigurationPrimaryTranslationsHandler struct{}

// Handle fetches TMDB's list of primary translation locales.
func (ConfigurationPrimaryTranslationsHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	translations, _, err := client.Configuration.GetPrimaryTranslations(ctx)
	if err != nil {
		return nil, err
	}
	return translations, nil
}
