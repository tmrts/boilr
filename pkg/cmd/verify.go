package cmd

import (
	"errors"
	"os"
	"path/filepath"

	cli "github.com/spf13/cobra"
)

var (
	ErrTemplateInvalid = errors.New("verify: given template is invalid")
)

var Verify = &cli.Command{
	Use:   "verify",
	Short: "Verifies whether a template is valid or not",
	Run: func(_ *cli.Command, args []string) {
		templatePath := args[0]

		info, err := os.Stat(templatePath)
		if err != nil {
			panic(err)
		}

		if !info.IsDir() {
			panic(ErrTemplateInvalid)
		}

		templateDirInfo, err := os.Stat(filepath.Join(templatePath, "template"))
		if err != nil {
			if os.IsNotExist(err) {
				panic(ErrTemplateInvalid)

			}

			panic(err)
		}

		if !templateDirInfo.IsDir() {
			panic(ErrTemplateInvalid)
		}
	},
}
