// Package cfg used for process configuration
package cfg

import (
	"errors"
	"flag"
	"fmt"
	"os/user"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/mfederowicz/tmdb-sync/consts"

	"github.com/BurntSushi/toml"
	"github.com/spf13/afero"
)

// Config struct for app.
type Config struct {
	APIKey          string `toml:"api_key"`
	ReadAccessToken string `toml:"read_access_token"`
	AuthVersion     string `toml:"auth_version"`
	ConfigPath      string `toml:"config_path"`
	SessionPath     string `toml:"session_path"`
	AccountPath     string `toml:"account_path"`
	OutputDir       string `toml:"output_dir"`
	PerPage         int    `toml:"per_page"`
	PagesLimit      int    `toml:"pages_limit"`
	Verbose         bool   `toml:"verbose"`
}

// InitConfig of app
func InitConfig(fs afero.Fs) (*Config, error) {
	flagMap := map[string]string{}
	flag.VisitAll(func(f *flag.Flag) {
		flagMap[f.Name] = f.Value.String()
	})

	if len(flagMap["c"]) == consts.ZeroValue {
		return nil, errors.New("config file not exists")
	}

	configFromFile, err := ReadConfigFromFile(fs, flagMap["c"])
	if err != nil {
		return nil, fmt.Errorf("init config error : %w", err)
	}

	return MergeConfigs(DefaultConfig(), configFromFile, flagMap)
}

// GenUsedFlagMap map of used flags
func GenUsedFlagMap() map[string]bool {
	flagset := map[string]bool{}

	flag.Visit(func(f *flag.Flag) {
		key := string(f.Name[0])
		flagset[key] = true
	})

	return flagset
}

// MergeConfigs from two sources file and flags
func MergeConfigs(defaultConfig *Config, fileConfig *Config, flagConfig map[string]string) (*Config, error) {
	flagset := GenUsedFlagMap()

	defaultConfig.APIKey = processOptionAPIKey(defaultConfig, fileConfig)
	defaultConfig.ReadAccessToken = processOptionReadAccessToken(defaultConfig, fileConfig)
	defaultConfig.AuthVersion = processOptionAuthVersion(defaultConfig, fileConfig)
	defaultConfig.PerPage = processOptionPerPage(defaultConfig, fileConfig)
	defaultConfig.PagesLimit = processOptionPagesLimit(defaultConfig, fileConfig)
	defaultConfig.OutputDir = processOptionOutputDir(defaultConfig, fileConfig)
	defaultConfig.Verbose = processOptionVerbose(defaultConfig, fileConfig, flagConfig, flagset)

	sessionPath, err := processOptionSessionPath(defaultConfig, fileConfig)
	if err != nil {
		return nil, fmt.Errorf("config error : %w", err)
	}
	defaultConfig.SessionPath = sessionPath

	accountPath, err := processOptionAccountPath(defaultConfig, fileConfig)
	if err != nil {
		return nil, fmt.Errorf("config error : %w", err)
	}
	defaultConfig.AccountPath = accountPath

	defaultConfig.ConfigPath = processOptionConfigPath(defaultConfig, fileConfig, flagConfig, flagset)

	if err := normalizeConfig(defaultConfig); err != nil {
		return nil, fmt.Errorf("config error : %w", err)
	}

	return defaultConfig, nil
}

func processOptionAPIKey(defaultConfig *Config, fileConfig *Config) string {
	if len(fileConfig.APIKey) > consts.ZeroValue {
		defaultConfig.APIKey = fileConfig.APIKey
	}
	return defaultConfig.APIKey
}

func processOptionReadAccessToken(defaultConfig *Config, fileConfig *Config) string {
	if len(fileConfig.ReadAccessToken) > consts.ZeroValue {
		defaultConfig.ReadAccessToken = fileConfig.ReadAccessToken
	}
	return defaultConfig.ReadAccessToken
}

func processOptionAuthVersion(defaultConfig *Config, fileConfig *Config) string {
	if len(fileConfig.AuthVersion) > consts.ZeroValue {
		defaultConfig.AuthVersion = fileConfig.AuthVersion
	}
	return defaultConfig.AuthVersion
}

func processOptionOutputDir(defaultConfig *Config, fileConfig *Config) string {
	if len(fileConfig.OutputDir) > consts.ZeroValue {
		defaultConfig.OutputDir = fileConfig.OutputDir
	}
	return defaultConfig.OutputDir
}

func processOptionPerPage(defaultConfig *Config, fileConfig *Config) int {
	if fileConfig.PerPage > consts.ZeroValue {
		defaultConfig.PerPage = fileConfig.PerPage
	}
	return defaultConfig.PerPage
}

