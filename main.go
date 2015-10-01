package main

import "github.com/spf13/cobra"

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Creates boiler-plate code for a project",
	Run: func(cmd *cobra.Command, args []string) {
		tmpl, err := template.Get(args[0])
		if err != nil {
			panic(err)
		}

		/*
		 *err := tmpl.Persist()
		 *if err != nil {
		 *    panic(err)
		 *}
		 */
	},
}

func main() {
	mainCmd := &cobra.Command{
		Use: "main",
	}

	mainCmd.AddCommand(buildCmd)

	mainCmd.Execute()
}
