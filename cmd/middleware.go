package cmd

import (
	"github.com/spf13/cobra"
	"github.com/theHamdiz/gun/internal/generator"
	"github.com/theHamdiz/it"
)

var middlewareCmd = &cobra.Command{
	Use:   "middleware [name]",
	Short: "Generate a middleware",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		middlewareName := args[0]

		err := generator.GenerateMiddleware(middlewareName)
		if err != nil {
			return err
		}

		it.Infof("Middleware '%s' created successfully!\n", middlewareName)
		return nil
	},
}

func init() {
	generateCmd.AddCommand(middlewareCmd)
}
