package cmd

import (
	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate code components like models, handlers, routes, views, etc.",
}

func init() {
	rootCmd.AddCommand(generateCmd)
}
