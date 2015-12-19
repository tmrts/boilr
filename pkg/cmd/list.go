package cmd

import (
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
	cli "github.com/spf13/cobra"

	"github.com/tmrts/tmplt/pkg/template"
	"github.com/tmrts/tmplt/pkg/tmplt"
	"github.com/tmrts/tmplt/pkg/util/exit"
	"github.com/tmrts/tmplt/pkg/util/validate"
)

func ListTemplates() (map[string]bool, error) {
	d, err := os.Open(tmplt.Configuration.TemplateDirPath)
	if err != nil {
		return nil, err
	} else {
		defer d.Close()
	}

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
		for name, _ := range templateNames {
			tmplPath, err := tmplt.TemplatePath(name)
			if err != nil {
				exit.Fatal(fmt.Errorf("list: %s", err))
			}

			tmpl, err := template.Get(tmplPath)
			if err != nil {
				exit.Fatal(fmt.Errorf("list: %s", err))
			}

			data = append(data, tmpl.Info().String())
		}

		// TODO Wrap in a util function
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Tag", "Repository", "Created"})

		for _, datum := range data {
			table.Append(datum)
		}

		if len(data) == 0 {
			table.Append([]string{"", "", ""})
		}

		table.Render()
	},
}
