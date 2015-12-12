package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	cli "github.com/spf13/cobra"

	"github.com/tmrts/tmplt/pkg/tmplt"
	"github.com/tmrts/tmplt/pkg/util/exit"
	"github.com/tmrts/tmplt/pkg/util/osutil"
	"github.com/tmrts/tmplt/pkg/util/validate"
)

var Save = &cli.Command{
	Use:   "save <template-path> <template-name>",
	Short: "Saves a project template to template registry",
	Run: func(c *cli.Command, args []string) {
		MustValidateArgs(args, []validate.String{
			validate.UnixPath,
			validate.Alphanumeric,
		})

		sourceDir, templateName := args[0], args[1]

		targetDir := filepath.Join(tmplt.Configuration.TemplateDirPath, templateName)

		switch err := osutil.FileExists(targetDir); {
		case os.IsNotExist(err):
			break
		case err == nil:
			shouldOverwrite := GetBoolFlag(c, "force")

			if err != nil {
				exit.Error(fmt.Errorf("save: %v", err))
			}

			if !shouldOverwrite {
				exit.OK("Exiting")
			}
		default:
			exit.Error(err)
		}

		if _, err := exec.Command("/usr/bin/cp", "-r", "--remove-destination", sourceDir, targetDir).Output(); err != nil {
			// TODO create exec package
			exit.Error(err)
		}
	},
}
