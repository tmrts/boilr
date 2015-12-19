package cmd

import (
	"errors"

	cli "github.com/spf13/cobra"
	"github.com/tmrts/boilr/pkg/util/exit"
	"github.com/tmrts/boilr/pkg/util/validate"
)

var (
	ErrTemplateInvalid = errors.New("validate: given template is invalid")
)

var Validate = &cli.Command{
	Use:   "validate",
	Short: "Validates whether a template is valid or not",
	Run: func(_ *cli.Command, args []string) {
		MustValidateArgs(args, []validate.Argument{
			{"template-path", validate.UnixPath},
		})

		templatePath := args[0]

		MustValidateTemplate(templatePath)

		exit.OK("Template is valid")
	},
}
