package cmd

import (
	"time"

	cli "github.com/spf13/cobra"

	"github.com/tmrts/tmplt/pkg/cmd/util"
	"github.com/tmrts/tmplt/pkg/template"
	"github.com/tmrts/tmplt/pkg/tmplt"
	"github.com/tmrts/tmplt/pkg/util/exit"
	"github.com/tmrts/tmplt/pkg/util/validate"
)

func MustValidateArgs(args []string, validations []validate.String) {
	if errs := util.ValidateArgs(args, validations); len(errs) > 0 {
		exit.Error(errs...)
	}
}

var Use = &cli.Command{
	Use:   "use <template-name>",
	Short: "Executes a project template",
	Run: func(_ *cli.Command, args []string) {
		MustValidateArgs(args, []validate.String{
			validate.Alphanumeric,
		})

		tmplPath, err := tmplt.TemplatePath(args[0])
		if err != nil {
			panic(err)
		}

		tmpl, err := template.Get(tmplPath)
		if err != nil {
			panic(err)
		}

		metadata := template.Metadata{
			Name:    "test-project-1",
			Author:  "test-author",
			Email:   "test@mail.com",
			Date:    time.Now().Format("Mon Jan 2 2006 15:04:05"),
			Version: "0.0.1",
		}

		err = tmpl.Execute(args[1], metadata)
		if err != nil {
			panic(err)
		}
	},
}
