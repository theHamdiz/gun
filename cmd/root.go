package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/theHamdiz/it"
)

var rootCmd = &cobra.Command{
	Use:   "gun",
	Short: "Gun is a CLI tool for generating Go project boilerplate code.",
	Long:  `Gun helps you quickly scaffold Go projects with Fiber, Tailwind CSS, and other modern tools.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		it.LogErrorWithStack(err)
		os.Exit(1)
	}
}
