package cmd

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	cli "github.com/spf13/cobra"
	"github.com/tmrts/boilr/pkg/boilr"
	"github.com/tmrts/boilr/pkg/util/exit"
)

const bashrcText = `
# Enables shell command completion for boilr
source $HOME/bin/boilr
`

func configureBashCompletion() error {
	bash_completion_file := filepath.Join(boilr.Configuration.ConfigDirPath, "completion.bash")

	if err := Root.GenBashCompletionFile(bash_completion_file); err != nil {
		return err
	}

	if err := Root.GenBashCompletionFile(bash_completion_file); err != nil {
		return err
	}

	bashrcPath := filepath.Join(os.Getenv("HOME"), ".bashrc")
	if bashrcPath == "" {
		return errors.New("environment variable ${HOME} should be set")
	}

	f, err := os.OpenFile(bashrcPath, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}

	defer f.Close()

	if _, err = f.WriteString(bashrcText); err != nil {
		return err
	}

	return nil
}

// ConfigureBashCompletion generates bash auto-completion script and installs it.
var ConfigureBashCompletion = &cli.Command{
	Hidden: true,
	Use:    "configure-bash-completion",
	Short:  "Configure bash the auto-completion",
	Run: func(c *cli.Command, _ []string) {
		if err := configureBashCompletion(); err != nil {
			exit.Fatal(fmt.Errorf("configure-bash-completion: %s", err))
		}

		exit.OK("Successfully configured bash auto-completion")
	},
}
