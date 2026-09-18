// Package main github.com/mfederowicz/tmdb-sync.
package main

import (
	"context"
	"flag"
	"os"

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
	_debugShort = flag.Bool("d", false, consts.DebugUsage)
	_debug      = flag.Bool("debug", false, consts.DebugUsage)
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
		os.Exit(1)
	}

	options, err := cfg.OptionsFromConfig(fs, config)
	if err != nil {
		printer.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	client := internal.NewClient(nil)
	client.UpdateHeaders(options.Headers)
	client.Debug = options.Debug

	if _, _, err := client.Auth.ValidateKey(context.Background()); err != nil {
		printer.Printf("Error: invalid api_key/read_access_token: %v\n", err)
		os.Exit(1)
	}

	if !cmds.ModulesRuntime(args, fs, config, client, options) {
		os.Exit(1)
	}
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
