package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	cli "github.com/spf13/cobra"

	"github.com/tmrts/tmplt/pkg/tmplt"
	"github.com/tmrts/tmplt/pkg/util/exit"
	"github.com/tmrts/tmplt/pkg/util/inpututil"
	"github.com/tmrts/tmplt/pkg/util/osutil"
)

var Save = &cli.Command{
	Use:   "save",
	Short: "Saves a project template to template registry",
	Run: func(_ *cli.Command, args []string) {
		templateName, sourceDir := args[0], args[1]

		targetDir := filepath.Join(tmplt.TemplateDirPath, templateName)

		switch err := osutil.FileExists(targetDir); {
		case os.IsNotExist(err):
			break
		case err == nil:
			// Template Already Exists Ask If Should be Replaced
			shouldOverride, err := inpututil.ScanYesOrNo("Template already exists. Override?", false)
			if err != nil {
				exit.Error(fmt.Errorf("save: %v", err))
			}

			if !shouldOverride {
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
