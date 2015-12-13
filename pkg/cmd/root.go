package cmd

import cli "github.com/spf13/cobra"

var Root = &cli.Command{
	Use: "tmplt",
}

func Run() {
	// TODO use command factories instead of global command variables

	Init.PersistentFlags().BoolP("force", "f", false, "Recreate directories if they exist")
	Root.AddCommand(Init)

	Root.AddCommand(Use)

	Save.PersistentFlags().BoolP("force", "f", false, "Overwrite existing template with the same name")
	Root.AddCommand(Save)

	Root.AddCommand(Validate)

	Root.AddCommand(Version)

	Root.Execute()
}
