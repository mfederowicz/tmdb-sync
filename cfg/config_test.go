package cfg

import (
	"flag"
	"testing"

	"github.com/spf13/afero"
)

func resetFlags() {
	flag.CommandLine = flag.NewFlagSet("test", flag.ContinueOnError)
}

func TestReadConfigFromFile(t *testing.T) {
	fs := afero.NewMemMapFs()
	toml := `
api_key = "abc123"
session_path = "/home/test/session.json"
per_page = 20
pages_limit = 5
`
	if err := afero.WriteFile(fs, "/config.toml", []byte(toml), 0o644); err != nil {
		t.Fatalf("setup: %v", err)
	}

	config, err := ReadConfigFromFile(fs, "/config.toml")
	if err != nil {
		t.Fatalf("ReadConfigFromFile() error = %v", err)
	}
	if config.APIKey != "abc123" {
		t.Errorf("APIKey = %q, want %q", config.APIKey, "abc123")
	}
	if config.PerPage != 20 {
		t.Errorf("PerPage = %d, want 20", config.PerPage)
	}
}

func TestReadConfigFromFile_Missing(t *testing.T) {
	fs := afero.NewMemMapFs()
	if _, err := ReadConfigFromFile(fs, "/nope.toml"); err == nil {
		t.Fatal("expected error for missing file")
	}
}

func TestReadConfigFromFile_Empty(t *testing.T) {
	fs := afero.NewMemMapFs()
	if err := afero.WriteFile(fs, "/empty.toml", []byte(""), 0o644); err != nil {
		t.Fatalf("setup: %v", err)
	}
	if _, err := ReadConfigFromFile(fs, "/empty.toml"); err == nil {
		t.Fatal("expected error for empty file")
	}
}

func TestMergeConfigs_FileOverridesDefault(t *testing.T) {
	resetFlags()
	defaultConfig := DefaultConfig()
	fileConfig := &Config{
		APIKey:      "from-file",
		SessionPath: "/home/test/session.json",
		PerPage:     50,
	}

	merged, err := MergeConfigs(defaultConfig, fileConfig, map[string]string{})
	if err != nil {
		t.Fatalf("MergeConfigs() error = %v", err)
	}
	if merged.APIKey != "from-file" {
		t.Errorf("APIKey = %q, want %q", merged.APIKey, "from-file")
	}
	if merged.PerPage != 50 {
		t.Errorf("PerPage = %d, want 50", merged.PerPage)
	}
	if merged.AuthVersion != "v3" {
		t.Errorf("AuthVersion = %q, want v3 (default preserved)", merged.AuthVersion)
	}
}

func TestMergeConfigs_MissingCredentials(t *testing.T) {
	resetFlags()
	defaultConfig := DefaultConfig()
	fileConfig := &Config{SessionPath: "/home/test/session.json"}

	if _, err := MergeConfigs(defaultConfig, fileConfig, map[string]string{}); err == nil {
		t.Fatal("expected error when neither api_key nor read_access_token is set")
	}
}

func TestMergeConfigs_BadSessionPath(t *testing.T) {
	resetFlags()
	defaultConfig := DefaultConfig()
	fileConfig := &Config{APIKey: "abc", SessionPath: "/home/test/session.txt"}

	if _, err := MergeConfigs(defaultConfig, fileConfig, map[string]string{}); err == nil {
		t.Fatal("expected error for non-json session_path")
	}
}
