package cli

import (
	"context"
	"fmt"

	"github.com/mfederowicz/tmdb-sync/cfg"
	"github.com/mfederowicz/tmdb-sync/internal"
	"github.com/mfederowicz/tmdb-sync/printer"
	"github.com/mfederowicz/tmdb-sync/str"

	"github.com/spf13/afero"
)

const tmdbApproveURLV4Fmt = "https://www.themoviedb.org/auth/access?request_token=%s"

// CreateAccessTokenInteractively runs the v4 request-token -> browser
// approval -> access-token flow and persists the result to
// config.AccessTokenPath.
func CreateAccessTokenInteractively(fs afero.Fs, config *cfg.Config, client *internal.Client) (*str.AccessTokenV4, error) {
	ctx := context.Background()

	requestToken, _, err := client.Auth.CreateRequestTokenV4(ctx, "")
	if err != nil {
		return nil, fmt.Errorf("create v4 request token: %w", err)
	}

	approveURL := fmt.Sprintf(tmdbApproveURLV4Fmt, requestToken.RequestToken)
	printer.Println("Approve tmdb-sync access in your browser:", approveURL)
	OpenBrowser(approveURL)
	printer.Println("Press Enter once you have approved access...")
	WaitForEnter()

	accessToken, _, err := client.Auth.CreateAccessTokenV4(ctx, requestToken.RequestToken)
	if err != nil {
		return nil, fmt.Errorf("create v4 access token: %w", err)
	}
	if !accessToken.Valid() {
		return nil, fmt.Errorf("tmdb refused to create an access token for the approved request token")
	}

	if err := cfg.WriteAccessToken(fs, config.AccessTokenPath, accessToken); err != nil {
		return nil, fmt.Errorf("persist access token: %w", err)
	}

	return accessToken, nil
}
