package handlers

import (
	"context"

	"github.com/mfederowicz/tmdb-sync/cfg"
	"github.com/mfederowicz/tmdb-sync/cli"
	"github.com/mfederowicz/tmdb-sync/internal"

	"github.com/spf13/afero"
)

// AuthLoginHandler handles `auth -v4 -a login`: the interactive request
// token -> browser approval -> access token flow, persisting the result.
type AuthLoginHandler struct {
	Fs     afero.Fs
	Config *cfg.Config
}

// Handle runs the interactive v4 login.
func (h AuthLoginHandler) Handle(_ context.Context, client *internal.Client) (any, error) {
	accessToken, err := cli.CreateAccessTokenInteractively(h.Fs, h.Config, client)
	if err != nil {
		return nil, err
	}
	return accessToken, nil
}
