package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	cli "github.com/spf13/cobra"
	"github.com/tmrts/boilr/pkg/boilr"
	"github.com/tmrts/boilr/pkg/util/exit"
)

func configureBashCompletion() error {
	var bash_completion_dir string
	if bash_completion_dir = os.Getenv("BASH_COMPLETION_COMPAT_DIR"); bash_completion_dir == "" {
		bash_completion_dir = "/etc/bash_completion.d"
	}

	bash_completion_file := filepath.Join(bash_completion_dir, boilr.AppName+".bash")

	if err := Root.GenBashCompletionFile(bash_completion_file); err != nil {
		if strings.Contains(err.Error(), "permission") {
			return fmt.Errorf("couldn't configure bash completion for %s: permission denied", boilr.AppName)
		}

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
