package main

import (
	cli "github.com/spf13/cobra"

	"github.com/tmrts/tmplt/pkg/cmd"
)

func main() {
	mainCmd := &cli.Command{
		Use: "tmplt",
	}

	mainCmd.AddCommand(cmd.Use)
	mainCmd.AddCommand(cmd.Save)
	mainCmd.AddCommand(cmd.Verify)
	mainCmd.AddCommand(cmd.Version)

	mainCmd.Execute()
}
