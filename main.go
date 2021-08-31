package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

var version = "unknown"

func main() {
	app := &cli.App{
		Name:            "artifactory",
		Usage:           "a plugin of drone",
		Copyright:       "Loc Ngo <xuanloc0511@gmail.com>",
		Description:     "The plugin for publishing artifacts to Jfrog Artifactory",
		Version:         version,
		UsageText:       "artifactory [global options]",
		HideHelpCommand: true,
		Action:          run,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "username",
				EnvVars: []string{"PLUGIN_USERNAME"},
			},
			&cli.StringFlag{
				Name:    "password",
				EnvVars: []string{"PLUGIN_PASSWORD"},
			},
			&cli.StringFlag{
				Name:    "url",
				EnvVars: []string{"PLUGIN_URL"},
			},
			&cli.StringSliceFlag{
				Name:    "poms",
				EnvVars: []string{"PLUGIN_POMS"},
			},
			&cli.StringSliceFlag{
				Name:    "files",
				EnvVars: []string{"./**/target/*.jar"},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func run(ctx *cli.Context) error {
	return nil
}
