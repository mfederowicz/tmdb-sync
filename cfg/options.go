package cfg

import (
	"encoding/json"
	"fmt"

	"github.com/mfederowicz/tmdb-sync/internal"
	"github.com/mfederowicz/tmdb-sync/str"

	"github.com/spf13/afero"
)

// OptionsFromConfig builds runtime str.Options from a loaded *Config,
// including request headers and any previously persisted session.
func OptionsFromConfig(fs afero.Fs, config *Config) (*str.Options, error) {
	headers := map[string]any{}
	if len(config.ReadAccessToken) > 0 {
		headers["Authorization"] = fmt.Sprintf("Bearer %s", config.ReadAccessToken)
	} else if len(config.APIKey) > 0 {
		headers[internal.APIKeyParam] = config.APIKey
	}

	options := &str.Options{
		Headers: headers,
		Verbose: config.Verbose,
	}

	session, err := readSession(fs, config.SessionPath)
	if err == nil {
		options.Session = session
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
	return afero.WriteFile(fs, path, data, 0o644)
}
