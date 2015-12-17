package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	cli "github.com/spf13/cobra"

	"github.com/tmrts/tmplt/pkg/tmplt"
	"github.com/tmrts/tmplt/pkg/util/exec"
	"github.com/tmrts/tmplt/pkg/util/exit"
	"github.com/tmrts/tmplt/pkg/util/osutil"
	"github.com/tmrts/tmplt/pkg/util/validate"
)

var Save = &cli.Command{
	Use:   "save <template-path> <template-name>",
	Short: "Save a project template to local template registry",
	Run: func(c *cli.Command, args []string) {
		MustValidateArgs(args, []validate.Argument{
			{"template-path", validate.UnixPath},
			{"template-name", validate.Alphanumeric},
		})

		tmplDir, templateName := args[0], args[1]

		MustValidateTemplate(tmplDir)

		targetDir := filepath.Join(tmplt.Configuration.TemplateDirPath, templateName)

		switch exists, err := osutil.DirExists(targetDir); {
		case err != nil:
			exit.Error(fmt.Errorf("save: %s", err))
		case exists:
			shouldOverwrite := GetBoolFlag(c, "force")

			if err != nil {
				exit.Error(fmt.Errorf("save: %v", err))
			}

			if !shouldOverwrite {
				exit.OK("Template %v already exists use -f to overwrite", templateName)
			}

			if err := os.RemoveAll(targetDir); err != nil {
				exit.Error(fmt.Errorf("save: %v", err))
			}
		}

		if _, err := exec.Cmd("cp", "-r", tmplDir, targetDir); err != nil {
			// TODO create exec package
			exit.Error(err)
		}

		exit.OK("Successfully saved the template %v", templateName)
	},
}
