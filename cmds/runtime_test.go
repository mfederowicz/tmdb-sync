package cmds

import (
	"bytes"
	"flag"
	"testing"

	"github.com/mfederowicz/tmdb-sync/cfg"
	"github.com/mfederowicz/tmdb-sync/internal"
	"github.com/mfederowicz/tmdb-sync/printer"
	"github.com/mfederowicz/tmdb-sync/str"

	"github.com/spf13/afero"
)

func captureOutput(f func()) string {
	var buf bytes.Buffer
	old := printer.Stdout
	printer.Stdout = &buf
	defer func() { printer.Stdout = old }()
	f()
	return buf.String()
}

func noopExec(afero.Fs, *internal.Client, *cfg.Config, *str.Options, []string) error {
	return nil
}

func TestModulesRuntime_UnknownCommand(t *testing.T) {
	old := flag.Usage
	flag.Usage = func() {}
	defer func() { flag.Usage = old }()

	Commands = []*Command{
		{Name: "movies", Abbrev: "m", Exec: noopExec},
	}

	out := captureOutput(func() {
		ModulesRuntime([]string{"bogus"}, afero.NewMemMapFs(), nil, nil, nil)
	})
	if out == "" {
		t.Fatal("expected an 'unknown command' message to be printed")
	}
}

func TestModulesRuntime_RejectsPartialPrefix(t *testing.T) {
	old := flag.Usage
	flag.Usage = func() {}
	defer func() { flag.Usage = old }()

	Commands = []*Command{
		{Name: "movies", Abbrev: "m", Exec: noopExec},
	}

	out := captureOutput(func() {
		ModulesRuntime([]string{"movi"}, afero.NewMemMapFs(), nil, nil, nil)
	})
	if out == "" {
		t.Fatal("expected an 'unknown command' message to be printed for a partial prefix")
	}
}

func TestModulesRuntime_ResolvesByExactName(t *testing.T) {
	var got []string
	Commands = []*Command{
		{Name: "movies", Abbrev: "m", Exec: func(_ afero.Fs, _ *internal.Client, _ *cfg.Config, _ *str.Options, args []string) error {
			got = args
			return nil
		}},
	}

	ModulesRuntime([]string{"movies", "-a", "popular"}, afero.NewMemMapFs(), nil, nil, nil)
	if len(got) != 2 || got[0] != "-a" || got[1] != "popular" {
		t.Errorf("args passed to Exec = %v, want [-a popular]", got)
	}
}

func TestModulesRuntime_ResolvesByAbbrev(t *testing.T) {
	var got []string
	Commands = []*Command{
		{Name: "movies", Abbrev: "m", Exec: func(_ afero.Fs, _ *internal.Client, _ *cfg.Config, _ *str.Options, args []string) error {
			got = args
			return nil
		}},
	}

	ModulesRuntime([]string{"m", "-a", "popular"}, afero.NewMemMapFs(), nil, nil, nil)
	if len(got) != 2 || got[0] != "-a" || got[1] != "popular" {
		t.Errorf("args passed to Exec = %v, want [-a popular]", got)
	}
}
