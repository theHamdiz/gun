package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gun",
	Short: "Gun is a CLI tool for generating Go project boilerplate code.",
	Long:  `Gun helps you quickly scaffold Go projects with Fiber, Tailwind CSS, and other modern tools.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
