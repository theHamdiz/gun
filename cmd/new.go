package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/theHamdiz/gun/internal/project"
)

var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Create new entities like project",
}

var newProjectCmd = &cobra.Command{
	Use:   "project [name]",
	Short: "Create a new Go project with Fiber and Tailwind CSS",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		projectName := args[0]
		style, _ := cmd.Flags().GetString("style")
		withChannels, _ := cmd.Flags().GetBool("with-channels")
		withSignals, _ := cmd.Flags().GetBool("with-signals")

		err := project.CreateProject(projectName, style, withChannels, withSignals)
		if err != nil {
			return err
		}

		fmt.Printf("Project '%s' created successfully!\n", projectName)
		return nil
	},
}

func init() {
	newCmd.AddCommand(newProjectCmd)
	rootCmd.AddCommand(newCmd)

	newProjectCmd.Flags().String("style", "tailwind", "Styling framework to use (tailwind, shadcn, both)")
	newProjectCmd.Flags().Bool("with-channels", false, "Include channel utilities")
	newProjectCmd.Flags().Bool("with-signals", false, "Include signal handling utilities")
}
