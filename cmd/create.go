package cmd

import (
	"log"
	"github.com/owbird/raijin/internal/project"

	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new Raijin project",
	Long:  `Raijin creates a new project with frontend and backend setup. It accepts vite params for the frontend.`,
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]

		if name == "" {
			log.Fatalln("Name is required")
		}

		if err := project.CreateScaffold(name, args[1:]); err != nil {
			log.Fatalln(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
}
