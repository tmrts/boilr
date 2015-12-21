package cmd

import cli "github.com/spf13/cobra"

// Root contains the root cli-command containing all the other commands.
var Root = &cli.Command{
	Use: "boilr",
}

// Run executes the cli-command root.
func Run() {
	// TODO trap c-c to rollback transactions
	// TODO use command factories instead of global command variables
	// TODO add describe command that shows template metadata information
	// TODO add create command that creates a minimal template template
	// TODO rename template-name to template-tag
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
