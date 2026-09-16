package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
)

// ConfigurationLanguagesHandler handles `configuration -a languages`.
type ConfigurationLanguagesHandler struct{}

// Handle fetches TMDB's list of languages.
func (ConfigurationLanguagesHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	languages, _, err := client.Configuration.GetLanguages(ctx)
	if err != nil {
		return nil, err
	}
	return languages, nil
}
