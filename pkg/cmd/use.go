package cmd

import (
	"fmt"
	"os"

	cli "github.com/spf13/cobra"

	"github.com/tmrts/tmplt/pkg/template"
	"github.com/tmrts/tmplt/pkg/tmplt"
	"github.com/tmrts/tmplt/pkg/util/exit"
	"github.com/tmrts/tmplt/pkg/util/validate"
)

func TemplateInRegistry(name string) (bool, error) {
	names, err := ListTemplates()
	if err != nil {
		return false, err
	}

	_, ok := names[name]
	return ok, nil
}

// TODO add --use-defaults flag to execute a template without user prompts
// TODO add --use-cache flag to execute a template from previous answers to prompts
var Use = &cli.Command{
	Use:   "use <template-name> <target-dir>",
	Short: "Executes a project template",
	Run: func(_ *cli.Command, args []string) {
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

		tmplPath, err := tmplt.TemplatePath(tmplName)
		if err != nil {
			exit.Fatal(fmt.Errorf("use: %s", err))
		}

		tmpl, err := template.Get(tmplPath)
		if err != nil {
			exit.Fatal(fmt.Errorf("use: %s", err))
		}

		if err := tmpl.Execute(targetDir); err != nil {
			// Deletes the target dir if execute transaction fails
			defer os.RemoveAll(targetDir)

			exit.Fatal(fmt.Errorf("use: %s", err))
		}

		exit.OK("Successfully executed the project template %v in %v", tmplName, targetDir)
	},
}
