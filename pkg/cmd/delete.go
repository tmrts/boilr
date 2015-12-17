package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	cli "github.com/spf13/cobra"

	"github.com/tmrts/tmplt/pkg/tmplt"
	"github.com/tmrts/tmplt/pkg/util/exit"
	"github.com/tmrts/tmplt/pkg/util/osutil"
	"github.com/tmrts/tmplt/pkg/util/validate"
)

var Delete = &cli.Command{
	Use:   "delete <template-name>",
	Short: "Delete a project template from the template registry",
	Run: func(c *cli.Command, args []string) {
		MustValidateArgs(args, []validate.Argument{
			{"template-path", validate.Alphanumeric},
		})

		templateName := args[0]

		targetDir := filepath.Join(tmplt.Configuration.TemplateDirPath, templateName)

		switch exists, err := osutil.DirExists(targetDir); {
		case err != nil:
			exit.Error(fmt.Errorf("delete: %s", err))
		case !exists:
			exit.Error(fmt.Errorf("Template %v doesn't exist", templateName))
		}

		// TODO Accept globs and multiple arguments
		if err := os.RemoveAll(targetDir); err != nil {
			exit.Error(fmt.Errorf("delete: %v", err))
		}

		exit.OK("Successfully deleted the template %v", templateName)
	},
}
