package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	cli "github.com/spf13/cobra"

	"github.com/tmrts/boilr/pkg/boilr"
	"github.com/tmrts/boilr/pkg/util/exec"
	"github.com/tmrts/boilr/pkg/util/exit"
	"github.com/tmrts/boilr/pkg/util/osutil"
	"github.com/tmrts/boilr/pkg/util/validate"
)

var Save = &cli.Command{
	// TODO rename template-name to template-tag
	Use:   "save <template-path> <template-name>",
	Short: "Save a project template to local template registry",
	Run: func(c *cli.Command, args []string) {
		MustValidateArgs(args, []validate.Argument{
			{"template-path", validate.UnixPath},
			{"template-name", validate.Alphanumeric},
		})

		tmplDir, templateName := args[0], args[1]

		MustValidateTemplate(tmplDir)

		targetDir := filepath.Join(boilr.Configuration.TemplateDirPath, templateName)

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

		if err := serializeMetadata(templateName, "local:"+tmplDir, targetDir); err != nil {
			exit.Error(fmt.Errorf("save: %s", err))
		}

		exit.OK("Successfully saved the template %v", templateName)
	},
}
