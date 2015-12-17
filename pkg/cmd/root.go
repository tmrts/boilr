package cmd

import cli "github.com/spf13/cobra"

var Root = &cli.Command{
	Use: "tmplt",
}

func Run() {
	// TODO trap c-c to rollback transactions
	// TODO use command factories instead of global command variables
	Init.PersistentFlags().BoolP("force", "f", false, "Recreate directories if they exist")
	Root.AddCommand(Init)

	Root.AddCommand(Use)

	Save.PersistentFlags().BoolP("force", "f", false, "Overwrite existing template with the same name")
	Root.AddCommand(Save)

	// Update Flag
	Root.AddCommand(Download)

	Root.AddCommand(Delete)

	Root.AddCommand(List)

	Root.AddCommand(Validate)

	Root.AddCommand(Version)

	Root.AddCommand(Report)

	Root.Execute()
}
