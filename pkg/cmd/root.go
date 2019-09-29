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

	Template := &cli.Command{
		Use:   "template",
		Short: "Run a template command",
	}

	Template.AddCommand(Delete)

	Download.PersistentFlags().BoolP("force", "f", false, "Overwrite existing template with the same name")
	Download.PersistentFlags().StringP("log-level", "l", "error", "log-level for output")
	Template.AddCommand(Download)

	List.PersistentFlags().BoolP("dont-prettify", "", false, "Print only the template names without fancy formatting")
	Template.AddCommand(List)

	Template.AddCommand(Rename)

	Save.PersistentFlags().BoolP("force", "f", false, "Overwrite existing template with the same name")
	Template.AddCommand(Save)

	Use.PersistentFlags().BoolP("use-defaults", "f", false, "Uses default values in project.json instead of prompting the user")
	Use.PersistentFlags().StringP("log-level", "l", "error", "log-level for output")
	Template.AddCommand(Use)

	Template.AddCommand(Validate)

	Init.PersistentFlags().BoolP("force", "f", false, "Recreate directories if they exist")
	Root.AddCommand(Init)

	Root.AddCommand(ConfigureBashCompletion)

	Root.AddCommand(Template)

	Version.PersistentFlags().BoolP("dont-prettify", "", false, "Only print the version without fancy formatting")
	Root.AddCommand(Version)

	Root.Execute()
}
