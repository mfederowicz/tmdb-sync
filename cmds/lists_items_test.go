package cmds

import (
	"testing"

	"github.com/mfederowicz/tmdb-sync/str"
)

func TestItemFlagSet(t *testing.T) {
	var f itemFlag
	for _, v := range []string{"movie:100", "tv:200:nice: really"} {
		if err := f.Set(v); err != nil {
			t.Fatalf("Set(%q) error = %v", v, err)
		}
	}
	want := []str.ListMediaV4{{MediaType: "movie", MediaID: 100}, {MediaType: "tv", MediaID: 200, Comment: "nice: really"}}
	if len(f) != 2 || f[0] != want[0] || f[1] != want[1] {
		t.Errorf("items = %+v, want %+v", f, want)
	}

	for _, bad := range []string{"", "movie", "person:1", "movie:x", "movie:0"} {
		if err := (&itemFlag{}).Set(bad); err == nil {
			t.Errorf("Set(%q) error = nil, want error", bad)
		}
	}
}