func processOptionPagesLimit(defaultConfig *Config, fileConfig *Config) int {
	if fileConfig.PagesLimit > consts.ZeroValue {
		defaultConfig.PagesLimit = fileConfig.PagesLimit
	}
	return defaultConfig.PagesLimit
}

func processOptionVerbose(defaultConfig *Config, fileConfig *Config, flagConfig map[string]string, flagset map[string]bool) bool {
	if fileConfig.Verbose {
		defaultConfig.Verbose = fileConfig.Verbose
	}
	if flagset["v"] {
		if boolValue, err := strconv.ParseBool(flagConfig["v"]); err == nil {
			defaultConfig.Verbose = boolValue
		}
	}
	return defaultConfig.Verbose
}

func processOptionSessionPath(defaultConfig *Config, fileConfig *Config) (string, error) {
	if len(fileConfig.SessionPath) > consts.ZeroValue {
		defaultConfig.SessionPath = fileConfig.SessionPath
	}

	sessionPath, err := expandTilde(defaultConfig.SessionPath)
	if err != nil {
		return "", fmt.Errorf("failed to expand tilde from sessionPath: %w", err)
	}
	return sessionPath, nil
}

func processOptionAccountPath(defaultConfig *Config, fileConfig *Config) (string, error) {
	if len(fileConfig.AccountPath) > consts.ZeroValue {
		defaultConfig.AccountPath = fileConfig.AccountPath
	}

	accountPath, err := expandTilde(defaultConfig.AccountPath)
	if err != nil {
		return "", fmt.Errorf("failed to expand tilde from accountPath: %w", err)
	}
	return accountPath, nil
}

func processOptionConfigPath(defaultConfig *Config, fileConfig *Config, flagConfig map[string]string, flagset map[string]bool) string {
	if len(fileConfig.ConfigPath) > consts.ZeroValue {
		defaultConfig.ConfigPath = fileConfig.ConfigPath
	}
	if flagset["c"] && flagConfig["c"] != consts.EmptyString {
		defaultConfig.ConfigPath = flagConfig["c"]
	}
	return defaultConfig.ConfigPath
}

// ReadConfigFromFile reads config from file stored on disc
func ReadConfigFromFile(fs afero.Fs, filename string) (*Config, error) {
	var config Config

	file, err := afero.ReadFile(fs, filename)
	if err != nil {
		return nil, fmt.Errorf("cannot read the config file : %w", err)
	}

	if len(string(file)) == consts.ZeroValue {
		return nil, errors.New("empty file content")
	}

	if _, err := toml.Decode(string(file), &config); err != nil {
		return nil, fmt.Errorf("cannot parse the config file : %w", err)
	}

	return &config, nil
}

// DefaultConfig config with default values
func DefaultConfig() *Config {
	return &Config{
		APIKey:          consts.EmptyString,
		ReadAccessToken: consts.EmptyString,
		AuthVersion:     "v3",
		ConfigPath:      buildDefaultConfigPath(),
		SessionPath:     buildDefaultSessionPath(),
		AccountPath:     buildDefaultAccountPath(),
		OutputDir:       consts.EmptyString,
		PerPage:         consts.ZeroValue,
		PagesLimit:      consts.PagesLimit,
		Verbose:         false,
	}
}

func normalizeConfig(config *Config) error {
	if len(config.APIKey) == consts.ZeroValue && len(config.ReadAccessToken) == consts.ZeroValue {
		return errors.New("api_key or read_access_token is required, update your config file")
	}

	if len(config.SessionPath) == consts.ZeroValue || !strings.HasSuffix(config.SessionPath, "json") {
		return errors.New("session_path should be a json file, update your config file")
	}

	return nil
}

func expandTilde(path string) (string, error) {
	if len(path) > consts.ZeroValue && path[consts.ZeroValue] == '~' {
		usr, err := user.Current()
		if err != nil {
			return consts.EmptyString, err
		}
		return filepath.Join(usr.HomeDir, path[1:]), nil
	}
	return path, nil
}

func buildDefaultConfigPath() string {
	absPath, err := expandTilde("~/tmdb-sync.toml")
	if err != nil {
		panic(err)
	}
	return absPath
}

func buildDefaultSessionPath() string {
	absPath, err := expandTilde("~/.config/tmdb-sync/session.json")
	if err != nil {
		panic(err)
	}
	return absPath
}

func buildDefaultAccountPath() string {
	absPath, err := expandTilde("~/.config/tmdb-sync/account.json")
	if err != nil {
		panic(err)
	}
	return absPath
}
