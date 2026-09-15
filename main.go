// Package main github.com/mfederowicz/tmdb-sync.
package main

import (
	"flag"

	"github.com/mfederowicz/tmdb-sync/cfg"
	"github.com/mfederowicz/tmdb-sync/cli"
	"github.com/mfederowicz/tmdb-sync/cmds"
	"github.com/mfederowicz/tmdb-sync/consts"
	"github.com/mfederowicz/tmdb-sync/internal"
	"github.com/mfederowicz/tmdb-sync/printer"

	"github.com/spf13/afero"
)

var (
	_verbose    = flag.Bool("v", false, consts.VerboseUsage)
	_version    = flag.Bool("version", false, consts.VersionUsage)
	_configPath = flag.String("c", cfg.DefaultConfig().ConfigPath, consts.ConfigUsage)
)

func init() {
	flag.Usage = func() {
		cmds.HelpFunc(nil)
	}
	flag.Parse()
}

func main() {
	args, noflags := handleArgs()
	if noflags {
		return
	}

	fs := afero.NewOsFs()
	config, err := cfg.InitConfig(fs)
	if err != nil {
		printer.Printf("Error: %v\n", err)
		return
	}

	options, err := cfg.OptionsFromConfig(fs, config)
	if err != nil {
		printer.Printf("Error: %v\n", err)
		return
	}

	client := internal.NewClient(nil)
	client.UpdateHeaders(options.Headers)

	cmds.ModulesRuntime(args, fs, config, client, options)
}

func handleArgs() ([]string, bool) {
	if *_version {
		printer.Println(cli.GenAppVersion())
		return []string{}, true
	}

	args := flag.Args()
	if len(args) == consts.ZeroValue {
		flag.Usage()
		return []string{}, true
	}

	return args, false
}
