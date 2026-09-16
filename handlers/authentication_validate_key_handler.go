package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/internal"
)

// AuthenticationValidateKeyHandler handles `authentication -a validate-key`.
type AuthenticationValidateKeyHandler struct{}

// Handle confirms the configured api_key/read_access_token is valid.
func (AuthenticationValidateKeyHandler) Handle(ctx context.Context, client *internal.Client) (any, error) {
	status, _, err := client.Auth.ValidateKey(ctx)
	if err != nil {
		return nil, err
	}
	return status, nil
}
