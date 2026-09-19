package cfg

import (
	"flag"
	"path/filepath"
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

// mergeFromTOML runs a config file body through ReadConfigFromFile and MergeConfigs
// like InitConfig does, so tests see whether a key was actually present.
func mergeFromTOML(t *testing.T, body string) (*Config, error) {
	t.Helper()
	resetFlags()
	fs := afero.NewMemMapFs()
	if err := afero.WriteFile(fs, "/cfg.toml", []byte(body), 0o600); err != nil {
		t.Fatal(err)
	}
	fileConfig, err := ReadConfigFromFile(fs, "/cfg.toml")
	if err != nil {
		t.Fatalf("ReadConfigFromFile() error = %v", err)
	}
	return MergeConfigs(DefaultConfig(), fileConfig, map[string]string{})
}

func TestMergeConfigs_PagesLimit(t *testing.T) {
	base := "api_key = \"k\"\nsession_path = \"/s.json\"\n"
	tests := []struct {
		name    string
		extra   string
		want    int
		wantErr bool
	}{
		{name: "absent keeps default", extra: "", want: 10},
		{name: "zero means all pages", extra: "pages_limit = 0\n", want: 0},
		{name: "positive overrides", extra: "pages_limit = 3\n", want: 3},
		{name: "negative is rejected", extra: "pages_limit = -1\n", wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			merged, err := mergeFromTOML(t, base+tt.extra)
			if tt.wantErr {
				if err == nil {
					t.Fatal("expected an error")
				}
				return
			}
			if err != nil {
				t.Fatalf("MergeConfigs() error = %v", err)
			}
			if merged.PagesLimit != tt.want {
				t.Errorf("PagesLimit = %d, want %d", merged.PagesLimit, tt.want)
			}
		})
	}
}

func TestMergeConfigs_OutputDirExpandsTilde(t *testing.T) {
	home, err := expandTilde("~")
	if err != nil {
		t.Fatal(err)
	}

	merged, err := mergeFromTOML(t, "api_key = \"k\"\nsession_path = \"/s.json\"\noutput_dir = \"~/tmdb\"\n")
	if err != nil {
		t.Fatalf("MergeConfigs() error = %v", err)
	}
	if want := filepath.Join(home, "tmdb"); merged.OutputDir != want {
		t.Errorf("OutputDir = %q, want %q", merged.OutputDir, want)
	}
}

func TestExpandTilde(t *testing.T) {
	home, err := expandTilde("~")
	if err != nil {
		t.Fatal(err)
	}

	tests := []struct{ in, want string }{
		{"~", home},
		{"~/x/y.json", filepath.Join(home, "x/y.json")},
		{"~other/x", "~other/x"}, // another user's home is not ours to guess
		{"/abs/path", "/abs/path"},
		{"rel/~/path", "rel/~/path"},
		{"", ""},
	}
	for _, tt := range tests {
		got, err := expandTilde(tt.in)
		if err != nil {
			t.Fatalf("expandTilde(%q) error = %v", tt.in, err)
		}
		if got != tt.want {
			t.Errorf("expandTilde(%q) = %q, want %q", tt.in, got, tt.want)
		}
	}
}
