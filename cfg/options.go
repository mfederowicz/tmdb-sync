package cfg

import (
	"encoding/json"
	"fmt"

	"github.com/mfederowicz/tmdb-sync/consts"
	"github.com/mfederowicz/tmdb-sync/internal"
	"github.com/mfederowicz/tmdb-sync/str"

	"github.com/spf13/afero"
)

// OptionsFromConfig builds runtime str.Options from a loaded *Config,
// including request headers and any previously persisted session.
func OptionsFromConfig(fs afero.Fs, config *Config) (*str.Options, error) {
	headers := map[string]any{}
	if len(config.ReadAccessToken) > 0 {
		headers[consts.HeaderAuthorization] = fmt.Sprintf("Bearer %s", config.ReadAccessToken)
	} else if len(config.APIKey) > 0 {
		headers[internal.APIKeyParam] = config.APIKey
	}

	options := &str.Options{
		Headers: headers,
		Verbose: config.Verbose,
		Debug:   config.Debug,
	}

	session, err := readSession(fs, config.SessionPath)
	if err == nil {
		options.Session = session
	}

	account, err := readAccount(fs, config.AccountPath)
	if err == nil {
		options.Account = account
	}

	guestSession, err := readGuestSession(fs, config.GuestSessionPath)
	if err == nil {
		options.GuestSession = guestSession
	}

	accessToken, err := readAccessToken(fs, config.AccessTokenPath)
	if err == nil {
		options.AccessTokenV4 = accessToken
	}

	return options, nil
}

func readSession(fs afero.Fs, path string) (*str.Session, error) {
	data, err := afero.ReadFile(fs, path)
	if err != nil {
		return nil, err
	}

	var session str.Session
	if err := json.Unmarshal(data, &session); err != nil {
		return nil, err
	}

	return &session, nil
}

// WriteSession persists a session to disk at path, using afero for testability
// but os.WriteFile-equivalent permissions.
func WriteSession(fs afero.Fs, path string, session *str.Session) error {
	data, err := json.Marshal(session)
	if err != nil {
		return err
	}
	return afero.WriteFile(fs, path, data, consts.X644)
}

func readAccount(fs afero.Fs, path string) (*str.Account, error) {
	data, err := afero.ReadFile(fs, path)
	if err != nil {
		return nil, err
	}

	var account str.Account
	if err := json.Unmarshal(data, &account); err != nil {
		return nil, err
	}

	return &account, nil
}

// WriteAccount caches a resolved account's details to disk at path, so later
// invocations of other account actions can reuse its id without a network
// round trip or repeating -i.
func WriteAccount(fs afero.Fs, path string, account *str.Account) error {
	data, err := json.Marshal(account)
	if err != nil {
		return err
	}
	return afero.WriteFile(fs, path, data, consts.X644)
}

func readGuestSession(fs afero.Fs, path string) (*str.GuestSession, error) {
	data, err := afero.ReadFile(fs, path)
	if err != nil {
		return nil, err
	}

	var guestSession str.GuestSession
	if err := json.Unmarshal(data, &guestSession); err != nil {
		return nil, err
	}

	return &guestSession, nil
}

// WriteGuestSession caches a created guest session to disk at path, so later
// invocations of guest-sessions rated-* actions can reuse its id without
// repeating -i.
func WriteGuestSession(fs afero.Fs, path string, guestSession *str.GuestSession) error {
	data, err := json.Marshal(guestSession)
	if err != nil {
		return err
	}
	return afero.WriteFile(fs, path, data, consts.X644)
}

func readAccessToken(fs afero.Fs, path string) (*str.AccessTokenV4, error) {
	data, err := afero.ReadFile(fs, path)
	if err != nil {
		return nil, err
	}

	var accessToken str.AccessTokenV4
	if err := json.Unmarshal(data, &accessToken); err != nil {
		return nil, err
	}

	return &accessToken, nil
}

// WriteAccessToken persists a v4 user access token to disk at path, readable
// by the owner only since the token grants access to the user's account.
func WriteAccessToken(fs afero.Fs, path string, accessToken *str.AccessTokenV4) error {
	data, err := json.Marshal(accessToken)
	if err != nil {
		return err
	}
	return afero.WriteFile(fs, path, data, consts.X600)
}
