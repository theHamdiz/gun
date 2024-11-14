package cmd

import (
	"github.com/spf13/cobra"
	"github.com/theHamdiz/gun/internal/generator"
	"github.com/theHamdiz/it"
)

var viewCmd = &cobra.Command{
	Use:   "view [resource]",
	Short: "Generate views for a resource",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		resourceName := args[0]
		fieldsStr, _ := cmd.Flags().GetString("fields")
		fields := parseFields(fieldsStr)

		err := generator.GenerateViews(resourceName, fields)
		if err != nil {
			return err
		}

		it.Infof("Views for '%s' created successfully!\n", resourceName)
		return nil
	},
}

func init() {
	generateCmd.AddCommand(viewCmd)
	viewCmd.Flags().String("fields", "", "Fields for the resource (e.g., 'Name:string Email:string')")
}
