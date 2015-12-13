package cmd

import (
	"fmt"
	"os"

	cli "github.com/spf13/cobra"

	"github.com/tmrts/tmplt/pkg/tmplt"
	"github.com/tmrts/tmplt/pkg/util/exit"
	"github.com/tmrts/tmplt/pkg/util/validate"
)

var List = &cli.Command{
	Use:   "list <template-path> <template-name>",
	Short: "Lists templates found in the template registry",
	Run: func(c *cli.Command, args []string) {
		MustValidateArgs(args, []validate.Argument{})

		d, err := os.Open(tmplt.Configuration.TemplateDirPath)
		if err != nil {
			exit.Error(fmt.Errorf("list: %s", err))
		} else {
			defer d.Close()
		}

		templateNames, err := d.Readdirnames(-1)
		if err != nil {
			exit.Error(fmt.Errorf("list: %s", err))
		}

		for _, name := range templateNames {
			fmt.Println(name)
		}
	},
}
