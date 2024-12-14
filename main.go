package main

import (
	"fmt"
	"os"
	"todo-cli/cmd"
)

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func readData() {
	file, err := os.Open("/var/lib/todo-cli/data.csv")
	if err != nil {
		fmt.Println("Hooray!")
		fmt.Printf("%v", file)
	} else {
		fmt.Printf("err: %v\n", err)
	}
}

func main() {
	cmd.RootCmd.AddCommand(cmd.GreetCmd)
	Execute()
	readData()
	cmd.TestCmd()
}
