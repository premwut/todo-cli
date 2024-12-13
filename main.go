package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var greetCmd = &cobra.Command{
	Use:   "greet",
	Short: "Prints a greeting message",
	Long:  "This command prints a custom greeting message to the console.",
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		if name == "" {
			fmt.Println("Hello, World!")
		} else {
			fmt.Printf("Hello, %s!\n", name)
		}
	},
}

var rootCmd = &cobra.Command{
	Use:   "todo",
	Short: "A to-do list CLI tool",
	Long:  "This is a simple CLI tool built with Cobra.",
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func main() {
	rootCmd.AddCommand(greetCmd)
	Execute()

}
