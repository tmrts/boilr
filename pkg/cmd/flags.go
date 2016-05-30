package cmd

import cli "github.com/spf13/cobra"

// GetBoolFlag retrieves the named boolean command-line
// flag given the command that contains it.
func GetBoolFlag(c *cli.Command, name string) bool {
	return c.PersistentFlags().Lookup(name).Value.String() == "true"
}

func GetStringFlag(c *cli.Command, name string) string {
	return c.PersistentFlags().Lookup(name).Value.String()
}
