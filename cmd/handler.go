package cmd

import (
	"github.com/spf13/cobra"
	"github.com/theHamdiz/gun/internal/generator"
	"github.com/theHamdiz/it"
)

var handlerCmd = &cobra.Command{
	Use:   "handler [resource]",
	Short: "Generate handlers for a resource",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		resourceName := args[0]

		err := generator.GenerateHandler(resourceName)
		if err != nil {
			return err
		}

		it.Infof("Handlers for '%s' created successfully!\n", resourceName)
		return nil
	},
}

func init() {
	generateCmd.AddCommand(handlerCmd)
}
