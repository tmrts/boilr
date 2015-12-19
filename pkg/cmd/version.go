package cmd

import (
	"fmt"

	cli "github.com/spf13/cobra"

	"github.com/tmrts/boilr/pkg/boilr"
	"github.com/tmrts/boilr/pkg/util/tlog"
	"github.com/tmrts/boilr/pkg/util/validate"
)

var Version = &cli.Command{
	Use:   "version",
	Short: "Show the boilr version information",
	Run: func(_ *cli.Command, args []string) {
		MustValidateArgs(args, []validate.Argument{})

		tlog.Info(fmt.Sprint("Current version is ", boilr.Version))
	},
}
