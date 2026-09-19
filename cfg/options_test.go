package cfg

import (
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
