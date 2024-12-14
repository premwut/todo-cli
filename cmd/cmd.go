package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

type TodoItem struct {
	action string
	status TodoStatus
}

type TodoStatus string

const (
	TodoStatusPending TodoStatus = "pending"
	TodoStatusDone    TodoStatus = "done"
)

func TestCmd() {
	fmt.Println("Hello from cmd")
}

var GreetCmd = &cobra.Command{
	Use:   "greet",
	Short: "Prints a greeting message",
	Long:  "This command prints a custom greeting message to the console.",
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		if name == "" {
			fmt.Println("Hello, World!")
		} else {
			fmt.Printf("Hello, %s!\n", name)
			item := TodoItem{action: "Do coding", status: "smth"}
			fmt.Println("Item:", item)
		}
	},
}

var RootCmd = &cobra.Command{
	Use:   "todo",
	Short: "A to-do list CLI tool",
	Long:  "This is a simple CLI tool built with Cobra.",
}
