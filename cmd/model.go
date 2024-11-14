package cmd

import (
	"strings"

	"github.com/spf13/cobra"
	"github.com/theHamdiz/gun/internal/generator"
	"github.com/theHamdiz/it"
)

var modelCmd = &cobra.Command{
	Use:   "model [name]",
	Short: "Generate a model",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		modelName := args[0]
		fieldsStr, _ := cmd.Flags().GetString("fields")
		fields := parseFields(fieldsStr)

		err := generator.GenerateModel(modelName, fields)
		if err != nil {
			return err
		}

		it.Infof("Model '%s' created successfully!\n", modelName)
		return nil
	},
}

func init() {
	generateCmd.AddCommand(modelCmd)
	modelCmd.Flags().String("fields", "", "Fields for the model (e.g., 'Name:string Email:string')")
}

func parseFields(fieldsStr string) []generator.Field {
	var fields []generator.Field
	if fieldsStr == "" {
		return fields
	}

	pairs := strings.Split(fieldsStr, " ")
	for _, pair := range pairs {
		parts := strings.Split(pair, ":")
		if len(parts) != 2 {
			continue
		}
		fields = append(fields, generator.Field{
			Name: parts[0],
			Type: parts[1],
		})
	}
	return fields
}
