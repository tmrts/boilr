package cmd

import (
	"fmt"

	cli "github.com/spf13/cobra"

	"github.com/tmrts/tmplt/pkg/tmplt"
	"github.com/tmrts/tmplt/pkg/util/validate"
)

var Version = &cli.Command{
	Use:   "version",
	Short: "Show the tmplt version information",
	Run: func(_ *cli.Command, args []string) {
		MustValidateArgs(args, []validate.Argument{})

		fmt.Println("Current version is", tmplt.Version)
	},
}
