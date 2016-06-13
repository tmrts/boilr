package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	cli "github.com/spf13/cobra"

	"github.com/tmrts/boilr/pkg/boilr"
	"github.com/tmrts/boilr/pkg/template"
	"github.com/tmrts/boilr/pkg/util/exit"
	"github.com/tmrts/boilr/pkg/util/osutil"
	"github.com/tmrts/boilr/pkg/util/validate"
)

// TemplateInRegistry checks whether the given name exists in the template registry.
func TemplateInRegistry(name string) (bool, error) {
	names, err := ListTemplates()
	if err != nil {
		return false, err
	}

	_, ok := names[name]
	return ok, nil
}

// TODO add --use-cache flag to execute a template from previous answers to prompts
// Use contains the cli-command for using templates located in the local template registry.
var Use = &cli.Command{
	Use:   "use <template-tag> <target-dir>",
	Short: "Execute a project template in the given directory",
	Run: func(cmd *cli.Command, args []string) {
		MustValidateArgs(args, []validate.Argument{
			{"template-tag", validate.UnixPath},
			{"target-dir", validate.UnixPath},
		})

		MustValidateTemplateDir()

		tmplName := args[0]
		targetDir, err := filepath.Abs(args[1])
		if err != nil {
			exit.Fatal(fmt.Errorf("use: %s", err))
		}

		templateFound, err := TemplateInRegistry(tmplName)
		if err != nil {
			exit.Fatal(fmt.Errorf("use: %s", err))
		}

		if !templateFound {
			exit.Fatal(fmt.Errorf("Template %q couldn't be found in the template registry", tmplName))
		}

		tmplPath, err := boilr.TemplatePath(tmplName)
		if err != nil {
			exit.Fatal(fmt.Errorf("use: %s", err))
		}

		tmpl, err := template.Get(tmplPath)
		if err != nil {
			exit.Fatal(fmt.Errorf("use: %s", err))
		}

		if shouldUseDefaults := GetBoolFlag(cmd, "use-defaults"); shouldUseDefaults {
			tmpl.UseDefaultValues()
		}

		executeTemplate := func() error {
			parentDir := filepath.Dir(targetDir)

			exists, err := osutil.DirExists(parentDir)
			if err != nil {
				return err
			}

			if !exists {
				return fmt.Errorf("use: parent directory %q doesn't exist", parentDir)
			}

			tmpDir, err := ioutil.TempDir("", "boilr-use-template")
			if err != nil {
				return err
			}
			defer os.RemoveAll(tmpDir)

			if err := tmpl.Execute(tmpDir); err != nil {
				return err
			}

			// Complete the template execution transaction by copying the temporary dir to the target directory.
			return osutil.CopyRecursively(tmpDir, targetDir)
		}

		if err := executeTemplate(); err != nil {
			exit.Fatal(fmt.Errorf("use: %v", err))
		}

		exit.OK("Successfully executed the project template %v in %v", tmplName, targetDir)
	},
}
