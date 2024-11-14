package cmd

import (
	"github.com/spf13/cobra"
	"github.com/theHamdiz/gun/internal/generator"
	"github.com/theHamdiz/it"
)

var routeCmd = &cobra.Command{
	Use:   "route [resource]",
	Short: "Generate routes for a resource",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		resourceName := args[0]

		err := generator.GenerateRoute(resourceName)
		if err != nil {
			return err
		}

		it.Infof("Routes for '%s' created successfully!\n", resourceName)
		return nil
	},
}

func init() {
	generateCmd.AddCommand(routeCmd)
}
