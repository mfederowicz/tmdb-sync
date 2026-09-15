// Package cli holds interactive CLI helpers: auth flow, browser launch, version.
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

const tmdbApproveURLFmt = "https://www.themoviedb.org/authenticate/%s"

// HandleToken makes sure options.Session holds a valid TMDB v3 session id,
// running the browser-approval flow and persisting the result if not.
func HandleToken(fs afero.Fs, config *cfg.Config, client *internal.Client, options *str.Options) error {
	if options.Session.Valid() {
		return nil
	}

	session, err := createSessionInteractively(fs, config, client)
	if err != nil {
		return err
	}

	options.Session = session
	return nil
}

func createSessionInteractively(fs afero.Fs, config *cfg.Config, client *internal.Client) (*str.Session, error) {
	ctx := context.Background()

	requestToken, _, err := client.Auth.CreateRequestToken(ctx)
	if err != nil {
		return nil, fmt.Errorf("create request token: %w", err)
	}

	approveURL := fmt.Sprintf(tmdbApproveURLFmt, requestToken.RequestToken)
	printer.Println("Approve tmdb-sync access in your browser:", approveURL)
	OpenBrowser(approveURL)
	printer.Println("Press Enter once you have approved access...")
	WaitForEnter()

	session, _, err := client.Auth.CreateSession(ctx, requestToken.RequestToken)
	if err != nil {
		return nil, fmt.Errorf("create session: %w", err)
	}
	if !session.Valid() {
		return nil, fmt.Errorf("tmdb refused to create a session for the approved request token")
	}

	if err := cfg.WriteSession(fs, config.SessionPath, session); err != nil {
		return nil, fmt.Errorf("persist session: %w", err)
	}

	return session, nil
}
