package cmd

import (
	"fmt"
	"os"

	cli "github.com/spf13/cobra"

	"github.com/tmrts/boilr/pkg/boilr"
	"github.com/tmrts/boilr/pkg/template"
	"github.com/tmrts/boilr/pkg/util/exit"
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
	Use:   "use <template-name> <target-dir>",
	Short: "Execute a project template in the given directory",
	Run: func(c *cli.Command, args []string) {
		MustValidateArgs(args, []validate.Argument{
			{"template-name", validate.UnixPath},
			{"target-dir", validate.UnixPath},
		})

		tmplName, targetDir := args[0], args[1]

		if ok, err := TemplateInRegistry(tmplName); err != nil {
			exit.Fatal(fmt.Errorf("use: %s", err))
		} else if !ok {
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

		if shouldUseDefaults := GetBoolFlag(c, "use-defaults"); shouldUseDefaults {
			tmpl.UseDefaultValues()
		}

		if err := tmpl.Execute(targetDir); err != nil {
			// Deletes the target dir if execute transaction fails
			defer os.RemoveAll(targetDir)

			exit.Fatal(fmt.Errorf("use: %s", err))
		}

		exit.OK("Successfully executed the project template %v in %v", tmplName, targetDir)
	},
}
