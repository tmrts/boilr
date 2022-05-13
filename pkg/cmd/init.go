package cmd

import (
	"fmt"

	"boilr/pkg/boilr"
	"boilr/pkg/util/exit"
	"boilr/pkg/util/osutil"
	cli "github.com/spf13/cobra"
)

// Init contains the cli-command for initializing the local template
// registry in case it's not initialized.
var Init = &cli.Command{
	Use:   "init",
	Short: "Initialize directories required by boilr (By default done by installation script)",
	Run: func(c *cli.Command, _ []string) {
		// Check if .config/boilr exists
		if exists, err := osutil.DirExists(boilr.Configuration.TemplateDirPath); exists {
			if shouldRecreate := GetBoolFlag(c, "force"); !shouldRecreate {
				exit.GoodEnough("template registry is already initialized use -f to reinitialize")
			}
		} else if err != nil {
			exit.Error(fmt.Errorf("init: %s", err))
		}

		if err := osutil.CreateDirs(boilr.Configuration.TemplateDirPath); err != nil {
			exit.Error(err)
		}

		exit.OK("Initialization complete")
	},
}
