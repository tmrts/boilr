package cmd

import cli "github.com/spf13/cobra"

var Root = &cli.Command{
	Use: "tmplt",
}

func Run() {
	Init.PersistentFlags().BoolP("force", "f", false, "Recreate directories if they exist")
	Root.AddCommand(Init)

	Root.AddCommand(Use)

	Save.PersistentFlags().BoolP("force", "f", false, "Overwrite existing template with the same name")
	Root.AddCommand(Save)

	Root.AddCommand(Verify)

	Root.AddCommand(Version)

	Root.Execute()
}
