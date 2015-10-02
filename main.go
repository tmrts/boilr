package main

import (
	"github.com/spf13/cobra"
	"github.com/tmrts/cookie/pkg/cmd"
)

func main() {
	mainCmd := &cobra.Command{
		Use: "main",
	}

	mainCmd.AddCommand(cmd.Build)

	mainCmd.Execute()
}
