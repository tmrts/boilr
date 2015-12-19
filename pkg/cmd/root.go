package cmd

import cli "github.com/spf13/cobra"

var Root = &cli.Command{
	Use: "boilr",
}

func Run() {
	// TODO add loading bars or progress bars to commands that take time
	// TODO trap c-c to rollback transactions
	// TODO use command factories instead of global command variables
	Init.PersistentFlags().BoolP("force", "f", false, "Recreate directories if they exist")
	Root.AddCommand(Init)

	Use.PersistentFlags().BoolP("use-defaults", "f", false, "Uses default values in project.json instead of prompting the user")
	Root.AddCommand(Use)

	Save.PersistentFlags().BoolP("force", "f", false, "Overwrite existing template with the same name")
	Root.AddCommand(Save)

	Download.PersistentFlags().BoolP("force", "f", false, "Overwrite existing template with the same name")
	Root.AddCommand(Download)

	Root.AddCommand(Delete)

	Root.AddCommand(List)

	Root.AddCommand(Validate)

	Root.AddCommand(Version)

	Root.AddCommand(Report)

	Root.Execute()
}
