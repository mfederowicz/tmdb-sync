package cfg

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/mfederowicz/tmdb-sync/consts"
	"github.com/mfederowicz/tmdb-sync/str"

	"github.com/spf13/afero"
)

func TestAccessTokenRoundTrip(t *testing.T) {
	fs := afero.NewMemMapFs()
	want := &str.AccessTokenV4{Success: true, AccessToken: "tok", AccountID: "acc"}

	if err := WriteAccessToken(fs, "/tok.json", want); err != nil {
		t.Fatalf("WriteAccessToken() error = %v", err)
	}
	info, err := fs.Stat("/tok.json")
	if err != nil {
		t.Fatalf("stat: %v", err)
	}
	if info.Mode().Perm() != consts.X600 {
		t.Errorf("perm = %o, want %o", info.Mode().Perm(), consts.X600)
	}

	options, err := OptionsFromConfig(fs, &Config{AccessTokenPath: "/tok.json"})
	if err != nil {
		t.Fatalf("OptionsFromConfig() error = %v", err)
	}
	if options.AccessTokenV4 == nil || *options.AccessTokenV4 != *want {
		t.Errorf("AccessTokenV4 = %+v, want %+v", options.AccessTokenV4, want)
	}
}

func TestOptionsFromConfig_NoAccessToken(t *testing.T) {
	options, err := OptionsFromConfig(afero.NewMemMapFs(), &Config{AccessTokenPath: "/missing.json"})
	if err != nil {
		t.Fatalf("OptionsFromConfig() error = %v", err)
	}
	if options.AccessTokenV4 != nil {
		t.Errorf("AccessTokenV4 = %+v, want nil", options.AccessTokenV4)
	}
}

func TestWriters_CreateParentDirAndPrivateMode(t *testing.T) {
	writers := map[string]func(fs afero.Fs, path string) error{
		"session": func(fs afero.Fs, path string) error {
			return WriteSession(fs, path, &str.Session{SessionID: "s"})
		},
		"account": func(fs afero.Fs, path string) error {
			return WriteAccount(fs, path, &str.Account{ID: 1})
		},
		"guest session": func(fs afero.Fs, path string) error {
			return WriteGuestSession(fs, path, &str.GuestSession{GuestSessionID: "g"})
		},
		"access token": func(fs afero.Fs, path string) error {
			return WriteAccessToken(fs, path, &str.AccessTokenV4{AccessToken: "t"})
		},
	}

	for name, write := range writers {
		t.Run(name, func(t *testing.T) {
			// A real filesystem: MemMapFs creates parents implicitly and would hide the bug.
			fs := afero.NewOsFs()
			path := filepath.Join(t.TempDir(), "missing", "config", "file.json")

			if err := write(fs, path); err != nil {
				t.Fatalf("write into a missing directory: %v", err)
			}
			info, err := fs.Stat(path)
			if err != nil {
				t.Fatalf("stat: %v", err)
			}
			if info.Mode().Perm() != consts.X600 {
				t.Errorf("perm = %o, want %o", info.Mode().Perm(), consts.X600)
			}
		})
	}
}

func TestWriters_TightenExistingFileMode(t *testing.T) {
	fs := afero.NewOsFs()
	path := filepath.Join(t.TempDir(), "session.json")
	if err := os.WriteFile(path, []byte("{}"), 0o644); err != nil {
		t.Fatal(err)
	}

	if err := WriteSession(fs, path, &str.Session{SessionID: "s"}); err != nil {
		t.Fatalf("WriteSession() error = %v", err)
	}
	info, err := os.Stat(path)
	if err != nil {
		t.Fatal(err)
	}
	if info.Mode().Perm() != consts.X600 {
		t.Errorf("perm = %o, want %o after rewriting a 0644 file", info.Mode().Perm(), consts.X600)
	}
}
