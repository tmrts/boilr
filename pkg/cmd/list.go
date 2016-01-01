package cmd

import (
	"fmt"
	"os"

	cli "github.com/spf13/cobra"

	"github.com/tmrts/boilr/pkg/boilr"
	"github.com/tmrts/boilr/pkg/template"
	"github.com/tmrts/boilr/pkg/util/exit"
	"github.com/tmrts/boilr/pkg/util/tabular"
	"github.com/tmrts/boilr/pkg/util/validate"
)

// ListTemplates returns a list of templates saved in the local template registry.
func ListTemplates() (map[string]bool, error) {
	d, err := os.Open(boilr.Configuration.TemplateDirPath)
	if err != nil {
		return nil, err
	}
	defer d.Close()

	names, err := d.Readdirnames(-1)
	if err != nil {
		return nil, err
	}

	nameSet := make(map[string]bool)
	for _, name := range names {
		nameSet[name] = true
	}

	return nameSet, nil
}

// List contains the cli-command for printing a list of saved templates.
var List = &cli.Command{
	Use:   "list <template-path> <template-name>",
	Short: "List project templates found in the local template registry",
	Run: func(c *cli.Command, args []string) {
		MustValidateArgs(args, []validate.Argument{})

		templateNames, err := ListTemplates()
		if err != nil {
			exit.Error(fmt.Errorf("list: %s", err))
		}

		var data [][]string
		for name := range templateNames {
			tmplPath, err := boilr.TemplatePath(name)
			if err != nil {
				exit.Fatal(fmt.Errorf("list: %s", err))
			}

			tmpl, err := template.Get(tmplPath)
			if err != nil {
				exit.Fatal(fmt.Errorf("list: %s", err))
			}

			data = append(data, tmpl.Info().String())
		}

		tabular.Print([]string{"Tag", "Repository", "Created"}, data)
	},
}
