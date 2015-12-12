package cmd

import (
	"fmt"

	cli "github.com/spf13/cobra"

	"github.com/tmrts/tmplt/pkg/tmplt"
)

var Version = &cli.Command{
	Use:   "version",
	Short: "Version information for tmplt",
	Run: func(_ *cli.Command, args []string) {
		fmt.Println("Current version is", tmplt.Version)
	},
}
