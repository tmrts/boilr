package cmd

import cli "github.com/spf13/cobra"

func GetBoolFlag(c *cli.Command, name string) bool {
	return c.PersistentFlags().Lookup(name).Value.String() == "true"
}
