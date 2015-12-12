package cmd

import (
	"errors"
	"os"
	"path/filepath"

	cli "github.com/spf13/cobra"
	"github.com/tmrts/tmplt/pkg/util/exit"
	"github.com/tmrts/tmplt/pkg/util/validate"
)

var (
	ErrTemplateInvalid = errors.New("verify: given template is invalid")
)

var Verify = &cli.Command{
	Use:   "verify",
	Short: "Verifies whether a template is valid or not",
	Run: func(_ *cli.Command, args []string) {
		MustValidateArgs(args, []validate.String{
			validate.UnixPath,
		})

		templatePath := args[0]

		info, err := os.Stat(templatePath)
		if err != nil {
			panic(err)
		}

		if !info.IsDir() {
			exit.Error(ErrTemplateInvalid)
		}

		templateDirInfo, err := os.Stat(filepath.Join(templatePath, "template"))
		if err != nil {
			if os.IsNotExist(err) {
				exit.Error(ErrTemplateInvalid)

			}

			panic(err)
		}

		if !templateDirInfo.IsDir() {
			exit.Error(ErrTemplateInvalid)
		}
	},
}
